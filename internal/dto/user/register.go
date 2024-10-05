package user

import "errors"

type RegisterIn struct {
	UserName        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

func (r *RegisterIn) Validate() error {
	if r.UserName == "" {
		return errors.New("name is required")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if len(r.Password) < 8 {
		return errors.New("password is required")
	}
	if r.PasswordConfirm != r.Password {
		return errors.New("different passwords")
	}
	return nil
}
