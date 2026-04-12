package routes

import (
	"fiber-boilerplate/api/handlers"
	"fiber-boilerplate/pkg/auth"

	"github.com/gofiber/fiber/v3"
)

func AuthRouter(app *fiber.App, service auth.Service) {
	router := app.Group("/auth").Name("auth.")
	router.Post("/register", handlers.RegisterHandler(service)).Name("register")
	router.Post("/login", handlers.LoginHandler(service)).Name("login")
}
