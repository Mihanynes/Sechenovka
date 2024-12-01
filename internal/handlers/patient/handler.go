package patient

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	patientInfoService patientInfoService
}

func NewHandler(patientInfoService patientInfoService) *handler {
	return &handler{patientInfoService: patientInfoService}
}

func (h *handler) GetPatientInfo(c *fiber.Ctx) error {
	var getPatientInfoIn GetPatientInfoIn
	err := c.BodyParser(&getPatientInfoIn)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	patientInfo, err := h.patientInfoService.GetPatientInfo(model.UserId(getPatientInfoIn.UserId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"patient_info": patientInfo})
}
