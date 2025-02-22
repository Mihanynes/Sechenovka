package user_info

import (
	"Sechenovka/internal/model"
	"errors"
	"gorm.io/gorm"
)

type UserInfoStorage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserInfoStorage {
	return &UserInfoStorage{
		db: db,
	}
}

func (s *UserInfoStorage) GetUserByUsername(username string) (*UserInfo, error) {
	userFromDB := UserInfo{}
	result := s.db.First(&userFromDB, "username = ?", username)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &userFromDB, nil
}

func (s *UserInfoStorage) GetUserByUserId(userId model.UserId) (*UserInfo, error) {
	userFromDB := UserInfo{}
	result := s.db.First(&userFromDB, "user_id = ?", userId.String())
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &userFromDB, nil
}

func (s *UserInfoStorage) SaveUser(user *model.UserInfo, userId model.UserId) error {
	userToSave := UserInfo{
		UserID:     userId.String(),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Phone:      user.Phone,
		Snils:      user.Snils,
		Email:      user.Email,
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
