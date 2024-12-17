package questions

import (
	"Sechenovka/internal/model"
)

type questionService interface {
	GetFirstUserQuestion(userId model.UserId) (int, *model.Question, error)
	GetQuestionByQuestionId(questionId int) (*model.Question, error)
}
