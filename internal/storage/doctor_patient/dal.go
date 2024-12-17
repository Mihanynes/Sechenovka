package doctor_patient

type DoctorPatient struct {
	DoctorId  string `gorm:"index;not null"`
	PatientId string `gorm:"index;not null"`
}
