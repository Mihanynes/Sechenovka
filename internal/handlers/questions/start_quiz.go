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
	out.CorrelationId = uuid.New().String()
	return c.Status(fiber.StatusOK).JSON(out)
}

func modelToDto(question *model.Question) QuestionOut {
	var out QuestionOut
	out.QuestionText = question.QuestionText
	for _, option := range question.Options {
		dtoOption := Option{
			AnswerId:       option.AnswerId,
			Points:         option.Points,
			NextQuestionId: option.NextQuestionId,
		}
		out.Options = append(out.Options, dtoOption)
	}
	return out
}
