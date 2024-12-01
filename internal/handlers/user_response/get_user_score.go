package user_response

import "github.com/gofiber/fiber/v2"

func (h *handler) GetUserScore(c *fiber.Ctx) error {
	var getUserScoreIn GetUserScoreIn
	err := c.BodyParser(&getUserScoreIn)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	score, err := h.userResponseStorage.GetUserTotalScore(getUserScoreIn.UserId, getUserScoreIn.CorrelationId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"score": score})
}
