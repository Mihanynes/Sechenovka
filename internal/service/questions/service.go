package questions

import (
	"Sechenovka/internal/models"
	"fmt"
)

const firstQuestion = "Есть ли у вас боль или дискомфорт в грудной клетке?"

type service struct {
	questions []*models.Question
}

func New(questions []*models.Question) *service {
	return &service{
		questions: questions,
	}
}

func (s *service) GetOptionsByQuestionText(questionText string) (*models.Question, error) {
	if questionText == "" {
		return s.getOptionsByQuestionText(firstQuestion)
	}
	return s.getOptionsByQuestionText(questionText)
}

// Получение опций по тексту вопроса
func (s *service) getOptionsByQuestionText(questionText string) (*models.Question, error) {
	for _, question := range s.questions {
		if question.Text == questionText {
			return question, nil
		}
	}
	return nil, fmt.Errorf("вопрос с текстом '%s' не найден", questionText)
}
