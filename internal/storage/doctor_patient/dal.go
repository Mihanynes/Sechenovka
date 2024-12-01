package doctor_patient

import "github.com/google/uuid"

type DoctorPatient struct {
	DoctorID  uuid.UUID `gorm:"type:uuid;not null;index"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index"`
}
