package user

type User struct {
	UserID     string `gorm:"type:text;primaryKey"` // UUID как текст
	Username   string `gorm:"type:varchar;uniqueIndex;not null"`
	FirstName  string `gorm:"type:varchar;not null"`
	LastName   string `gorm:"type:varchar;not null"`
	MiddleName string `gorm:"type:varchar;not null"`
	Phone      string `gorm:"type:varchar"`
	Snils      string `gorm:"type:varchar;not null"`
	Email      string `gorm:"type:varchar(100);not null"`
	Password   string `gorm:"type:varchar(100);not null"`
	IsAdmin    bool   `gorm:"default:false;not null"`
}
