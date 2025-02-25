package questions

import (
	"errors"
)

type QuestionIn struct {
	QuestionId int `json:"question_id"`
}

type QuestionOut struct {
	QuestionText     string   `json:"question_text,omitempty"`
	ImgName          string   `json:"img_name"`
	Options          []Option `json:"options,omitempty"`
	PassNum          int      `json:"pass_num,omitempty"`
	IsMultipleChoice bool     `json:"is_multiple_choice"`
}

type Option struct {
	AnswerText     string `json:"response_text"`
	AnswerId       int    `json:"response_id"`
	Points         int    `json:"points"`
	NextQuestionId int    `json:"next_question_id"`
}

func (q *QuestionIn) Validate() error {
	if q.QuestionId <= 0 {
		return errors.New("wrong question id")
	}
	return nil
}
