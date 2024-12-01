package auth

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user"
	"github.com/google/uuid"
)

type userStorage interface {
	SaveUser(user *model.User, userId uuid.UUID) error
	GetUserByUsername(username string) (*user.User, error)
}
