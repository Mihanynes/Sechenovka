package questions

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetQuestion(c *fiber.Ctx) error {
	var questionIn QuestionIn
	err := json.Unmarshal(c.Body(), &questionIn)
	fmt.Println(questionIn)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if err = questionIn.Validate(); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	question, err := h.questionService.GetOptionsByQuestionId(questionIn.QuestionId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(modelToDto(question))
}
