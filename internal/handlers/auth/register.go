package auth

import (
	"Sechenovka/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) RegisterAdmin(c *fiber.Ctx) error {
	var userIn RegisterIn

	if err := json.Unmarshal(c.Body(), &userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err := userIn.ValidateAdmin()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err = h.authService.Register(DtoToModel(userIn))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "user successfully registered"})
}

func (h *handler) RegisterUser(c *fiber.Ctx) error {
	var userIn RegisterIn

	if err := json.Unmarshal(c.Body(), &userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err := userIn.ValidateUser()
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
		Username:   user.Username,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Phone:      user.Phone,
		Snils:      user.Snils,
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    user.IsAdmin,
	}
}
