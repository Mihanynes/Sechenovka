package questions

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/utils/queue"
)

type questionService interface {
	GetFirstQuestion() (*model.Question, error)
	GetOptionsByQuestionText(questionText string) (*model.Question, error)
}

type historyQueue interface {
	Add(item queue.Message)
}

type historyStorage interface {
	GetUserScore(correlationId string) (int, error)
}
