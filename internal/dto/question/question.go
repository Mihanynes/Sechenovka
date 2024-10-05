package question

import "Sechenovka/internal/service/history_saver"

type QuestionIn struct {
	CorrelationID string `json:"correlation_id"`
	QuestionText  string `json:"next_question_text"`
	Points        int    `json:"points"`
	Answer        string `json:"answer"`
}

type QuestionOut struct {
	QuestionText  string   `json:"question_text,omitempty"`
	Options       []Option `json:"options,omitempty"`
	CorrelationID string   `json:"correlation_id,omitempty"`
}

type Option struct {
	Answer           string `json:"answer"`
	Points           int    `json:"points"`
	NextQuestionText string `json:"next_question_text"`
}

func (q *QuestionIn) ToUserResponse() *history_saver.UserResponse {
	return &history_saver.UserResponse{
		Response: history_saver.Response{
			Answer: q.Answer,
			Score:  q.Points,
		},
		CorrelationId: q.CorrelationID,
	}
}
