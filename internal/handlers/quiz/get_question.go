package quiz

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (h *handler) GetQuestion(c *fiber.Ctx) error {
	questionIdString := c.Query("QuestionId")
	quizIdString := c.Query("QuizId")

	var err error
	defer func() {
		if err != nil {
			log.Print(fmt.Errorf("Handler[GetUserInfo] error: %v", err))
		}
	}()

	if questionIdString == "" || quizIdString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "QuestionId and QuizId are required"})
	}

	questionId, err := strconv.Atoi(questionIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid QuestionId"})
	}

	quizId, err := strconv.Atoi(quizIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid QuizId"})
	}

	question, err := h.questionService.GetQuestionByQuestionId(questionId, quizId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(modelToDto(question))
}
