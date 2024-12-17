package user_result

import "time"

type UserResult struct {
	Id        int       `gorm:"type:int;primaryKey"`
	UserId    string    `gorm:"index;not null"` // Идентификатор пользователя
	PassNum   int       `gorm:"index;not null"`
	Score     int       `gorm:"type:int;not null"`
	IsFailed  bool      `gorm:"type:boolean;not null"` //true - хорошо, false - плохо
	Timestamp time.Time `gorm:"autoCreateTime"`        // Время создания записи
}
