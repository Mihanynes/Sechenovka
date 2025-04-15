package admin

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	userStorage *user.UserStorage
}

func New(userStorage *user.UserStorage) *handler {
	return &handler{userStorage: userStorage}
}

func (h *handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userStorage.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}

	return c.JSON(users)
}

func (h *handler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user ID is required",
		})
	}

	err := h.userStorage.DeleteUserByID(userID)
	if err != nil {
		if errors.Is(err, model.ErrUserNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to delete user",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
