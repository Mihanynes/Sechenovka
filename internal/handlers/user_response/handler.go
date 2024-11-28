package user_response

type handler struct {
	userResponseService userResponseService
}

func New(userResponseService userResponseService) *handler {
	return &handler{
		userResponseService: userResponseService,
	}
}
