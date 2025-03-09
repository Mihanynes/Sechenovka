package quiz

import "time"

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

type QuizList struct {
	List []Quiz `json:"list"`
}

type Quiz struct {
	QuizId          int       `json:"quiz_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	IsAvailable     bool      `json:"is_available"`
	TimeToPassAgain int64     `json:"time_to_pass_again"`
	NextTimeCan     time.Time `json:"next_time_can,omitempty"`
}
