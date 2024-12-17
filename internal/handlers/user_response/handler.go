package user_response

import (
	question_service "Sechenovka/internal/service/question_config"
	"Sechenovka/internal/storage/doctor_patient"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
)

type handler struct {
	userResponseService   userResponseService
	userResponseStorage   *user_respons_storage.UserResponseStorage
	questionConfigService *question_service.QuestionConfigService
	doctorPatientsStorage *doctor_patient.DoctorPatientsStorage
	userResultStorage     *user_result.UserResultStorage
}

func New(
	userResponseService userResponseService,
	userResponseStorage *user_respons_storage.UserResponseStorage,
	questionConfigService *question_service.QuestionConfigService,
	doctorPatientsStoragee *doctor_patient.DoctorPatientsStorage,
	userResultStorage *user_result.UserResultStorage,
) *handler {
	return &handler{
		userResponseService:   userResponseService,
		userResponseStorage:   userResponseStorage,
		questionConfigService: questionConfigService,
		doctorPatientsStorage: doctorPatientsStoragee,
		userResultStorage:     userResultStorage,
	}
}
