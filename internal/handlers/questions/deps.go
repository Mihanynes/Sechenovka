package questions

import (
	"Sechenovka/internal/models"
	"Sechenovka/internal/queue"
)

type questionService interface {
	GetOptionsByQuestionText(questionText string) (*models.Question, error)
}

type historyQueue interface {
	Add(item queue.Message)
}
