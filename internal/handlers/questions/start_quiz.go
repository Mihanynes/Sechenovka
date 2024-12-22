package questions

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) StartQuiz(c *fiber.Ctx) error {
	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	passNum, question, err := h.questionService.GetFirstUserQuestion(userId)
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
