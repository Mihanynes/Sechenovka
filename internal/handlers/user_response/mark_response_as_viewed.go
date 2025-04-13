package user_response

import (
	"Sechenovka/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) MarkResponseAsViewed(c *fiber.Ctx) error {
	var dtoIn MarkResponseAsViewedIn

	if err := json.Unmarshal(c.Body(), &dtoIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.userResponseStorage.UpdateIsViewed(model.UserIdFromString(dtoIn.PatientId), dtoIn.QuizId, dtoIn.PassNum, dtoIn.ResponseId, dtoIn.IsViewed)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}
