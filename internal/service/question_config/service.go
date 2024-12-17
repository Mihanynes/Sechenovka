package question_config

import (
	"Sechenovka/internal/model"
	"fmt"
)

const firstQuestion = 1

type QuestionConfigService struct {
	questions []*model.Question
}

func New(questions []*model.Question) *QuestionConfigService {
	return &QuestionConfigService{
		questions: questions,
	}
}

func (s *QuestionConfigService) GetFirstQuestion() (*model.Question, error) {
	return s.GetQuestionByQuestionId(firstQuestion)
}

// GetOptionsByQuestionText Получение опций ответа по тексту вопроса
func (s *QuestionConfigService) GetQuestionByQuestionId(questionId int) (*model.Question, error) {
	for _, question := range s.questions {
		if question.QuestionId == questionId {
			return question, nil
		}
	}
	return nil, fmt.Errorf("вопрос с текстом '%v' не найден", questionId)
}

func (s *QuestionConfigService) GetQuestionByResponseId(responseId int) (*model.Question, error) {
	for _, question := range s.questions {
		for _, option := range question.Options {
			if option.AnswerId == responseId {
				return question, nil
			}
		}
	}
	return nil, fmt.Errorf("ответ с id '%v' не найден", responseId)
}

func (s *QuestionConfigService) GetOptionByResponseId(responseId int) (*model.Option, error) {
	for _, question := range s.questions {
		for _, option := range question.Options {
			if option.AnswerId == responseId {
				return option, nil
			}
		}
	}
	return nil, fmt.Errorf("ответ с id '%v' не найден", responseId)
}
