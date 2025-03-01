package question_config

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_responses"
)

type userResponseStorage interface {
	GetLastUserResponse(userId model.UserId, quizId int) (*user_responses.UserResponse, error)
}
