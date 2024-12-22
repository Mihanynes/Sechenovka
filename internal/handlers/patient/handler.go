package patient

import (
	"Sechenovka/internal/model"
	"github.com/gofiber/fiber/v2"
	"os"
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

func (h *handler) UploadAvatar(c *fiber.Ctx) error {
	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error retrieving the file")
	}

	file.Filename = userId.String()

	// Создайте файл для сохранения
	outFile, err := os.Create(file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating file")
	}
	defer outFile.Close()

	// Сохраните файл
	if err := c.SaveFile(file, outFile.Name()); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving file")
	}

	return c.SendString("File uploaded successfully")

}
