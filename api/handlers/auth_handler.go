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
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthFailureResponse(err))
		}

		tokens, err := service.Login(c, p)
		if err != nil {
			return c.Status(err.Code).JSON(presenter.AuthFailureResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AuthSuccessResponse(tokens))
	}
}
