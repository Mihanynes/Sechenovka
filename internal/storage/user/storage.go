package user

import (
	"Sechenovka/internal/model"
	"errors"
	"gorm.io/gorm"
)

type UserStorage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (s *UserStorage) GetUserByUsername(username string) (*User, error) {
	userFromDB := User{}
	result := s.db.First(&userFromDB, "username = ?", username)
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &userFromDB, nil
}

func (s *UserStorage) GetUserByUserId(userId model.UserId) (*User, error) {
	userFromDB := User{}
	result := s.db.First(&userFromDB, "user_id = ?", userId.String())
	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &userFromDB, nil
}

//func (s *UserStorage) GetPatientsByDoctorId(doctorID model.UserID) ([]User, error) {
//	var patients []User
//	err := s.db.Table("users").
//		Select("users.*").
//		Joins("join doctor_patients on doctor_patients.patient_id = users.user_id").
//		Where("doctor_patients.doctor_id = ?", doctorID).
//		Find(&patients).Error
//	if err != nil {
//		return nil, errors.New("could not retrieve patients: " + err.Error())
//	}
//	return patients, nil
//}

func (s *UserStorage) SaveUser(user *model.User, userId model.UserId) error {
	userToSave := User{
		UserID:     userId.String(),
		Username:   user.Username,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Phone:      user.Phone,
		Snils:      user.Snils,
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    user.IsAdmin,
	}
	result := s.db.Create(&userToSave)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return errors.New("user already exists")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
