package model

type UserResponse struct {
	UserId        UserId
	QuestionId    int
	QuestionText  string
	Response      Response
	CorrelationId string
}

type Response struct {
	AnswerId   int
	AnswerText string
	Score      int
}
