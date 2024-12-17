package questions

import (
	"Sechenovka/internal/model"
)

type questionService interface {
	GetFirstQuestion() (*model.Question, error)
	GetQuestionByQuestionId(questionId int) (*model.Question, error)
}
