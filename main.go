package main

import (
	"Sechenovka/config"
	authhandler "Sechenovka/internal/handlers/auth"
	"Sechenovka/internal/handlers/middleware"
	questionshandler "Sechenovka/internal/handlers/questions"
	"Sechenovka/internal/queue"
	authservice "Sechenovka/internal/service/auth"
	"Sechenovka/internal/service/history_saver"
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

	db := storage.ConnectDB()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	authService := authservice.New(logger, db)
	authHandler := authhandler.New(authService)

	historyStorage := history_saver.NewStorage(db)
	historySaver := history_saver.NewSaver(historyStorage)

	queue := queue.NewProcessQueue(10, historySaver)

	questionsConfig, err := config.GetQuestionsConfig()
	if err != nil {
		log.Fatal(err)
	}
	questionsService := questionservice.New(questionsConfig)
	questionsHandler := questionshandler.New(questionsService, queue)

	middleware := middleware.New(db)

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
