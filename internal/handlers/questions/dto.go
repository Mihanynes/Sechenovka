package questions

import (
	"Sechenovka/internal/model"
	"errors"
)

type QuestionIn struct {
	CorrelationId string `json:"correlation_id"`
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

func (q *QuestionIn) Validate() error {
	if q.QuestionText == "" {
		return errors.New("questions text is required")
	}
	if q.CorrelationId == "" {
		return errors.New("correlation id is required")
	}
	if q.Answer == "" {
		return errors.New("answer is required")
	}
	return nil
}

func (q *QuestionIn) ValidateCorrelationId() error {
	if q.CorrelationId == "" {
		return errors.New("correlation id is required")
	}
	return nil
}

func (q *QuestionIn) ToUserResponse() *model.UserResponse {
	return &model.UserResponse{
		Response: model.Response{
			Answer: q.Answer,
			Score:  q.Points,
		},
		CorrelationId: q.CorrelationId,
	}
}
