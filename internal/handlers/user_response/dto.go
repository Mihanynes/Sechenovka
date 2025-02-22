package user_response

import (
	"errors"
)

type SaveUserResponseIn struct {
	ResponseIds []int `json:"response_ids"`
	PassNum     int   `json:"pass_num"`
}

func (q *SaveUserResponseIn) Validate() error {
	if len(q.ResponseIds) == 0 {
		return errors.New("response ids is required")
	}
	if q.PassNum == 0 {
		return errors.New("pass num is required")
	}
	return nil
}

type GetUserResponsesIn struct {
	UserId  string `json:"user_id"`
	PassNum int    `json:"pass_num"`
}

type GetUserResponsesOut struct {
	QuestionText  string `json:"question_text"`
	AnswerText    string `json:"answer_text"`
	ResponseScore int    `json:"response_score"`
}

type GetUserResponsesOutList struct {
	Responses []GetUserResponsesOut `json:"responses"`
}

type GetUsersResultIn struct {
	PassNum int `json:"pass_num"`
}

type GetUsersResultOutList struct {
	UserResults []GetUsersResultOut `json:"user_results"`
}

type GetUsersResultOut struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserScore int    `json:"user_score"`
	IsFailed  bool   `json:"is_failed"`
}
