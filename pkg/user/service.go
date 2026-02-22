package user

import (
	"fiber-boilerplate/pkg/entities"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers(c fiber.Ctx) ([]entities.User, error)
	// Define other user-related methods here
}

type userService struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) UserService {
	// Return a concrete implementation of UserService
	// This is a simplified example; in a real app, this would be more complex
	return &userService{
		DB: db,
	}
}

// Implement UserService methods here
func (s *userService) GetUsers(c fiber.Ctx) ([]entities.User, error) {
	var users []entities.User
	// Implement logic to retrieve users from the database
	query := s.DB.WithContext(c.Context()).Order("created_at asc")
	result := query.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
