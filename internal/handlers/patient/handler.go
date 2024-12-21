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
	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	patientInfo, err := h.patientInfoService.GetPatientInfo(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(patientInfo)
}
