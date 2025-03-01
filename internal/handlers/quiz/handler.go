package quiz

type handler struct {
	questionService quizService
}

func New(questionService quizService) *handler {
	return &handler{
		questionService: questionService,
	}
}
