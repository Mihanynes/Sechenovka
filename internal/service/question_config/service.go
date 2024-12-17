package question_config

import (
	"Sechenovka/internal/model"
	"fmt"
)

const firstQuestion = 1

type QuestionConfigService struct {
	questions           []*model.Question
	userResponseStorage userResponseStorage
}

func New(questions []*model.Question, userResponseStorage userResponseStorage) *QuestionConfigService {
	return &QuestionConfigService{
		questions:           questions,
		userResponseStorage: userResponseStorage,
	}
}

func (s *QuestionConfigService) GetFirstUserQuestion(userId model.UserId) (int, *model.Question, error) {
	res, err := s.userResponseStorage.GetLastUserResponse(userId)
	if err != nil {
		return 0, nil, err
	}

	if res == nil {
		// Если последнего ответа нет, возвращаем первый вопрос с passNum 1
		question, err := s.GetQuestionByQuestionId(firstQuestion)
		if err != nil {
			return 0, nil, err
		}
		return 1, question, nil
	}

	resp, err := s.GetOptionByResponseId(res.ResponseId)
	if err != nil {
		return 0, nil, err
	}

	if !resp.IsEnded {
		// Если текущий вопрос не завершен, возвращаем следующий вопрос с текущим passNum
		nextQuestion, err := s.GetQuestionByQuestionId(resp.NextQuestionId)
		if err != nil {
			return 0, nil, err
		}
		return res.PassNum, nextQuestion, nil
	}

	// Если текущий вопрос завершен, возвращаем первый вопрос с passNum + 1
	question, err := s.GetQuestionByQuestionId(firstQuestion)
	if err != nil {
		return 0, nil, err
	}
	return res.PassNum + 1, question, nil
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
