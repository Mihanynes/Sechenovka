package doctor_patient

import "Sechenovka/internal/storage/user"

type DoctorPatient struct {
	DoctorID  string    `gorm:"index;not null;foreignKey:DoctorID;references:UserID"`
	PatientID string    `gorm:"index;not null;foreignKey:PatientID;references:UserID"`
	User      user.User `gorm:"foreignKey:PatientID;references:UserID"`
}
