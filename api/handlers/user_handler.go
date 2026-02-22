package handlers

import (
	"fiber-boilerplate/pkg/user"

	"github.com/gofiber/fiber/v3"
)

func GetUsers(service user.UserService) fiber.Handler {
	return func(c fiber.Ctx) error {
		users, err := service.GetUsers(c)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Failed to retrieve users",
				"data":    nil,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Users retrieved successfully",
			"data":    users,
		})
	}
}
