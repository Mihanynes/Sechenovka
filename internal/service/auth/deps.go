package auth

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user"
)

type userStorage interface {
	SaveUser(user *model.User, userId model.UserId) error
	GetUserByUsername(username string) (*user.User, error)
}
