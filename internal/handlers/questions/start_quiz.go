package questions

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *handler) StartQuiz(c *fiber.Ctx) error {
	question, err := h.questionService.GetFirstQuestion()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	out := modelToDto(question)
	out.CorrelationID = uuid.New().String()
	return c.Status(fiber.StatusOK).JSON(out)
}

func modelToDto(question *model.Question) QuestionOut {
	var out QuestionOut
	out.QuestionText = question.Text
	for _, option := range question.Options {
		dtoOption := Option{
			Answer:           option.Answer,
			Points:           option.Points,
			NextQuestionText: option.NextQuestionText,
		}
		out.Options = append(out.Options, dtoOption)
	}
	return out
}
