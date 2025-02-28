package user_response

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *handler) GetUserResponses(c *fiber.Ctx) error {
	dtoUserId := c.Query("UserId")
	dtoPassNum, err := strconv.Atoi(c.Query("PassNum"))
	if err != nil {
		return err
	}

	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	isAdmin, err := model.IsAdminFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if userId.String() != dtoUserId && !isAdmin {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "access denied"})
	}

	responses, err := h.userResponseStorage.GetUserResponsesByPassNum(model.UserIdFromString(dtoUserId), dtoPassNum)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	dtoOut := GetUserResponsesOutList{}
	for _, response := range responses {
		responseConfig, err := h.questionConfigService.GetOptionByResponseId(response.ResponseId)
		if err != nil {
			continue
			//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			//	"message": err.Error(),
			//})
		}
		questionConfig, err := h.questionConfigService.GetQuestionByQuestionId(responseConfig.QuestionId)
		if err != nil {
			continue
			//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			//	"message": err.Error(),
			//})
		}
		dtoOut.Responses = append(dtoOut.Responses, GetUserResponsesOut{
			QuestionText:  questionConfig.QuestionText,
			AnswerText:    responseConfig.AnswerText,
			ResponseScore: responseConfig.Points,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dtoOut)
}
