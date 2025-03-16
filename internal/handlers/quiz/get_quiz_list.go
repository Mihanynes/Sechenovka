package quiz

import (
	"Sechenovka/config"
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *handler) GetQuizListForUser(c *fiber.Ctx) error {
	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.questionService.GetQuizList(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *handler) GetQuizInfo(c *fiber.Ctx) error {
	quizIdString := c.Query("QuizId")
	if quizIdString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "QuizId required"})
	}

	quizId, err := strconv.Atoi(quizIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(config.QuizInfo[quizId])
}
