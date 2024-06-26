package main

import (
	"Sechenovka/controllers"
	"Sechenovka/initializers"
	"Sechenovka/lib"
	"Sechenovka/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New()
	micro := fiber.New()

	quiz, err := lib.ReadYaml(lib.GetAbsolutePath("database/questions.yaml"))
	if err != nil {
		log.Fatal(err.Error())
	}

	quizController := controllers.NewQuizController(quiz)

	fmt.Println(quiz)

	app.Mount("/api", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controllers.SignUpUser)
		router.Post("/login", controllers.SignInUser)
		router.Get("/logout", middleware.BasicAuth, controllers.LogoutUser)
	})

	micro.Post("/questions", middleware.BasicAuth, quizController.GetQuestion())

	//micro.Get("/users/me", middleware.BasicAuth, controllers.GetMe)
	//
	//micro.All("*", func(c *fiber.Ctx) error {
	//	path := c.Path()
	//	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
	//		"status":  "fail",
	//		"message": fmt.Sprintf("Path: %v does not exists on this server", path),
	//	})
	//})

	log.Fatal(app.Listen(":8000"))
}

func init() {
	initializers.ConnectDB()
}
