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
	return c.Status(fiber.StatusOK).JSON(h.toDto(userResults))
}

func (h *handler) toDto(usersResult []user_result.UserResult) GetUsersResultOutList {
	results := make([]GetUsersResultOut, len(usersResult))
	for i, userResult := range usersResult {
		userId := userResult.UserId
		userInfo, err := h.userInfoStorage.GetUserByUserId(model.UserIdFromString(userId))
		if err != nil {
			continue
		}

		results[i] = GetUsersResultOut{
			UserId:    userResult.UserId,
			FirstName: userInfo.FirstName,
			LastName:  userInfo.LastName,
			UserScore: userResult.TotalScore,
			IsFailed:  false,
		}
	}
	return GetUsersResultOutList{
		UserResults: results,
	}
}
