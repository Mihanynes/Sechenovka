package auth

import (
	"Sechenovka/internal/model"
	"github.com/google/uuid"
)

type userStorage interface {
	SaveUser(user *model.User, userId uuid.UUID) error
}
