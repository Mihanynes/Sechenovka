package user_response

import (
	"Sechenovka/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

func (h *handler) GetUserResponses(c *fiber.Ctx) error {
	dtoUserId := c.Query("UserId")
	passNumString := c.Query("PassNum")
	quizIdString := c.Query("QuizId")

	var err error
	defer func() {
		if err != nil {
			log.Print(fmt.Errorf("Handler[GetUserInfo] error: %v", err))
		}
	}()

	if dtoUserId == "" || passNumString == "" || quizIdString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no query params"})
	}

	passNum, err := strconv.Atoi(passNumString)
	if err != nil || passNum < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid passNum"})
	}

	quizId, err := strconv.Atoi(quizIdString)
	if err != nil || quizId < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid quizId"})
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

	responses, err := h.userResponseStorage.GetUserResponsesByPassNum(model.UserIdFromString(dtoUserId), passNum, quizId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	dtoOut := GetUserResponsesOutList{}
	for _, response := range responses {
		responseConfig, err := h.questionConfigService.GetOptionByResponseId(response.ResponseId, quizId)
		if err != nil {
			continue
			//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			//	"message": err.Error(),
			//})
		}
		questionConfig, err := h.questionConfigService.GetQuestionByQuestionId(responseConfig.QuestionId, quizId)
		if err != nil {
			continue
			//return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			//	"message": err.Error(),
			//})
		}
		dtoOut.Responses = append(dtoOut.Responses, GetUserResponsesOut{
			QuizId:        quizId,
			QuestionText:  questionConfig.QuestionText,
			AnswerText:    responseConfig.AnswerText,
			ResponseScore: responseConfig.Points,
		})
	}
	return c.Status(fiber.StatusOK).JSON(dtoOut)
}
