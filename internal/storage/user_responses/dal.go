package user_responses

import (
	"time"
)

// Основная структура для сохранения данных в БД
type UserResponse struct {
	Id     int    `gorm:"type:int;primaryKey"`
	UserId string `gorm:"index;not null"`
	//QuestionId    int       `gorm:"type:int;not null"`
	Response      Response  `gorm:"embedded;embeddedPrefix:response_"` // Ответ пользователя с оценкой
	CorrelationId string    `gorm:"type:varchar(36)"`                  // Correlation ID для отслеживания запросов
	Timestamp     time.Time `gorm:"autoCreateTime"`                    // Время создания записи
}

type Response struct {
	//AnswerId int
	AnswerText string `gorm:"type:text"`
	Score      int
}
