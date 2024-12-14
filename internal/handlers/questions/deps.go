package questions

import (
	"Sechenovka/internal/model"
)

type questionService interface {
	GetFirstQuestion() (*model.Question, error)
	GetOptionsByQuestionId(questionId int) (*model.Question, error)
}
