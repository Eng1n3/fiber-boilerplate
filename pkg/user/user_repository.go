package user

import (
	"errors"
	"fiber-boilerplate/pkg/entities"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound = errors.New("User not found!")
)

type Repository interface {
	CreateUser(c fiber.Ctx, user *entities.User) *fiber.Error
	GetUsers(c fiber.Ctx) ([]entities.User, *fiber.Error)
	GetUserByEmail(c fiber.Ctx, email string) (*entities.User, *fiber.Error)
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

func (s *userRepository) CreateUser(c fiber.Ctx, user *entities.User) *fiber.Error {
	result := s.DB.WithContext(c.Context()).Create(user)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}
	return nil
}

func (s *userRepository) GetUsers(c fiber.Ctx) ([]entities.User, *fiber.Error) {
	var users []entities.User
	result := s.DB.WithContext(c.Context()).Find(&users)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return users, nil
}

func (s *userRepository) GetUserByEmail(c fiber.Ctx, email string) (*entities.User, *fiber.Error) {
	var user entities.User
	result := s.DB.WithContext(c.Context()).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fiber.NewError(fiber.StatusNotFound, "User not found!")
	}
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return &user, nil
}
