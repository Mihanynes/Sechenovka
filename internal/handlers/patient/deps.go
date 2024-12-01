package patient

import "Sechenovka/internal/model"

type patientInfoService interface {
	GetPatientInfo(userId model.UserId) (*model.PatientInfo, error)
}
