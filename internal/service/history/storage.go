package history

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

type UserScore struct {
	Answer string `json:"answer"`
	Score  int    `json:"score"`
}

// Метод для получения суммы score по correlationId
func (s *historyStorage) GetUserScore(correlationId string) (int, error) {
	var totalScore int64 // Используем int64 для избежания переполнения

	// Выполняем запрос к базе данных
	err := s.db.Model(&UserResponse{}).
		Select("SUM(response_score)").
		Where("correlation_id = ?", correlationId).
		Scan(&totalScore).Error

	if err != nil {
		return 0, err
	}

	return int(totalScore), nil
}
