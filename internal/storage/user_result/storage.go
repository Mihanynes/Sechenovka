package user_result

import (
	"Sechenovka/internal/model"
	"errors"
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

func (s *UserResultStorage) UpdateIsViewed(userId model.UserId, quizId int, passNum int, isViewed bool) error {
	return s.db.Model(&UserResult{}).
		Where("user_id = ? AND quiz_id = ? AND pass_num = ?", userId.String(), quizId, passNum).
		Update("is_viewed", isViewed).Error
}

func (s *UserResultStorage) SaveUserResult(userId model.UserId, userScore int, passNum int, quizId int, isFailed bool) error {
	dal := UserResult{
		UserID:     userId.String(),
		PassNum:    passNum,
		TotalScore: userScore,
		IsFailed:   isFailed,
		QuizId:     quizId,
	}
	return s.db.Create(&dal).Error
}

func (s *UserResultStorage) GetUsersResults(userIds []model.UserId) ([]UserResult, error) {
	var userResults []UserResult
	userIdsStr := model.ConvertUserIdsToStrings(userIds)

	err := s.db.Where("user_id IN ?", userIdsStr).
		Find(&userResults).Order("created_at DESC").Error
	if err != nil {
		return nil, err
	}

	return userResults, nil
}

func (s *UserResultStorage) GetUserResultByQuizId(userId model.UserId, quizId int) (*UserResult, error) {
	userResult := UserResult{}
	err := s.db.Where("user_id = ?", userId.String()).
		Where("quiz_id = ?", quizId).Last(&userResult).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &userResult, nil
}
