package question_config

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_responses"
)

type userResponseStorage interface {
	GetLastUserResponse(userId model.UserId) (*user_responses.UserResponse, error)
}
