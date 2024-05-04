package models

type QuestionAnswer struct {
	Question Question
	Answer   []Answer
}
type Question struct {
	QuestionId   int    `json:"question_id"`
	QuestionText string `json:"question_text"`
}

type Answer struct {
	AnswerId     int    `json:"answer_id"`
	AnswerText   string `json:"answer_text"`
	NextQuestion *Question
}

type Quiz struct {
	Questions []QuestionAnswer
}
