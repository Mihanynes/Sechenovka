package main

import (
	"Sechenovka/config"
	"Sechenovka/db"
	authhandler "Sechenovka/internal/handlers/auth"
	"Sechenovka/internal/handlers/middleware"
	"Sechenovka/internal/handlers/patient"
	questions_handler "Sechenovka/internal/handlers/questions"
	user_response_handler "Sechenovka/internal/handlers/user_response"
	auth_service "Sechenovka/internal/service/auth"
	"Sechenovka/internal/service/patient_info"
	user_response_service "Sechenovka/internal/service/patient_response"
	question_service "Sechenovka/internal/service/question_config"
	"Sechenovka/internal/storage/doctor_patient"
	"Sechenovka/internal/storage/user"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"log/slog"
	"os"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, DELETE",
		AllowCredentials: false,
	}))
	micro := fiber.New()
	micro.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, DELETE",
		AllowCredentials: false,
	}))
	app.Mount("/api", micro)
	app.Use(logger.New())

	db := db.ConnectDB()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	initConfig, err := config.GetQuestionsConfig()
	if err != nil {
		log.Fatal(err)
	}

	//storage
	userStorage := user.New(db)
	userResultStorage := user_result.New(db)
	doctorPatientStorage := doctor_patient.New(db)
	userResponseStorage := user_respons_storage.New(db)

	authService := auth_service.New(userStorage, logger)

	patientInfoService := patient_info.New(userStorage)
	patientInfoHandler := patient.NewHandler(patientInfoService)

	questionsConfigService := question_service.New(initConfig, userResponseStorage)
	questionsHandler := questions_handler.New(questionsConfigService)

	authHandler := authhandler.New(authService, doctorPatientStorage)
	userResponseService := user_response_service.New(userResponseStorage, userResultStorage, questionsConfigService)
	userResponseHandler := user_response_handler.New(userResponseService, userResponseStorage, questionsConfigService, doctorPatientStorage, userResultStorage)

	middleware := middleware.New(db)

	app.Static("/public", "./public/images")

	micro.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Добро пожаловать! Перейдите по /api/, чтобы увидеть изображения.")
	})
	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register/user", middleware.AdminAuth, authHandler.RegisterUser)
		router.Post("/register/admin", authHandler.RegisterAdmin)
		router.Post("/login", authHandler.Login)
		//router.Post("/logout", middleware.UserAuth, authHandler.LogoutUser)
	})
	micro.Route("/questions", func(router fiber.Router) {
		router.Post("/start", middleware.UserAuth, questionsHandler.StartQuiz)
		router.Post("/get", middleware.UserAuth, questionsHandler.GetQuestion)
	})
	micro.Route("user/response", func(router fiber.Router) {
		router.Post("/save", middleware.UserAuth, userResponseHandler.SaveUserResponse)
		router.Post("/get", middleware.UserAuth, userResponseHandler.GetUserResponses)
		router.Post("/results", middleware.AdminAuth, userResponseHandler.GetUsersResult)
	})
	micro.Route("/user/info", func(router fiber.Router) {
		router.Post("/patient", middleware.UserAuth, patientInfoHandler.GetPatientInfo)
	})

	log.Fatal(app.Listen(":8080"))
}
