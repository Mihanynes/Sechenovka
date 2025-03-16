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

	return c.Status(fiber.StatusOK).JSON(PatientIdList{
		PatientIds: model.ConvertUserIdsToStrings(res),
	})
}
