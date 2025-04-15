package model

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var ErrUserAlreadyExists = errors.New("User already exists")
var ErrUserNotFound = errors.New("user not found")

type User struct {
	Username   string
	FirstName  string
	MiddleName string
	LastName   string
	Phone      string
	Email      string
	Password   string
	IsAdmin    bool
}

type UserId uuid.UUID

func (u UserId) String() string {
	return uuid.UUID(u).String()
}

func UserIdFromCtx(c *fiber.Ctx) (UserId, error) {
	rowUserId, ok := c.Locals("userId").(string)
	if !ok {
		return UserId{}, errors.New("Failed to get userId from context")
	}
	return UserIdFromString(rowUserId), nil
}

func IsAdminFromCtx(c *fiber.Ctx) (bool, error) {
	isAdmin, ok := c.Locals("isAdmin").(bool)
	if !ok {
		return false, errors.New("Failed to get isAdmin from context")
	}
	return isAdmin, nil
}
func UserIdFromString(s string) UserId {
	res, err := uuid.Parse(s)
	if err != nil {
		return UserId(uuid.Nil)
	}
	return UserId(res)
}

func UserIdsFromStrings(s []string) []UserId {
	res := make([]UserId, len(s))
	for i, user := range s {
		res[i] = UserIdFromString(user)
	}
	return res
}

func ConvertUserIdsToStrings(ids []UserId) []string {
	strIds := make([]string, len(ids))
	for i, id := range ids {
		strIds[i] = id.String()
	}
	return strIds
}
