package presenter

import (
	"github.com/gofiber/fiber/v3"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func AuthRegisterSuccessResponse(c fiber.Ctx) fiber.Map {
	traceID := c.Locals("trace_id")
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":  true,
		"message": "Registration successful!",
		"meta": fiber.Map{
			"trace_id":  traceID,
			"timestamp": timestamp,
		},
	}
}

func AuthRegisterFailureResponse(c fiber.Ctx, err interface{}, message string) fiber.Map {
	traceID := c.Locals("trace_id")
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":  false,
		"message": message,
		"meta": fiber.Map{
			"trace_id":  traceID,
			"timestamp": timestamp,
		},
		"errors": err,
	}
}

func AuthLoginSuccessResponse(c fiber.Ctx, tokens Tokens) fiber.Map {
	traceID := c.Locals("trace_id")
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":  true,
		"message": "Login successful!",
		"data":    tokens,
		"meta": fiber.Map{
			"trace_id":  traceID,
			"timestamp": timestamp,
		},
	}
}

func AuthLoginFailureResponse(c fiber.Ctx, err interface{}, message string) fiber.Map {
	traceID := c.Locals("trace_id")
	timestamp := c.Locals("timestamp")
	return fiber.Map{
		"status":  false,
		"message": message,
		"meta": fiber.Map{
			"trace_id":  traceID,
			"timestamp": timestamp,
		},
		"errors": err,
	}
}
