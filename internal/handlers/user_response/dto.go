package user_response

import (
	"errors"
	"github.com/google/uuid"
)

type SaveUserResponseIn struct {
	UserId        uuid.UUID `json:"user_id"`
	QuestionText  string    `json:"question_text"`
	CorrelationId string    `json:"correlation_id"`
	ResponseScore int       `json:"response_score"`
}

type GetUserScoreIn struct {
	UserId        uuid.UUID `json:"user_id"`
	CorrelationId string    `json:"correlation_id"`
}

func (q *SaveUserResponseIn) Validate() error {
	if q.QuestionText == "" {
		return errors.New("questions text is required")
	}
	if q.CorrelationId == "" {
		return errors.New("correlation id is required")
	}
	return nil
}
