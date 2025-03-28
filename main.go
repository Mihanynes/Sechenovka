package main

import (
	"Sechenovka/config"
	"Sechenovka/db"
	authhandler "Sechenovka/internal/handlers/auth"
	"Sechenovka/internal/handlers/middleware"
	"Sechenovka/internal/handlers/patient"
	questions_handler "Sechenovka/internal/handlers/quiz"
	user_response_handler "Sechenovka/internal/handlers/user_response"
	auth_service "Sechenovka/internal/service/auth"
	"Sechenovka/internal/service/patient_info"
	user_response_service "Sechenovka/internal/service/patient_response"
	question_service "Sechenovka/internal/service/quiz"
	"Sechenovka/internal/storage/doctor_patient"
	"Sechenovka/internal/storage/user"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"Sechenovka/internal/storage/user_result"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	httpLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	slog "log/slog"
	"os"
	"runtime/debug"
)

func main() {
	var err error
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))
	defer func() {
		if panicErr := recover(); panicErr != nil {
			logger.Error(fmt.Sprintf("recover panic: %+v\n%s", panicErr, debug.Stack()))

		}

		if err != nil {
			logger.Error(err.Error())
		}
	}()

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
	app.Use(httpLogger.New())

	db := db.ConnectDB()

	initConfig, err := config.GetQuestionsConfig()
	if err != nil {
		logger.Error(err.Error())
	}

	//storage
	userStorage := user.New(db)
	userResultStorage := user_result.New(db)
	doctorPatientStorage := doctor_patient.New(db)
	userResponseStorage := user_respons_storage.New(db)

	authService := auth_service.New(userStorage, logger)

	patientInfoService := patient_info.New(userStorage)
	patientInfoHandler := patient.NewHandler(patientInfoService)

	questionsConfigService := question_service.New(initConfig, userResponseStorage, userResultStorage)
	questionsHandler := questions_handler.New(questionsConfigService)

	authHandler := authhandler.New(authService, doctorPatientStorage)
	userResponseService := user_response_service.New(userResponseStorage, userResultStorage, questionsConfigService)
	userResponseHandler := user_response_handler.New(userResponseService, userResponseStorage, questionsConfigService, doctorPatientStorage, userResultStorage, userStorage)

	middleware := middleware.New(db)

	app.Static("/public/questions", "./public/questions")
	app.Static("/public/avatars", "./public/avatars")

	micro.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Добро пожаловать! Перейдите по /api/, чтобы увидеть изображения.")
	})
	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register/user", middleware.AdminAuth, authHandler.RegisterUser)
		router.Post("/register/admin", authHandler.RegisterAdmin)
		router.Post("/login", authHandler.Login)
	})
	micro.Route("/questions", func(router fiber.Router) {
		router.Post("/start", middleware.UserAuth, questionsHandler.StartQuiz)
		router.Get("/get", middleware.UserAuth, questionsHandler.GetQuestion)
	})
	micro.Route("/user/response", func(router fiber.Router) {
		router.Post("/save", middleware.UserAuth, userResponseHandler.SaveUserResponse)
		router.Get("/get", middleware.UserAuth, userResponseHandler.GetUserResponses)
	})

	micro.Route("/admin", func(router fiber.Router) {
		router.Get("/patient/results", middleware.AdminAuth, userResponseHandler.GetUsersResult)
		router.Get("/patient/result", middleware.AdminAuth, userResponseHandler.GetUserResult)
		router.Get("/patient/list", middleware.AdminAuth, userResponseHandler.GetPatientList)
		router.Get("/patient/info", middleware.AdminAuth, patientInfoHandler.GetPatientInfo)
		router.Get("/quiz/info", middleware.AdminAuth, questionsHandler.GetQuizInfo)
	})

	micro.Route("/user/info", func(router fiber.Router) {
		router.Post("/uploadAvatar", middleware.UserAuth, patientInfoHandler.UploadAvatar)
		router.Get("/get", middleware.UserAuth, patientInfoHandler.GetUserInfo)
	})
	micro.Get("/quiz/list", middleware.UserAuth, questionsHandler.GetQuizListForUser)

	log.Fatal(app.Listen(":8080"))
}
