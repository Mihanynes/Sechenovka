package patient_response

import (
	"Sechenovka/internal/model"
)

type service struct {
	userResponsesStorage userResponsesStorage
	questionsConfig      questionsConfig
}

func New(storage userResponsesStorage, questionsConfig questionsConfig) *service {
	return &service{
		userResponsesStorage: storage,
		questionsConfig:      questionsConfig,
	}
}

// SaveUserResponse возвращает true первым аргументом, если пациент перешел порог баллов
func (s *service) SaveUserResponse(userResponse *model.UserResponse) (bool, error) {
	err := s.userResponsesStorage.SaveUserResponse(userResponse)
	if err != nil {
		return false, err
	}

	currentTotalScore, err := s.userResponsesStorage.GetUserTotalScore(userResponse.UserId, userResponse.CorrelationId)
	if err != nil {
		return false, err
	}

	question, err := s.questionsConfig.GetOptionsByQuestionText(userResponse.QuestionText)
	if err != nil {
		return false, err
	}
	// Если пациент перешел порог, то надо отсылать уведомление
	if question.ScoreToFail != nil && currentTotalScore >= *question.ScoreToFail {
		return true, nil
		// TODO: тут бы послать уведомление врачу, что пациенту плохо
	}
	return false, nil
}
