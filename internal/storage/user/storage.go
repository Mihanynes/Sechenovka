package user

import (
	"Sechenovka/internal/model"
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
		return nil, model.ErrUserNotFound
	}
	return &userFromDB, nil
}

func (s *UserStorage) GetUserByUserId(userId model.UserId) (*User, error) {
	userFromDB := User{}
	result := s.db.First(&userFromDB, "user_id = ?", userId.String())
	if result.Error != nil {
		return nil, model.ErrUserNotFound
	}
	return &userFromDB, nil
}

func (s *UserStorage) GetAllUsers() ([]User, error) {
	var users []User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *UserStorage) DeleteUserByID(userID string) error {
	result := s.db.Where("user_id = ?", userID).Delete(&User{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return model.ErrUserNotFound // Возвращаем ошибку, если пользователь не найден
	}
	return nil
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
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    user.IsAdmin,
	}
	result := s.db.Create(&userToSave)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
