package questions

import (
	dto "Sechenovka/internal/dto/question"
	"Sechenovka/internal/models"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	questionService questionService
	historyQueue    historyQueue
}

func New(questionService questionService, historyQueue historyQueue) *handler {
	return &handler{
		questionService: questionService,
		historyQueue:    historyQueue,
	}
}

func (h *handler) GetQuestion(c *fiber.Ctx) error {
	var questionIn dto.QuestionIn
	err := c.BodyParser(&questionIn)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	h.historyQueue.Add(questionIn.ToUserResponse())

	question, err := h.questionService.GetOptionsByQuestionText(questionIn.QuestionText)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(modelToDto(question))
}

func modelToDto(question *models.Question) dto.QuestionOut {
	var out dto.QuestionOut
	out.QuestionText = question.Text
	for _, option := range question.Options {
		dtoOption := dto.Option{
			Answer:           option.Answer,
			Points:           option.Points,
			NextQuestionText: option.NextQuestionText,
		}
		out.Options = append(out.Options, dtoOption)
	}
	return out
}
