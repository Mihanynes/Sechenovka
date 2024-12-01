package model

type User struct {
	Id       int    `gorm:"type:int;primaryKey;autoIncrement"`
	Username string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(100);not null"`
	IsAdmin  bool   `gorm:"type:varchar(50);default:false;not null"`
}

type UserId int64
