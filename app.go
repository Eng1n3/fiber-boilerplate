package app

import (
	"fiber-boilerplate/api/handlers"
	"fiber-boilerplate/api/routes"
	"fiber-boilerplate/database"
	"fiber-boilerplate/middleware"
	"fiber-boilerplate/pkg/auth"
	"fiber-boilerplate/pkg/user"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

func App() *fiber.App {

	db := database.Connect()
	userRepo := user.NewRepository(db)

	app := fiber.New()

	app.Use(middleware.RecoverConfig())
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(helmet.New())

	// app.Route("/auth", routes.AuthRouter).Name("auth")
	authService := auth.NewService(userRepo)
	userService := user.NewService(userRepo)

	routes.AuthRouter(app, authService)
	routes.UserRouter(app, userService)

	// 404 Handler
	app.Use(handlers.NotFoundHandler())
	return app
}
