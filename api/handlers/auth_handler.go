package handlers

import (
	"fiber-boilerplate/pkg/auth"

	"github.com/gofiber/fiber/v3"
)

func LoginHandler(service auth.AuthService) fiber.Handler {
	return func(c fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		token, err := service.Login(username, password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Invalid credentials",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Login successful!",
			"token":   token,
		})
	}
}
