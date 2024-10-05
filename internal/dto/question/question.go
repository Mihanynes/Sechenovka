package question

type QuestionIn struct {
	QuestionText string `json:"question_text"`
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
