package telegram

import (
	"Sechenovka/internal/handlers/auth"
)

type userStorage interface {
	SetField(field string, value interface{}, userID string) error
}

type authService interface {
	Login(username string, password string) (*auth.LoginOut, error)
}
