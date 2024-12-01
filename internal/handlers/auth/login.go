package auth

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Login(c *fiber.Ctx) error {
	var userIn LoginIn

	if err := json.Unmarshal(c.Body(), &userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	userId, err := h.authService.Login(userIn.Username, userIn.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "signed in successfully", "userId": userId})
}
