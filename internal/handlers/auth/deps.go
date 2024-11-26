package auth

import "Sechenovka/internal/model"

type authService interface {
	Login(username, password string) error
	Register(user *model.User) error
}
