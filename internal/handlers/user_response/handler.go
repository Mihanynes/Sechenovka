package user_response

import (
	question_service "Sechenovka/internal/service/quiz"
	"Sechenovka/internal/storage/doctor_patient"
	"Sechenovka/internal/storage/user"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
)

type handler struct {
	userResponseService   userResponseService
	userResponseStorage   *user_respons_storage.UserResponseStorage
	questionConfigService *question_service.Service
	doctorPatientsStorage *doctor_patient.DoctorPatientsStorage
	userResultStorage     *user_result.UserResultStorage
	userInfoStorage       *user.UserStorage
}

func New(
	userResponseService userResponseService,
	userResponseStorage *user_respons_storage.UserResponseStorage,
	questionConfigService *question_service.Service,
	doctorPatientsStoragee *doctor_patient.DoctorPatientsStorage,
	userResultStorage *user_result.UserResultStorage,
	userInfoStorage *user.UserStorage,
) *handler {
	return &handler{
		userResponseService:   userResponseService,
		userResponseStorage:   userResponseStorage,
		questionConfigService: questionConfigService,
		doctorPatientsStorage: doctorPatientsStoragee,
		userResultStorage:     userResultStorage,
		userInfoStorage:       userInfoStorage,
	}
}
