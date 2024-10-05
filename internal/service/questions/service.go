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

func (s *service) GetFirstQuestion() (*models.Question, error) {
	return s.GetOptionsByQuestionText(firstQuestion)
}

// Получение опций по тексту вопроса
func (s *service) GetOptionsByQuestionText(questionText string) (*models.Question, error) {
	if questionText == "EOF" {
		return &models.Question{Text: "Тест завершен"}, nil
	}
	for _, question := range s.questions {
		if question.Text == questionText {
			return question, nil
		}
	}
	return nil, fmt.Errorf("вопрос с текстом '%s' не найден", questionText)
}
