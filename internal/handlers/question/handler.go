package question

import (
	dto "Sechenovka/internal/dto/question"
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type handler struct {
	questionService questionService
}

func New(questionService questionService) *handler {
	return &handler{
		questionService: questionService,
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
	if err = questionIn.Validate(); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	question, err := h.questionService.GetOptionsByQuestionText(questionIn.QuestionText)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(modelToDto(question))
}

//func (h *handler) GetScore(c *fiber.Ctx) error {
//	var questionIn dto.QuestionIn
//	err := c.BodyParser(&questionIn)
//	if err != nil {
//		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//	}
//	if err = questionIn.ValidateCorrelationId(); err != nil {
//		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
//	}
//
//	correlationId := questionIn.CorrelationId
//	score, err := h.historyStorage.GetUserScore(correlationId)
//	if err != nil {
//		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
//	}
//	return c.Status(fiber.StatusOK).JSON(fiber.Map{"score": score})
//}

func modelToDto(question *model.Question) dto.QuestionOut {
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