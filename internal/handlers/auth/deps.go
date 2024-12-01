package auth

import (
	"Sechenovka/internal/model"
)

type authService interface {
	Login(username string, password string) (string, error)
	Register(user *model.User) error
}
