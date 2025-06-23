package quiz

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_result"
	"fmt"
)

const firstQuestion = 1

type Service struct {
	quizConfig          map[int][]*model.Question
	userResponseStorage userResponseStorage
	userResultStorage   *user_result.UserResultStorage
}

func New(quizConfig map[int][]*model.Question, userResponseStorage userResponseStorage, userResultStorage *user_result.UserResultStorage) *Service {
	return &Service{
		quizConfig:          quizConfig,
		userResponseStorage: userResponseStorage,
		userResultStorage:   userResultStorage,
	}
}

func (s *Service) GetFirstUserQuestion(userId model.UserId, quizId int) (int, *model.Question, error) {
	res, err := s.userResponseStorage.GetLastUserResponse(userId, quizId)
	if err != nil {
		return 0, nil, err
	}

	if res == nil {
		// Если последнего ответа нет, возвращаем первый вопрос с passNum 1
		question, err := s.GetQuestionByQuestionId(firstQuestion, quizId)
		if err != nil {
			return 0, nil, err
		}
		return 1, question, nil
	}

	resp, err := s.GetOptionByResponseId(res.ResponseId, quizId)
	if err != nil {
		return 0, nil, err
	}

	question, err := s.GetQuestionByQuestionId(firstQuestion, quizId)
	if err != nil {
		return 0, nil, err
	}

	// Если текущий вопрос завершен досрочно по превышению баллов
	result, err := s.userResultStorage.GetUserResultByQuizId(userId, quizId)
	if err != nil {
		return 0, nil, err
	}
	if result == nil || result.PassNum == res.PassNum {
		return res.PassNum + 1, question, nil
	}

	if !resp.IsEnded {
		// Если текущий вопрос не завершен, возвращаем следующий вопрос с текущим passNum
		nextQuestion, err := s.GetQuestionByQuestionId(resp.NextQuestionId, quizId)
		if err != nil {
			return 0, nil, err
		}
		return res.PassNum, nextQuestion, nil
	}

	// Если текущий вопрос завершен, возвращаем первый вопрос с passNum + 1
	return res.PassNum + 1, question, nil
}

// GetOptionsByQuestionText Получение опций ответа по тексту вопроса
func (s *Service) GetQuestionByQuestionId(questionId int, quizId int) (*model.Question, error) {
	questions, ok := s.quizConfig[quizId]
	if !ok {
		return nil, fmt.Errorf("тест с quizId '%v' не найден", quizId)
	}

	for _, question := range questions {
		if question.QuestionId == questionId {
			return question, nil
		}
	}
	return nil, fmt.Errorf("вопрос с questionId '%v' не найден", questionId)
}

func (s *Service) GetQuestionByResponseId(responseId int, quizId int) (*model.Question, error) {
	questions, ok := s.quizConfig[quizId]
	if !ok {
		return nil, fmt.Errorf("тест с id '%v' не найден", quizId)
	}

	for _, question := range questions {
		for _, option := range question.Options {
			if option.AnswerId == responseId {
				return question, nil
			}
		}
	}
	return nil, fmt.Errorf("ответ с id '%v' не найден", responseId)
}

func (s *Service) GetOptionByResponseId(responseId int, quizId int) (*model.Option, error) {
	questions, ok := s.quizConfig[quizId]
	if !ok {
		return nil, fmt.Errorf("тест с id '%v' не найден", quizId)
	}

	for _, question := range questions {
		for _, option := range question.Options {
			if option.AnswerId == responseId {
				return option, nil
			}
		}
	}
	return nil, fmt.Errorf("ответ с id '%v' не найден", responseId)
}
