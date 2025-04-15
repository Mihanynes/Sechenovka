package auth

import (
	"Sechenovka/internal/model"
)

type authService interface {
	Login(username string, password string) (*LoginOut, error)
	Register(user *model.User) (*model.UserId, error)
}

type doctorPatientStorage interface {
	SaveDoctorPatientLink(doctorId, patientId model.UserId) error
}
