package handlers

import (
	"fiber-boilerplate/api/presenter"
	"fiber-boilerplate/pkg/auth"
	"fiber-boilerplate/pkg/validation"

	"github.com/gofiber/fiber/v3"
)

func LoginHandler(service auth.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		p := new(validation.Login)
		if err := c.Bind().Body(p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthFailureResponse(c, err, "Invalid credentials"))
		}

		tokens, err := service.Login(c, p)
		if err != nil {
			return c.Status(err.Code).JSON(presenter.AuthFailureResponse(c, err, "Invalid credentials"))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AuthSuccessResponse(c, tokens))
	}
}
