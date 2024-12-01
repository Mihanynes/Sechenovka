package doctor_patient

import (
	"Sechenovka/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) GetPatientsIdsByDoctorId(doctorID uuid.UUID) ([]model.UserId, error) {
	patients := make([]model.UserId, 0)
	err := s.db.Model(DoctorPatient{}).
		Select("patient_id").
		Where("doctor_id = ?", doctorID).
		Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return patients, nil
}
