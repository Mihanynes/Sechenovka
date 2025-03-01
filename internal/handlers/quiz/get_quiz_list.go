package quiz

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetQuizList(c *fiber.Ctx) error {
	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.questionService.GetQuizList(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
