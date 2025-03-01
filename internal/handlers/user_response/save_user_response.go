package user_response

import (
	"Sechenovka/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) SaveUserResponse(c *fiber.Ctx) error {
	var dtoIn SaveUserResponseIn
	err := json.Unmarshal(c.Body(), &dtoIn)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	isEnded, err := h.userResponseService.SaveUserResponses(userId, dtoIn.ResponseIds, dtoIn.PassNum, dtoIn.QuizId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"is_ended": isEnded})
}
