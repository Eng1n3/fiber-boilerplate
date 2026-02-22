package routes

import (
	"fiber-boilerplate/api/handlers"
	"fiber-boilerplate/pkg/user"

	"github.com/gofiber/fiber/v3"
)

func UserRouter(app *fiber.App, service user.UserService) {
	router := app.Group("/users")
	router.Get("/", handlers.GetUsers(service)).Name("users")
	router.Post("/", handlers.GetUsers(service)).Name("users")
}
