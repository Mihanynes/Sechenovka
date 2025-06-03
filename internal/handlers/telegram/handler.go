package telegram

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

const chatIdKey = "chat_id"

type handler struct {
	userStorage userStorage
	authService authService
}

func NewHandler(userStorage userStorage, authService authService) *handler {
	return &handler{userStorage: userStorage, authService: authService}
}

func (h *handler) SetChatId(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			log.Print(fmt.Errorf("Handler[SetChatId] error: %v", err))
		}
	}()

	username := c.Query("username")
	password := c.Query("password")
	chatId, err := strconv.ParseInt(c.Query("chatId"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid chat_id"})
	}

	user, err := h.authService.Login(username, password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	if !user.IsAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "user is not admin"})
	}

	err = h.userStorage.SetField(chatIdKey, chatId, user.UserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"chat_id": chatId})
}
