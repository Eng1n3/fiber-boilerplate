package app

import (
	"fiber-boilerplate/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

func App() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(helmet.New())

	app.Route("/auth", routes.AuthRoute).Name("auth")
	app.Route("/user", routes.UserRoute).Name("user")
	return app
}
