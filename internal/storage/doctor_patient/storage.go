package doctor_patient

import (
	"Sechenovka/internal/model"
	"gorm.io/gorm"
)

type DoctorPatientsStorage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *DoctorPatientsStorage {
	return &DoctorPatientsStorage{
		db: db,
	}
}

func (s *DoctorPatientsStorage) GetPatientsIdsByDoctorId(doctorID model.UserId) ([]model.UserId, error) {
	var patients []string
	err := s.db.Model(DoctorPatient{}).
		Select("patient_id").
		Where("doctor_id = ?", doctorID.String()).
		Find(&patients).Error
	if err != nil {
		return nil, err
	}
	return model.UserIdsFromStrings(patients), nil
}

func (s *DoctorPatientsStorage) CheckPatientLinkedToDoctor(doctorId, patientId model.UserId) bool {
	var link DoctorPatient
	result := s.db.Where("doctor_id = ? AND patient_id = ?", doctorId.String(), patientId.String()).First(&link)
	return result.Error == nil
}

func (s *DoctorPatientsStorage) SaveDoctorPatientLink(doctorId, patientId model.UserId) error {
	link := DoctorPatient{
		DoctorID:  doctorId.String(),
		PatientID: patientId.String(),
	}
	result := s.db.Create(&link)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *DoctorPatientsStorage) DeleteDoctorPatientLink(doctorID, patientID model.UserId) error {
	result := s.db.Delete(&DoctorPatient{}, "doctor_id = ? AND patient_id = ?", doctorID, patientID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
