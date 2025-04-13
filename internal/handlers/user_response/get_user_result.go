package user_response

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetUserResult(c *fiber.Ctx) error {
	var err error

	dtoUserId := c.Query("UserId")
	if dtoUserId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no query params"})
	}

	userId := model.UserIdFromString(dtoUserId)

	adminId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if isPatient := h.doctorPatientsStorage.CheckPatientLinkedToDoctor(adminId, userId); !isPatient {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user is not a patient of doctor who made request"})
	}

	userResults, err := h.userResultStorage.GetUsersResults([]model.UserId{userId})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(h.toDto(userResults))
}
