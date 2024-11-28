package model

type UserResponse struct {
	UserId        int
	QuestionText  string
	Response      Response
	CorrelationId string
}

type Response struct {
	Answer string
	Score  int
}
