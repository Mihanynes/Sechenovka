package questions

import (
	"Sechenovka/internal/models"
	"Sechenovka/internal/queue"
	"Sechenovka/internal/service/history_saver"
)

type questionService interface {
	GetFirstQuestion() (*models.Question, error)
	GetOptionsByQuestionText(questionText string) (*models.Question, error)
}

type historyQueue interface {
	Add(item queue.Message)
}

type historyStorage interface {
	SaveUserResponse(userResponse *history_saver.UserResponse) error
}
