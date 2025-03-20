package user_response

import (
	"Sechenovka/config"
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_result"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (h *handler) GetUsersResult(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			log.Print(fmt.Errorf("Handler[GetUsersResult] error: %v", err))
		}
	}()

	adminId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var patientIds []model.UserId

	patientIds, err = h.doctorPatientsStorage.GetPatientsIdsByDoctorId(adminId)
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
		userId := userResult.UserID
		userInfo, err := h.userInfoStorage.GetUserByUserId(model.UserIdFromString(userId))
		if err != nil {
			continue
		}

		results[i] = GetUsersResultOut{
			UserId:    userResult.UserID,
			FirstName: userInfo.FirstName,
			LastName:  userInfo.LastName,
			AvatarUrl: userInfo.UserID + ".png",
			QuizId:    userResult.QuizId,
			QuizName:  config.QuizInfo[userResult.QuizId].Name,
			UserScore: userResult.TotalScore,
			IsFailed:  userResult.IsFailed,
			PassNum:   userResult.PassNum,
			PassTime:  userResult.UpdatedAt,
		}
	}
	return GetUsersResultOutList{
		UserResults: results,
	}
}
