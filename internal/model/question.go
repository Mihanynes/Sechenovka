package model

type Option struct {
	AnswerText     string `yaml:"answer_text"`
	AnswerId       int    `yaml:"answer_id"`
	QuestionId     int    `yaml:"question_id"`
	Points         int    `yaml:"points"`
	NextQuestionId int    `yaml:"next_question_id,omitempty"`
	IsEnded        bool   `yaml:"is_ended,omitempty"`
}

type Question struct {
	QuestionText string    `yaml:"question_text"`
	QuestionId   int       `yaml:"question_id"`
	ImgName      string    `yaml:"img_name"`
	ScoreToFail  *int      `yaml:"score_to_fail,omitempty"`
	Options      []*Option `yaml:"options"`
}
