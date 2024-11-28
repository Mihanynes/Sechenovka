package question_config

import (
	"Sechenovka/internal/model"
	"fmt"
)

const firstQuestion = "Есть ли у вас боль или дискомфорт в грудной клетке?"

type service struct {
	questions []*model.Question
}

func New(questions []*model.Question) *service {
	return &service{
		questions: questions,
	}
}

func (s *service) GetFirstQuestion() (*model.Question, error) {
	return s.GetOptionsByQuestionText(firstQuestion)
}

// GetOptionsByQuestionText Получение опций ответа по тексту вопроса
func (s *service) GetOptionsByQuestionText(questionText string) (*model.Question, error) {
	if questionText == "EOF" {
		return &model.Question{Text: "Тест завершен"}, nil
	}
	for _, question := range s.questions {
		if question.Text == questionText {
			return question, nil
		}
	}
	return nil, fmt.Errorf("вопрос с текстом '%s' не найден", questionText)
}
