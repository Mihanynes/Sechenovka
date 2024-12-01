package user_response

import (
	"Sechenovka/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) SaveUserResponse(c *fiber.Ctx) error {
	var dtoIn SaveUserResponseIn
	err := json.Unmarshal(c.Body(), &dtoIn)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	isFailed, err := h.userResponseService.SaveUserResponse(dtoToModel(dtoIn))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"isFailed": isFailed})
}

func dtoToModel(in SaveUserResponseIn) *model.UserResponse {
	return &model.UserResponse{
		UserId:       model.UserId(in.UserId),
		QuestionText: in.QuestionText,
		Response: model.Response{
			Score: in.ResponseScore,
		},
		CorrelationId: in.CorrelationId,
	}
}
