package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	Username   string
	FirstName  string
	MiddleName string
	LastName   string
	Snils      string
	Email      string
	Password   string
	IsAdmin    bool
}

type UserId uuid.UUID

func (u UserId) String() string {
	return uuid.UUID(u).String()
}

func UserIdFromCtx(c *fiber.Ctx) (UserId, error) {
	rowUserId, ok := c.Locals("userId").(string) // Предполагается, что ID пользователя имеет тип uint
	if !ok {
		return UserId{}, errors.New("Failed to get userId from context")
	}
	return UserIdFromString(rowUserId), nil
}

func UserIdFromString(s string) UserId {
	return UserId(uuid.MustParse(s))
}
