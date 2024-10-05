package auth

import "Sechenovka/internal/models"

type authService interface {
	Login(username, password string) error
	Register(user *models.User) error
}
