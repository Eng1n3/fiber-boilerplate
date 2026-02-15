package routes

import (
	"fiber-boilerplate/handlers"

	"github.com/gofiber/fiber/v3"
)

func UserRoute(api fiber.Router) {
	api.Get("/", handlers.GetUserHandler)
	api.Post("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  c.Status(fiber.StatusOK),
			"message": "Hello from Express-style handlers post!",
		})
	})
}
