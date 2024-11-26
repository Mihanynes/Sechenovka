package model

type User struct {
	Id       int    `gorm:"type:int;primaryKey;autoIncrement"`
	Username string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Role     string `gorm:"type:varchar(50);default:'user';not null"`
}

type UserId int64
