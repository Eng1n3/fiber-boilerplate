package handlers

import "github.com/gofiber/fiber/v3"

func GetUserHandler(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Hello from Express-style handlers get!",
		"data": []fiber.Map{{
			"id":   1,
			"name": "John Doe",
		}, {
			"id":   2,
			"name": "Jane Doe",
		}},
	})
}
