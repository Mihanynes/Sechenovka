package user_history

import "time"

// Основная структура для сохранения данных в БД
type UserResponse struct {
	Id            int       `gorm:"type:int;primaryKey"`
	UserId        int       `gorm:"index;not null"`                    // Идентификатор пользователя
	Response      Response  `gorm:"embedded;embeddedPrefix:response_"` // Ответ пользователя с оценкой
	CorrelationId string    `gorm:"type:varchar(36)"`                  // Correlation ID для отслеживания запросов
	Timestamp     time.Time `gorm:"autoCreateTime"`                    // Время создания записи
}

type Response struct {
	Answer string `json:"answer"` // Сам ответ пользователя
	Score  int    `json:"score"`  // Оценка ответа
}
