package auth

import (
	"Sechenovka/internal/model"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *handler) RegisterAdmin(c *fiber.Ctx) error {
	var userIn RegisterAdminIn

	if err := json.Unmarshal(c.Body(), &userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err := userIn.ValidateAdmin()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	_, err = h.authService.Register(&model.User{
		Username:   userIn.Username,
		FirstName:  userIn.FirstName,
		MiddleName: userIn.MiddleName,
		LastName:   userIn.LastName,
		Password:   userIn.Password,
		IsAdmin:    true,
	})

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"message": "user already exists"})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "user successfully registered"})
}

func (h *handler) RegisterUser(c *fiber.Ctx) error {
	var userIn RegisterUserIn

	if err := json.Unmarshal(c.Body(), &userIn); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err := userIn.ValidateUser()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	patientUserId, err := h.authService.Register(DtoUserToModel(userIn))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	doctorUserId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	err = h.doctorPatientStorage.SaveDoctorPatientLink(doctorUserId, *patientUserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "user successfully registered"})
}

func DtoUserToModel(user RegisterUserIn) *model.User {
	return &model.User{
		Username:   user.Username,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Phone:      user.Phone,
		Email:      user.Email,
		Password:   user.Password,
		IsAdmin:    false,
	}
}
