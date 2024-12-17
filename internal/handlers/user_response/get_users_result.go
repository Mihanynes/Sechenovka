package user_response

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_result"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) GetUsersResult(c *fiber.Ctx) error {
	//var getUserScoreIn GetUsersResultIn
	//err := json.Unmarshal(c.Body(), &getUserScoreIn)
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error(), "body": string(c.Body())})
	//}
	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	patientIds, err := h.doctorPatientsStorage.GetPatientsIdsByDoctorId(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	userResults, err := h.userResultStorage.GetUsersResults(patientIds)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(toDto(userResults))
}

func toDto(usersResult []user_result.UserResult) GetUsersResultOutList {
	results := make([]GetUsersResultOut, len(usersResult))
	for i, userResult := range usersResult {
		results[i] = GetUsersResultOut{
			UserId:    userResult.UserId,
			FirstName: "",
			LastName:  "",
			UserScore: userResult.Score,
			IsFailed:  false,
		}
	}
	return GetUsersResultOutList{
		UserResults: results,
	}
}
