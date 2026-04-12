package handlers

import (
	"fiber-boilerplate/api/presenter"
	"fiber-boilerplate/pkg/auth"
	"fiber-boilerplate/pkg/validation"

	"github.com/gofiber/fiber/v3"
)

func RegisterHandler(service auth.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		p := new(validation.Register)
		if err := c.Bind().Body(p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthRegisterFailureResponse(c, err, "Invalid input"))
		}

		err := service.Register(c, p)
		if err != nil {
			return c.Status(err.Code).JSON(presenter.AuthRegisterFailureResponse(c, err, "Registration failed"))
		}

		return c.Status(fiber.StatusCreated).JSON(presenter.AuthRegisterSuccessResponse(c))
	}
}

func LoginHandler(service auth.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		p := new(validation.Login)
		if err := c.Bind().Body(p); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthLoginFailureResponse(c, err, "Invalid credentials"))
		}

		tokens, err := service.Login(c, p)
		if err != nil {
			return c.Status(err.Code).JSON(presenter.AuthLoginFailureResponse(c, nil, err.Error()))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AuthLoginSuccessResponse(c, tokens))
	}
}
