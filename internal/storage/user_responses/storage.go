package user_responses

import (
	"Sechenovka/internal/model"
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

func (s *storage) SaveUserResponse(userResponse *model.UserResponse) error {
	dal := &UserResponse{
		UserId: userResponse.UserId,
		Response: Response{
			Answer: userResponse.Response.Answer,
			Score:  userResponse.Response.Score,
		},
		CorrelationId: userResponse.CorrelationId,
	}
	return s.db.Create(dal).Error
}

// GetUserScore Метод для получения суммы score по correlationId
func (s *storage) GetUserScore(userId int, correlationId string) (int, error) {
	var totalScore int64

	err := s.db.Model(&UserResponse{}).
		Select("SUM(response_score)").
		Where("correlation_id = ?", correlationId).
		Where("user_id = ?", userId).
		Scan(&totalScore).Error

	if err != nil {
		return 0, err
	}

	return int(totalScore), nil
}
