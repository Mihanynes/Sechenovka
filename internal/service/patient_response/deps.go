package patient_response

import (
	"Sechenovka/internal/model"
)

type userResponsesStorage interface {
	SaveUserResponse(userResponse *model.UserResponse) error
	GetUserTotalScore(userId model.UserId, correlationId string) (int, error)
	GetUserResponses(userId model.UserId) ([]*model.UserResponse, error)
}

type questionsConfig interface {
	GetOptionsByQuestionText(questionText string) (*model.Question, error)
}
