package patient_response

import (
	"Sechenovka/internal/model"
	question_service "Sechenovka/internal/service/question_config"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
)

const isEnded = true

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
func (s *service) SaveUserResponse(userId model.UserId, responseId, passNum int) (bool, error) {
	err := s.userResponsesStorage.SaveUserResponse(userId, responseId, passNum)
	if err != nil {
		return false, err
	}

	prevUserResponses, err := s.userResponsesStorage.GetUserResponsesByPassNum(userId, passNum)
	if err != nil {
		return false, err
	}

	currentTotalScore, err := s.countCurrentScore(prevUserResponses)
	if err != nil {
		return false, err
	}

	question, err := s.questionsConfig.GetQuestionByResponseId(responseId)
	if err != nil {
		return false, err
	}

	// Если пациент перешел порог, то надо отсылать уведомление
	var isFailed bool
	if question.ScoreToFail != nil && currentTotalScore >= *question.ScoreToFail {
		isFailed = true
		return isFailed, nil
		// TODO: тут бы послать уведомление врачу, что пациенту плохо
	}

	// Если пациент завершил тест
	for _, option := range question.Options {
		if option.AnswerId == responseId && option.IsEnded {
			err = s.userResultStorage.UpdateUserResult(userId, currentTotalScore, passNum, isFailed)
			if err != nil {
				return false, err
			}
			return isEnded, nil
		}
	}

	return isFailed, nil
}

func (s *service) countCurrentScore(prevUserResponses []user_respons_storage.UserResponse) (int, error) {
	currentScore := 0
	for _, resp := range prevUserResponses {
		respConf, err := s.questionsConfig.GetOptionByResponseId(resp.ResponseId)
		if err != nil {
			return 0, err
		}
		currentScore += respConf.Points
	}
	return currentScore, nil
}
