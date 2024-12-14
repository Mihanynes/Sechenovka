package question_config

import (
	"Sechenovka/internal/model"
	"fmt"
)

const firstQuestion = 1

type service struct {
	questions []*model.Question
}

func New(questions []*model.Question) *service {
	return &service{
		questions: questions,
	}
}

func (s *service) GetFirstQuestion() (*model.Question, error) {
	return s.GetOptionsByQuestionId(firstQuestion)
}

// GetOptionsByQuestionText Получение опций ответа по тексту вопроса
func (s *service) GetOptionsByQuestionId(questionId int) (*model.Question, error) {
	for _, question := range s.questions {
		if question.QuestionId == questionId {
			return question, nil
		}
	}
	return nil, fmt.Errorf("вопрос с текстом '%v' не найден", questionId)
}
