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

	is_ended, err := h.userResponseService.SaveUserResponse(userId, dtoIn.ResponseId, dtoIn.PassNum)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"is_ended": is_ended})
}
