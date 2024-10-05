package question

import "Sechenovka/internal/service/history_saver"

type QuestionIn struct {
	UserID        int    `json:"user_id"`
	CorrelationID string `json:"correlation_id"`
	QuestionText  string `json:"question_text"`
	Points        int    `json:"points"`
	Answer        string `json:"answer"`
}

type QuestionOut struct {
	QuestionText string   `json:"question_text"`
	Options      []Option `json:"options"`
}

type Option struct {
	Answer           string `json:"answer"`
	Points           int    `json:"points"`
	NextQuestionText string `json:"next_question_text"`
}

func (q *QuestionIn) ToUserResponse() *history_saver.UserResponse {
	return &history_saver.UserResponse{
		UserId: q.UserID,
		Response: history_saver.Response{
			Answer: q.Answer,
			Score:  q.Points,
		},
		CorrelationId: q.CorrelationID,
	}
}
