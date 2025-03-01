package quiz

import (
	"Sechenovka/internal/model"
)

type quizService interface {
	GetFirstUserQuestion(userId model.UserId, quizId int) (int, *model.Question, error)
	GetQuestionByQuestionId(questionId int, quizId int) (*model.Question, error)
	GetQuizList(userId model.UserId) (QuizList, error)
}
