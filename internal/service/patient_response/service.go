package patient_response

import (
	"Sechenovka/internal/model"
	question_service "Sechenovka/internal/service/question_config"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
	"errors"
)

type service struct {
	userResponsesStorage *user_respons_storage.UserResponseStorage
	userResultStorage    *user_result.UserResultStorage
	questionsConfig      *question_service.QuestionConfigService
}

func New(
	storage *user_respons_storage.UserResponseStorage,
	userResultStorage *user_result.UserResultStorage,
	questionsConfig *question_service.QuestionConfigService,
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
	if question.ScoreToFail != nil && currentTotalScore >= *question.ScoreToFail {
		err = s.userResultStorage.SaveUserResult(userId, currentTotalScore, passNum, quizId, true)
		if err != nil {
			return false, err
		}
		return true, nil
		// TODO: тут бы послать уведомление врачу, что пациенту плохо
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
