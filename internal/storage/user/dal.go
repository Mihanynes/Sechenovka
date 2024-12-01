package user

type User struct {
	UserId     string `gorm:"type:text;primaryKey"` // UUID как текст
	UserName   string `gorm:"type:text;uniqueIndex;not null"`
	FirstName  string `gorm:"type:varchar;not null"`
	LastName   string `gorm:"type:varchar;not null"`
	MiddleName string `gorm:"type:varchar;not null"`
	Snils      string `gorm:"type:varchar;uniqueIndex;not null"`
	Email      string `gorm:"type:varchar(100);not null"`
	Password   string `gorm:"type:varchar(100);not null"`
	IsAdmin    bool   `gorm:"default:false;not null"`
}
