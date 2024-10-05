package questions

import (
	"Sechenovka/internal/models"
	"fmt"
)

type service struct {
	questions []*models.Question
}

func New(questions []*models.Question) *service {
	return &service{
		questions: questions,
	}
}

// Получение опций по тексту вопроса
func (s *service) GetOptionsByQuestionText(questionText string) ([]models.Option, error) {
	for _, q := range s.questions {
		if q.Text == questionText {
			return q.Options, nil
		}
	}
	return nil, fmt.Errorf("вопрос с текстом '%s' не найден", questionText)
}
