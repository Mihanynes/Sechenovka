package middleware

import (
	storage "Sechenovka/internal/storage/user"
	"encoding/base64"
	"fmt"
	"gorm.io/gorm"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type middleware struct {
	db *gorm.DB
}

func New(db *gorm.DB) *middleware {
	return &middleware{
		db: db,
	}
}

func (m *middleware) BasicAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		c.Set("WWW-Authenticate", `Basic realm="Restricted"`)
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization required")
	}

	authValue := strings.SplitN(authHeader, " ", 2)
	if len(authValue) != 2 || authValue[0] != "Basic" {
		return c.Status(fiber.StatusUnauthorized).SendString("Authorization header format must be Basic: %v")
	}

	// Декодируем базовое шифрование
	payload, err := base64.StdEncoding.DecodeString(authValue[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid base64 encoding in Authorization header")
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		return c.Status(fiber.StatusUnauthorized).SendString(fmt.Sprintf("Invalid auth header, must have 2 words: %v", authValue))
	}

	var user storage.User
	result := m.db.First(&user, "username = ?", strings.ToLower(pair[0]))
	if result.Error != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pair[1])) != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid username or password")
	}

	// Сохраняем только userId в контекст
	c.Locals("userId", user.UserId)
	return c.Next()
}
