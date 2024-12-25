package user_result

import (
	"Sechenovka/internal/storage/user"
	"gorm.io/gorm"
)

type UserResult struct {
	gorm.Model
	UserId     string    `gorm:"index;not null"` // Идентификатор пользователя
	PassNum    int       `gorm:"type:int;not null"`
	TotalScore int       `gorm:"type:int;"`
	IsFailed   bool      `gorm:"type:boolean;not null"` //true - хорошо, false - плохо
	User       user.User `gorm:"references:UserID"`
}
