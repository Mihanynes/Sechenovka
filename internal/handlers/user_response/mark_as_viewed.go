package user_response

import (
	"Sechenovka/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) MarkResultAsViewed(c *fiber.Ctx) error {
	var dtoIn MarkResultAsViewedIn

	if err := json.Unmarshal(c.Body(), &dtoIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.userResultStorage.UpdateIsViewed(model.UserIdFromString(dtoIn.PatientId), dtoIn.QuizId, dtoIn.PassNum)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok"})
}
