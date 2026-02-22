package presenter

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func AuthSuccessResponse(tokens Tokens) fiber.Map {
	return fiber.Map{
		"status":  true,
		"message": "Login successful!",
		"data":    tokens,
		"meta": fiber.Map{
			"trace_id":  nil,
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		},
	}
}

func AuthFailureResponse(err error) fiber.Map {
	return fiber.Map{
		"status":  false,
		"message": "Invalid credentials",
		"data":    nil,
		"meta": fiber.Map{
			"trace_id":  nil,
			"timestamp": time.Now().UTC().Format(time.RFC3339),
		},
		"errors": err.Error(),
	}
}
