package patient

import (
	"Sechenovka/internal/model"
	"os"
	"path/filepath"
	"strings"

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

func (h *handler) UploadAvatar(c *fiber.Ctx) error {
	userId, err := model.UserIdFromCtx(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error retrieving the file")
	}

	if userId.String() == "" {
		return c.Status(fiber.StatusBadRequest).SendString("User id is empty")
	}

	filename := userId.String()

	if strings.Contains(filename, "../") || strings.Contains(filename, "..\\") {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid file name")
	}

	// Определение базового пути
	basePath, err := os.Getwd()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error getting current directory " + err.Error())
	}

	publicPath := filepath.Join(basePath, "public", "images")

	// Создание директории, если ее нет
	if _, err := os.Stat(publicPath); os.IsNotExist(err) {
		err := os.MkdirAll(publicPath, 0755)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error creating directory: " + err.Error())
		}
	}

	// Добавление .jpg, если нет расширения
	ext := filepath.Ext(filename)
	if ext == "" {
		filename += ".jpg"
	}

	filePath := filepath.Join(publicPath, filename)

	// Создайте файл для сохранения
	outFile, err := os.Create(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating file: " + err.Error())
	}
	defer outFile.Close()

	// Сохраните файл
	if err := c.SaveFile(file, outFile.Name()); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving file: " + err.Error())
	}

	return c.SendString("File uploaded successfully")
}
