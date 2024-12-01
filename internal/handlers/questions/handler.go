package questions

type handler struct {
	questionService questionService
}

func New(questionService questionService) *handler {
	return &handler{
		questionService: questionService,
	}
}
