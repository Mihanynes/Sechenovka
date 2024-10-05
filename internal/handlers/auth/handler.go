package auth

type handler struct {
	authService authService
}

func New(authService authService) *handler {
	return &handler{
		authService: authService,
	}
}
