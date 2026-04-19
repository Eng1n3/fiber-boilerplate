package presenter

import (
	"github.com/gofiber/fiber/v3"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func AuthRegisterSuccessResponse(c fiber.Ctx) fiber.Map {
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":    true,
		"message":   "Registration successful!",
		"meta":      nil,
		"timestamp": timestamp,
	}
}

func AuthRegisterFailureResponse(c fiber.Ctx, err interface{}, message string) fiber.Map {
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":    false,
		"message":   message,
		"meta":      nil,
		"timestamp": timestamp,
		"errors":    err,
	}
}

func AuthLoginSuccessResponse(c fiber.Ctx, tokens Tokens) fiber.Map {
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":    true,
		"message":   "Login successful!",
		"data":      tokens,
		"meta":      nil,
		"timestamp": timestamp,
	}
}

func AuthLoginFailureResponse(c fiber.Ctx, err interface{}, message string) fiber.Map {
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":    false,
		"message":   message,
		"meta":      nil,
		"timestamp": timestamp,
		"errors":    err,
	}
}
