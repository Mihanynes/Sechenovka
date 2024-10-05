package main

import (
	"Sechenovka/config"
	authhandler "Sechenovka/internal/handlers/auth"
	"Sechenovka/internal/handlers/middleware"
	questionshandler "Sechenovka/internal/handlers/questions"
	authservice "Sechenovka/internal/service/auth"
	questionservice "Sechenovka/internal/service/questions"
	"Sechenovka/storage"
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
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, DELETE",
		AllowCredentials: true,
	}))

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	authService := authservice.New(logger)
	authHandler := authhandler.New(authService)

	questionsConfig, err := config.GetQuestionsConfig()
	if err != nil {
		log.Fatal(err)
	}
	questionsService := questionservice.New(questionsConfig)
	questionsHandler := questionshandler.New(questionsService)

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", authHandler.Register)
		router.Post("/login", authHandler.Login)
		router.Get("/logout", middleware.BasicAuth, authHandler.LogoutUser)
	})

	micro.Route("/questions", func(router fiber.Router) {
		router.Get("/question", questionsHandler.GetQuestion)
	})

	log.Fatal(app.Listen(":8080"))
}

func init() {
	storage.ConnectDB()
}
