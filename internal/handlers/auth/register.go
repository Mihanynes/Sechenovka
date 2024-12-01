package auth

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) Register(c *fiber.Ctx) error {
	var userIn RegisterIn

	if err := c.BodyParser(&userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err := userIn.Validate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err = h.authService.Register(DtoToModel(userIn))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "user successfully registered"})
}

func DtoToModel(user RegisterIn) *model.User {
	return &model.User{
		Username: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		IsAdmin:  user.IsAdmin,
	}
}
