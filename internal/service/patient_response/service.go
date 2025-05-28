package patient_response

import (
	"Sechenovka/internal/model"
	question_service "Sechenovka/internal/service/quiz"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type service struct {
	userResponsesStorage *user_respons_storage.UserResponseStorage
	userResultStorage    *user_result.UserResultStorage
	questionsConfig      *question_service.Service
}

func New(
	storage *user_respons_storage.UserResponseStorage,
	userResultStorage *user_result.UserResultStorage,
	questionsConfig *question_service.Service,
) *service {
	return &service{
		userResponsesStorage: storage,
		questionsConfig:      questionsConfig,
		userResultStorage:    userResultStorage,
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

		// Отправляем уведомление врачу
		go func() {
			payload := map[string]string{
				"title": "Тревога",
				"body":  "Пациенту стало плохо. Срочно проверьте его состояние!",
			}
			body, _ := json.Marshal(payload)

			resp, err := http.Post("http://push_sender:8081/api/notify", "application/json", bytes.NewReader(body))
			if err != nil {
				fmt.Println("Ошибка при отправке уведомления:", err)
				return
			}
			defer resp.Body.Close()
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
