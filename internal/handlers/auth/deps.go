package auth

import (
	"Sechenovka/internal/model"
	"github.com/google/uuid"
)

type authService interface {
	Login(snils string, password string) (uuid.UUID, error)
	Register(user *model.User) error
}
