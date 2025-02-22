package user_info

type UserInfo struct {
	UserID     string `gorm:"type:text;primaryKey"` // UUID как текст
	FirstName  string `gorm:"type:varchar;not null"`
	LastName   string `gorm:"type:varchar;not null"`
	MiddleName string `gorm:"type:varchar;not null"`
	Phone      string `gorm:"type:varchar"`
	Snils      string `gorm:"type:varchar;not null"`
	Email      string `gorm:"type:varchar(100);not null"`
}
