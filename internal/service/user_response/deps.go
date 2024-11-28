package user_response

import (
	"Sechenovka/internal/model"
)

type userResponsesStorage interface {
	SaveUserResponse(userResponse *model.UserResponse) error
	GetUserScore(correlationId string) (int, error)
}

type questionsConfig interface {
	GetOptionsByQuestionText(questionText string) (*model.Question, error)
}
