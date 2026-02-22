package handlers

import (
	"errors"
	"fiber-boilerplate/api/presenter"
	"fiber-boilerplate/pkg/auth"
	"fiber-boilerplate/pkg/entities"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func LoginHandler(service auth.Service) fiber.Handler {
	return func(c fiber.Ctx) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		tokens, err := service.Login(c, entities.User{
			Email:    email,
			Password: password,
		})
		fmt.Println(errors.Is(err, gorm.ErrRecordNotFound), 222111)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthFailureResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AuthSuccessResponse(tokens))
	}
}
