package quiz

import (
	"Sechenovka/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (h *handler) StartQuiz(c *fiber.Ctx) error {
	quizIdString := c.Query("QuizId")

	var err error
	defer func() {
		if err != nil {
			log.Print(fmt.Errorf("Handler[GetUserInfo] error: %v", err))
		}
	}()

	if quizIdString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "quizId is required"})
	}

	quizId, err := strconv.Atoi(quizIdString)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "quizId is not a number"})
	}

	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	passNum, question, err := h.questionService.GetFirstUserQuestion(userId, quizId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	out := modelToDto(question)
	out.PassNum = passNum

	return c.Status(fiber.StatusOK).JSON(out)
}

func modelToDto(question *model.Question) QuestionOut {
	var out QuestionOut
	out.QuestionText = question.QuestionText
	out.ImgName = question.ImgName
	out.Options = make([]Option, 0)
	out.IsMultipleChoice = question.IsMultipleChoice
	for _, option := range question.Options {
		dtoOption := Option{
			AnswerText:     option.AnswerText,
			AnswerId:       option.AnswerId,
			Points:         option.Points,
			NextQuestionId: option.NextQuestionId,
		}
		out.Options = append(out.Options, dtoOption)
	}
	return out
}
