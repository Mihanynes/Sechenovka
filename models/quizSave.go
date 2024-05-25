package models

type QA struct {
	QuestionText string `yaml:"question_text"`
	AnswerText   string `yaml:"answer_text"`
	AnswerScore  int    `yaml:"answer_score"`
}

type Test struct {
	Timestamp string `yaml:"timestamp"`
	Questions []QA   `yaml:"questions"`
}

type Data struct {
	History []Test `yaml:"history"`
}
