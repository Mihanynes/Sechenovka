package user_responses

import (
	"time"
)

// Основная структура для сохранения данных в БД
type UserResponse struct {
	Id         int       `gorm:"type:int;primaryKey"`
	UserId     string    `gorm:"index;not null"`
	ResponseId int       `gorm:"type:int"`
	PassNum    int       `gorm:"type:int"`       // Correlation ID для отслеживания запросов
	CreatedAt  time.Time `gorm:"autoCreateTime"` // Время создания записи
}
