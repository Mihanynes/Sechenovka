package user

import "github.com/google/uuid"

type User struct {
	UserId     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"` // Уникальный бизнес-идентификатор
	FirstName  string    `gorm:"type:varchar;not null"`
	LastName   string    `gorm:"type:varchar;not null"`
	MiddleName string    `gorm:"type:varchar;not null"`
	Snils      string    `gorm:"type:varchar;uniqueIndex;not null"`
	Email      string    `gorm:"type:varchar(100);not null"`
	Password   string    `gorm:"type:varchar(100);not null"`
	IsAdmin    bool      `gorm:"type:varchar(50);default:false;not null"`
}
