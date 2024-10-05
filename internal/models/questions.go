package models

// Опция ответа на вопрос
type Option struct {
	Answer           string `yaml:"answer"`
	Points           int    `yaml:"points"`
	NextQuestionText string `yaml:"next_question_text,omitempty"` // Ссылка на текст следующего вопроса (если есть)
}

type Question struct {
	Text    string   `yaml:"text"`
	Options []Option `yaml:"options"`
}
