package questions

import (
	"errors"
	"strings"
)

type QuestionIn struct {
	QuestionText string `json:"question_text"`
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
	if strings.TrimSpace(q.QuestionText) == "" {
		return errors.New("questions text is required")
	}
	return nil
}
