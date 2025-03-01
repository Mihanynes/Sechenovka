package user_responses

import (
	"Sechenovka/internal/storage/user"
	"gorm.io/gorm"
)

// Основная структура для сохранения данных в БД
type UserResponse struct {
	gorm.Model
	UserID     string    `gorm:"index;not null"`
	ResponseId int       `gorm:"type:int"`
	PassNum    int       `gorm:"type:int"`
	QuizId     int       `gorm:"type:int"`
	User       user.User `gorm:"references:UserID"`
}
