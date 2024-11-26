package user_history

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
	if err := s.db.Create(dal).Error; err != nil {
		return err
	}
	return nil
}

// GetUserScore Метод для получения суммы score по correlationId
func (s *storage) GetUserScore(correlationId string) (int, error) {
	var totalScore int64

	err := s.db.Model(&UserResponse{}).
		Select("SUM(response_score)").
		Where("correlation_id = ?", correlationId).
		Scan(&totalScore).Error

	if err != nil {
		return 0, err
	}

	return int(totalScore), nil
}
