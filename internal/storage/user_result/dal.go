package user_result

import (
	"Sechenovka/internal/storage/user"
	"gorm.io/gorm"
)

type UserResult struct {
	gorm.Model
	UserID     string    `gorm:"index;not null"` // Идентификатор пользователя
	PassNum    int       `gorm:"type:int;not null"`
	QuizId     int       `gorm:"type:int;not null"`
	TotalScore int       `gorm:"type:int;"`
	IsFailed   bool      `gorm:"type:boolean;not null"`
	User       user.User `gorm:"references:UserID"`
}
