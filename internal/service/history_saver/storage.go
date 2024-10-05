package history_saver

import (
	"gorm.io/gorm"
)

type historyStorage struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) *historyStorage {
	return &historyStorage{
		db: db,
	}
}

func (s *historyStorage) SaveUserResponse(userResponse *UserResponse) error {

	if err := s.db.Create(userResponse).Error; err != nil {
		return err
	}

	return nil
}
