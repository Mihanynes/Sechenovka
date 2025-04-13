package user_response

import (
	"errors"
	"time"
)

type SaveUserResponseIn struct {
	ResponseIds []int `json:"response_ids"`
	PassNum     int   `json:"pass_num"`
	QuizId      int   `json:"quiz_id"`
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

type GetUserResponsesOut struct {
	QuizId        int    `json:"quiz_id"`
	QuestionText  string `json:"question_text"`
	AnswerText    string `json:"answer_text"`
	ResponseScore int    `json:"response_score"`
	PassNum       int    `json:"pass_num"`
	ResponseId    int    `json:"response_id"`
	IsViewed      bool   `json:"is_viewed"`
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
	PatientInfo PatientInfo `json:"patient_info"`
	QuizId      int         `json:"quiz_id"`
	QuizName    string      `json:"quiz_name"`
	UserScore   int         `json:"user_score"`
	IsFailed    bool        `json:"is_failed"`
	PassNum     int         `json:"pass_num"`
	PassTime    time.Time   `json:"pass_time"`
	IsViewed    bool        `json:"is_viewed"`
}

type PatientIdList struct {
	Patients []PatientInfo `json:"patients"`
}

type PatientInfo struct {
	UserId    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	AvatarUrl string `json:"avatar_url"`
}

type MarkResultAsViewedIn struct {
	PatientId string `json:"patient_id"`
	QuizId    int    `json:"quiz_id"`
	PassNum   int    `json:"pass_num"`
}

type MarkResponseAsViewedIn struct {
	PatientId  string `json:"patient_id"`
	QuizId     int    `json:"quiz_id"`
	PassNum    int    `json:"pass_num"`
	ResponseId int    `json:"response_id"`
}
