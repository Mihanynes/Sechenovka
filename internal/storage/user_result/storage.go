package user_result

import (
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

func (s *storage) SaveUserResult(userId, userScore int, isFailed bool) error {
	dal := UserResult{
		UserId:   userId,
		Score:    userScore,
		IsFailed: isFailed,
	}
	return s.db.Create(&dal).Error
}
