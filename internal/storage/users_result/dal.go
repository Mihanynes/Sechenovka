package users_result

import "time"

type UserResult struct {
	Id        int       `gorm:"type:int;primaryKey"`
	UserId    int       `gorm:"index;not null"` // Идентификатор пользователя
	Score     int       `gorm:"type:int;not null"`
	Result    bool      `gorm:"type:boolean;not null"`
	Timestamp time.Time `gorm:"autoCreateTime"` // Время создания записи
}
