package config

import (
	"fiber-boilerplate/pkg/validation"

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
	}
}
