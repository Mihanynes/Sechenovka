package user

import (
	"Sechenovka/internal/model"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) GetUserBySnils(snils string) (*User, error) {
	userFromDB := User{}
	result := s.db.First(&userFromDB, "snils = ?", snils)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &userFromDB, nil
}

func (s *storage) GetUserByUserId(userId model.UserId) (*User, error) {
	userFromDB := User{}
	result := s.db.First(&userFromDB, "userId = ?", userId)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &userFromDB, nil
}

func (s *storage) SaveUser(user *model.User, userId uuid.UUID) error {
	userToSave := User{
		UserId:     userId,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Snils:      user.Snils,
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    user.IsAdmin,
	}
	result := s.db.Create(&userToSave)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return errors.New("user already exists")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
