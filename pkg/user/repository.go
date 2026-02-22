package user

import (
	"fiber-boilerplate/pkg/entities"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Repository interface {
	GetUsers(c fiber.Ctx) ([]entities.User, error)
	GetUserByEmail(c fiber.Ctx, email string) (*entities.User, error)
	// Define other user-related methods here
}

type userRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	// Return a concrete implementation of Repository
	// This is a simplified example; in a real app, this would be more complex
	return &userRepository{
		DB: db,
	}
}

func (s *userRepository) GetUsers(c fiber.Ctx) ([]entities.User, error) {
	var users []entities.User
	result := s.DB.WithContext(c.Context()).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *userRepository) GetUserByEmail(c fiber.Ctx, email string) (*entities.User, error) {
	var user entities.User
	result := s.DB.WithContext(c.Context()).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
