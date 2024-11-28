package user_response

type SaveUserResponseIn struct {
	UserId        int    `json:"user_id"`
	QuestionText  string `json:"question_text"`
	CorrelationId string `json:"correlation_id"`
	ResponseScore int    `json:"response_score"`
}
