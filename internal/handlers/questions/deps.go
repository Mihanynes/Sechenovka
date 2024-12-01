package questions

import (
	"Sechenovka/internal/model"
)

type questionService interface {
	GetFirstQuestion() (*model.Question, error)
	GetOptionsByQuestionText(questionText string) (*model.Question, error)
}
