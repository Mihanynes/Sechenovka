package user_result

import (
	"Sechenovka/internal/model"
	"gorm.io/gorm"
)

type UserResultStorage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserResultStorage {
	return &UserResultStorage{
		db: db,
	}
}

func (s *UserResultStorage) SaveUserResult(userId model.UserId, userScore, passNum int, isFailed bool) error {
	dal := UserResult{
		UserId:   userId.String(),
		Score:    userScore,
		PassNum:  passNum,
		IsFailed: isFailed,
	}
	return s.db.Create(&dal).Error
}

func (s *UserResultStorage) GetUsersResults(userIds []model.UserId) ([]UserResult, error) {
	var userResults []UserResult
	err := s.db.Where("user_id IN ?", userIds).Find(&userResults).Order("created_at DESC").Error
	if err != nil {
		return nil, err
	}
	return userResults, nil
}
