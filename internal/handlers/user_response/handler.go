package user_response

type handler struct {
	userResponseService userResponseService
	userResponseStorage userResponseStorage
}

func New(userResponseService userResponseService, userResponseStorage userResponseStorage) *handler {
	return &handler{
		userResponseService: userResponseService,
		userResponseStorage: userResponseStorage,
	}
}
