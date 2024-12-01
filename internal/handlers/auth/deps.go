package auth

import (
	"Sechenovka/internal/model"
)

type authService interface {
	Login(snils string, password string) (string, error)
	Register(user *model.User) error
}
