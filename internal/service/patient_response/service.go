package patient_response

import (
	"Sechenovka/internal/model"
	question_service "Sechenovka/internal/service/question_config"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
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

// SaveUserResponse возвращает true первым аргументом, если пациент перешел порог баллов
func (s *service) SaveUserResponse(userId model.UserId, responseId, passNum int) (bool, error) {
	err := s.userResponsesStorage.SaveUserResponse(userId, responseId, passNum)
	if err != nil {
		return false, err
	}

	currentTotalScore, err := s.userResponsesStorage.GetUserTotalScore(userId, passNum)
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
			err = s.userResultStorage.SaveUserResult(userId, currentTotalScore, passNum, isFailed)
			if err != nil {
				return false, err
			}
		}
	}

	return isFailed, nil
}
