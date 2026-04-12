package user

import (
	"fiber-boilerplate/pkg/entities"

	"github.com/gofiber/fiber/v3"
)

type Service interface {
	GetUsers(c fiber.Ctx) ([]entities.User, error)
	GetUserByEmail(c fiber.Ctx, email string) (*entities.User, error)
	// Define other user-related methods here
}

type service struct {
	repository Repository
}

func NewService(userRepo Repository) Service {
	// Return a concrete implementation of Service
	// This is a simplified example; in a real app, this would be more complex
	return &service{
		repository: userRepo,
	}
}

// Implement Service methods here
func (s *service) GetUsers(c fiber.Ctx) ([]entities.User, error) {
	// Implement logic to retrieve users from the database
	return s.repository.GetUsers(c)
}

func (s *service) GetUserByEmail(c fiber.Ctx, email string) (*entities.User, error) {
	return s.repository.GetUserByEmail(c, email)
}
