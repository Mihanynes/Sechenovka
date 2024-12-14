package model

type Option struct {
	AnswerText     string `yaml:"answer_text"`
	AnswerID       int    `yaml:"answer_id"`
	QuestionID     int    `yaml:"question_id"`
	Points         int    `yaml:"points"`
	NextQuestionID int    `yaml:"next_question_id,omitempty"`
	IsEnded        bool   `yaml:"is_ended,omitempty"`
}

type Question struct {
	QuestionText string    `yaml:"question_text"`
	QuestionID   int       `yaml:"question_id"`
	ScoreToFail  int       `yaml:"score_to_fail,omitempty"`
	Options      []*Option `yaml:"options"`
}
