package patient_info

import (
	"Sechenovka/internal/model"
	"errors"
)

type service struct {
	userStorage userStorage
}

func New(userStorage userStorage) *service {
	return &service{userStorage: userStorage}
}

func (s *service) GetPatientInfo(userId model.UserId) (*model.PatientInfo, error) {
	user, err := s.userStorage.GetUserByUserId(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user is nil")
	}
	if user.IsAdmin {
		return nil, errors.New("user is not a patient")
	}

	return &model.PatientInfo{
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Snils:      user.Snils,
		Email:      user.Email,
	}, err
}
