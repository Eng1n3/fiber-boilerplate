package handlers

import (
	"github.com/gofiber/fiber/v3"
)

func NotFoundHandler() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Route not found",
		})
	}
}
