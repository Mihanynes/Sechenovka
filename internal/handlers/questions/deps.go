package questions

import "Sechenovka/internal/models"

type questionService interface {
	GetOptionsByQuestionText(questionText string) (*models.Question, error)
}
