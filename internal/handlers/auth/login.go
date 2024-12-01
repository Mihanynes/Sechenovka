package auth

import (
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Login(c *fiber.Ctx) error {
	var userIn LoginIn

	if err := c.BodyParser(&userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err := h.authService.Login(userIn.Username, userIn.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "signed in successfully"})
}
