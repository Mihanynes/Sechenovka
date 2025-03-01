package user_response

import (
	"Sechenovka/internal/model"
	"Sechenovka/internal/storage/user_result"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *handler) GetUsersResult(c *fiber.Ctx) error {
	quizIdString := c.Query("QuizId")
	if quizIdString == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "QuizId is required"})
	}

	quizId, err := strconv.Atoi(quizIdString)
	if err != nil || quizId <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "QuizId is invalid"})
	}

	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var patientIds []model.UserId
	isAdmin, err := model.IsAdminFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if isAdmin {
		patientIds, err = h.doctorPatientsStorage.GetPatientsIdsByDoctorId(userId)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	} else {
		patientIds = append(patientIds, userId)
	}

	userResults, err := h.userResultStorage.GetUsersResults(patientIds, quizId)
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
