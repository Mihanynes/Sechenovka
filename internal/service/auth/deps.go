package auth

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_info"
	"github.com/google/uuid"
)

type userStorage interface {
	SaveUser(user *model.User, userId uuid.UUID) error
	GetUser(snils string) (*user_info.User, error)
}
