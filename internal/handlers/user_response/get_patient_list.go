package user_response

import (
	"Sechenovka/internal/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func (h *handler) GetPatientList(c *fiber.Ctx) error {
	var err error
	defer func() {
		if err != nil {
			log.Print(fmt.Errorf("Handler[GetPatientList] error: %v", err))
		}
	}()
	adminId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := h.doctorPatientsStorage.GetPatientsIdsByDoctorId(adminId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	patients := make([]PatientInfo, 0, len(res))
	for _, patientId := range res {
		user, err := h.userInfoStorage.GetUserByUserId(patientId)
		if err != nil || user == nil {
			log.Println("error while getting user info", err)
			continue
		}
		patients = append(patients, PatientInfo{
			UserId:    user.UserID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			AvatarUrl: user.UserID + ".png",
		})
	}

	return c.Status(fiber.StatusOK).JSON(PatientIdList{Patients: patients})
}
