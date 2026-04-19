package config

import (
	"errors"
	"fiber-boilerplate/pkg/validation"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

// Validator needs to implement the Validate method
func (v *structValidator) Validate(out any) error {

	// Custom password validation logic
	if l, ok := out.(*validation.Login); ok {
		if err := passwordValidation(l.Password); err != nil {
			return err
		}
	}
	return v.validate.Struct(out)
}

func FiberConfig() fiber.Config {
	return fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
		// Override default error handler
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	}
}
