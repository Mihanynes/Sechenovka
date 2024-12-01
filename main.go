package main

import (
	"Sechenovka/config"
	"Sechenovka/db"
	authhandler "Sechenovka/internal/handlers/auth"
	"Sechenovka/internal/handlers/middleware"
	questions_handler "Sechenovka/internal/handlers/questions"
	user_response_handler "Sechenovka/internal/handlers/user_response"
	auth_service "Sechenovka/internal/service/auth"
	question_service "Sechenovka/internal/service/question_config"
	user_response_service "Sechenovka/internal/service/user_response"
	"Sechenovka/internal/storage/user"
	user_respons_storage "Sechenovka/internal/storage/user_responses"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"log/slog"
	"os"
)

func main() {
	app := fiber.New()
	micro := fiber.New()
	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, DELETE",
		AllowCredentials: true,
	}))

	db := db.ConnectDB()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	userStorage := user.New(db)
	authService := auth_service.New(userStorage, logger)
	authHandler := authhandler.New(authService)

	initConfig, err := config.GetQuestionsConfig()
	if err != nil {
		log.Fatal(err)
	}
	questionsConfigService := question_service.New(initConfig)
	questionsHandler := questions_handler.New(questionsConfigService)

	userResponseStorage := user_respons_storage.New(db)
	userResponseService := user_response_service.New(userResponseStorage, questionsConfigService)
	userResponseHandler := user_response_handler.New(userResponseService, userResponseStorage)

	middleware := middleware.New(db)

	micro.Static("/", "./public/images/img.png")
	micro.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Добро пожаловать! Перейдите по /api/, чтобы увидеть изображения.")
	})
	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", authHandler.Register)
		router.Post("/login", authHandler.Login)
		router.Post("/logout", middleware.BasicAuth, authHandler.LogoutUser)
	})
	micro.Route("/questions", func(router fiber.Router) {
		router.Post("/start", questionsHandler.StartQuiz)
		router.Post("/questions", questionsHandler.GetQuestion)
	})
	micro.Route("/response", func(router fiber.Router) {
		router.Post("/save", userResponseHandler.SaveUserResponse)
	})

	log.Fatal(app.Listen(":8080"))
}
