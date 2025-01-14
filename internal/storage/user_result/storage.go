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
func (s *UserResultStorage) CreateUserResult(userId model.UserId, passNum int) error {
	dal := UserResult{
		UserId:  userId.String(),
		PassNum: passNum,
	}
	return s.db.Create(&dal).Error
}

func (s *UserResultStorage) UpdateUserResult(userId model.UserId, passNum int, userScore int, isFailed bool) error {
	return s.db.Model(&UserResult{}).
		Where("user_id = ? AND pass_num = ?", userId.String(), passNum).
		Updates(map[string]interface{}{
			"total_score": userScore,
			"is_failed":   isFailed,
		}).Error
}

func (s *UserResultStorage) GetUsersResults(userIds []model.UserId) ([]UserResult, error) {
	var userResults []UserResult
	err := s.db.Where("user_id IN ?", model.ConvertUserIdsToStrings(userIds)).Find(&userResults).Order("created_at DESC").Error
	if err != nil {
		return nil, err
	}
	return userResults, nil
}
