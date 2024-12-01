package user_info

import "github.com/google/uuid"

type User struct {
	Id         int       `gorm:"type:int;primaryKey;autoIncrement"`
	UserId     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"` // Уникальный бизнес-идентификатор
	FirstName  string    `gorm:"type:varchar;not null"`
	LastName   string    `gorm:"type:varchar;not null"`
	MiddleName string    `gorm:"type:varchar;not null"`
	Snils      string    `gorm:"type:varchar;not null"`
	Email      string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password   string    `gorm:"type:varchar(100);not null"`
	IsAdmin    bool      `gorm:"type:varchar(50);default:false;not null"`
}
