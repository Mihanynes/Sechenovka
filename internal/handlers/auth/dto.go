package auth

import (
	"errors"
)

type LoginIn struct {
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type LoginOut struct {
	UserId  string `json:"userId" `
	IsAdmin bool   `json:"isAdmin"`
}

type RegisterUserIn struct {
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type RegisterAdminIn struct {
	Username        string `json:"username"`
	FirstName       string `json:"first_name"`
	MiddleName      string `json:"middle_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	AdminToken      string `json:"admin_token"`
}

func (r *RegisterUserIn) ValidateUser() error {
	if r.FirstName == "" {
		return errors.New("name is required")
	}
	if r.LastName == "" {
		return errors.New("last name is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if r.PasswordConfirm != r.Password {
		return errors.New("different passwords")
	}
	return nil
}

func (r *RegisterAdminIn) ValidateAdmin() error {
	if r.FirstName == "" {
		return errors.New("name is required")
	}
	if r.LastName == "" {
		return errors.New("last name is required")
	}
	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if r.PasswordConfirm != r.Password {
		return errors.New("different passwords")
	}
	if r.AdminToken != "sechenovka" {
		return errors.New("admin token is required")
	}
	return nil
}
