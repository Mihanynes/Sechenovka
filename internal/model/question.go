package model

type Question struct {
	Text        string   `yaml:"text"`
	ScoreToFail *int     `yaml:"score_to_fail"`
	Options     []Option `yaml:"options"`
}

// Опция ответа на вопрос
type Option struct {
	Answer           string `yaml:"answer"`
	Points           int    `yaml:"points"`
	NextQuestionText string `yaml:"next_question_text,omitempty"` // Ссылка на текст следующего вопроса (если есть)
}
