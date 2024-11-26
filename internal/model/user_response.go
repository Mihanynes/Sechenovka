package model

type UserResponse struct {
	UserId        int
	Response      Response
	CorrelationId string
}

type Response struct {
	Answer string
	Score  int
}
