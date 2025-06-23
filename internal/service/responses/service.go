package responses

import (
	"Sechenovka/config"
	"Sechenovka/internal/model"
	question_service "Sechenovka/internal/service/quiz"
	"Sechenovka/internal/storage/doctor_patient"
	"Sechenovka/internal/storage/user"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const tgProducerPath = "http://telegram_producer:8082/send"

const (
	selfCheckMessage         = "Пациент %s %s %s завершил тест о состоянии здоровья c плохим результатом. Свяжитесь с ним."
	takingMedicationsMessage = "Пациент %s %s %s завершил тест о приеме препаратов c плохим результатом. Свяжитесь с ним."
)

var quizToMessage = map[int]string{
	config.SelfCheckQuiz:         selfCheckMessage,
	config.TakingMedicationsQuiz: takingMedicationsMessage,
}

type service struct {
	userResponsesStorage  *user_respons_storage.UserResponseStorage
	userResultStorage     *user_result.UserResultStorage
	questionsConfig       *question_service.Service
	doctorPatientsStorage *doctor_patient.DoctorPatientsStorage
	userStorage           *user.UserStorage
}

func New(
	storage *user_respons_storage.UserResponseStorage,
	userResultStorage *user_result.UserResultStorage,
	questionsConfig *question_service.Service,
	doctorPatientsStorage *doctor_patient.DoctorPatientsStorage,
	userStorage *user.UserStorage,
) *service {
	return &service{
		userResponsesStorage:  storage,
		questionsConfig:       questionsConfig,
		userResultStorage:     userResultStorage,
		doctorPatientsStorage: doctorPatientsStorage,
		userStorage:           userStorage,
	}
}

// SaveUserResponse возвращает true первым аргументом, если пациент перешел завершил тест
func (s *service) SaveUserResponses(userId model.UserId, responseIds []int, passNum int, quizId int) (bool, error) {
	if len(responseIds) == 0 {
		return false, errors.New("no responses to save")
	}
	question, err := s.questionsConfig.GetQuestionByResponseId(responseIds[0], quizId)
	if err != nil {
		return false, err
	}

	if !question.IsMultipleChoice && len(responseIds) > 1 {
		return false, errors.New("wrong number of responses")
	}

	for _, responseId := range responseIds {
		err := s.userResponsesStorage.SaveUserResponse(userId, responseId, passNum, quizId)
		if err != nil {
			return false, err
		}

	}

	prevUserResponses, err := s.userResponsesStorage.GetUserResponsesByPassNum(userId, passNum, quizId)
	if err != nil {
		return false, err
	}

	currentTotalScore, err := s.countCurrentScore(prevUserResponses, quizId)
	if err != nil {
		return false, err
	}

	// Если пациент перешел порог, то надо отсылать уведомление
	if checkIsFailed(currentTotalScore, question.ScoreToFail) {
		err = s.userResultStorage.SaveUserResult(userId, currentTotalScore, passNum, quizId, true)
		if err != nil {
			return false, err
		}
		log.Info("Пациент завершил тест с плохим результатом")

		go func() {
			doctorId, err := s.doctorPatientsStorage.GetDoctorIdByPatientId(userId)
			if err != nil {
				log.Error(err.Error())

			}
			doctor, err := s.userStorage.GetUserByUserId(*doctorId)
			if err != nil {
				log.Error(err.Error())

			}
			if doctor.ChatId == nil {
				log.Error("Нет чата для врача")

			}
			patient, err := s.userStorage.GetUserByUserId(userId)
			if err != nil {
				log.Error(err.Error())

			}
			message := fmt.Sprintf(quizToMessage[quizId], patient.FirstName, patient.LastName, patient.MiddleName)
			params := url.Values{}
			params.Set("chatId", strconv.FormatInt(*doctor.ChatId, 10))
			params.Set("message", message)
			requestUrl := tgProducerPath + "?" + params.Encode()
			_, err = http.Post(requestUrl, "application/x-www-form-urlencoded", strings.NewReader(""))
			if err != nil {
				log.Error(fmt.Sprintf("Ошибка при отправке уведомления: %v", err))
				return

			}
		}()

		return true, nil
	}

	// Если пациент завершил тест
	for _, option := range question.Options {
		if option.AnswerId == responseIds[0] && option.IsEnded {
			err = s.userResultStorage.SaveUserResult(userId, currentTotalScore, passNum, quizId, false)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}

	return false, nil
}

func (s *service) countCurrentScore(prevUserResponses []user_respons_storage.UserResponse, quizId int) (int, error) {
	currentScore := 0
	for _, resp := range prevUserResponses {
		respConf, err := s.questionsConfig.GetOptionByResponseId(resp.ResponseId, quizId)
		if err != nil {
			return 0, err
		}
		currentScore += respConf.Points
	}
	return currentScore, nil
}

func checkIsFailed(currentScore int, scoreToFail *int) bool {
	if scoreToFail == nil || *scoreToFail == 0 {
		return false
	}
	if *scoreToFail > 0 {
		return currentScore >= *scoreToFail
	} else {
		return currentScore <= (-1)*(*scoreToFail)
	}
}
