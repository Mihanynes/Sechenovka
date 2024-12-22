package user_result

import "gorm.io/gorm"

type UserResult struct {
	gorm.Model
	UserId     string `gorm:"index;not null"` // Идентификатор пользователя
	PassNum    int    `gorm:"type:int;not null"`
	TotalScore int    `gorm:"type:int;not null"`
	IsFailed   bool   `gorm:"type:boolean;not null"` //true - хорошо, false - плох
}
