package user_responses

import (
	"gorm.io/gorm"
)

// Основная структура для сохранения данных в БД
type UserResponse struct {
	gorm.Model
	UserId     string `gorm:"index;not null"`
	ResponseId int    `gorm:"index;type:int"`
	PassNum    int    `gorm:"type:int"`
}
