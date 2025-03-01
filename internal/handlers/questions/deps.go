package questions

import (
	"Sechenovka/internal/model"
)

type questionService interface {
	GetFirstUserQuestion(userId model.UserId, quizId int) (int, *model.Question, error)
	GetQuestionByQuestionId(questionId int, quizId int) (*model.Question, error)
}
