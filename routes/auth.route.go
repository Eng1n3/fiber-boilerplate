package routes

import (
	"fiber-boilerplate/handlers"

	"github.com/gofiber/fiber/v3"
)

func AuthRoute(api fiber.Router) {
	api.Post("/login", handlers.LoginHandler)
}
