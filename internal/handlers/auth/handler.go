package auth

type handler struct {
	authService          authService
	doctorPatientStorage doctorPatientStorage
}

func New(authService authService, doctorPatientStorage doctorPatientStorage) *handler {
	return &handler{
		authService:          authService,
		doctorPatientStorage: doctorPatientStorage,
	}
}
