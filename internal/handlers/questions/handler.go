package questions

import (
	dto "Sechenovka/internal/dto/question"
	"Sechenovka/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type handler struct {
	questionService questionService
	historyStorage  historyStorage
}

func New(questionService questionService, historyStorage historyStorage) *handler {
	return &handler{
		questionService: questionService,
		historyStorage:  historyStorage,
	}
}

func (h *handler) StartQuiz(c *fiber.Ctx) error {
	question, err := h.questionService.GetFirstQuestion()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	out := modelToDto(question)
	out.CorrelationID = uuid.New().String()
	return c.Status(fiber.StatusOK).JSON(out)
}

func (h *handler) GetQuestion(c *fiber.Ctx) error {
	var questionIn dto.QuestionIn
	err := c.BodyParser(&questionIn)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.historyStorage.SaveUserResponse(questionIn.ToUserResponse())
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

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
