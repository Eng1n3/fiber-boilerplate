package auth

import (
	"fiber-boilerplate/api/presenter"
	"fiber-boilerplate/pkg/user"
	"fiber-boilerplate/pkg/validation"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(c fiber.Ctx, v *validation.Login) (presenter.Tokens, *fiber.Error)
}

type service struct {
	userRepo user.Repository
}

func NewService(userRepo user.Repository) Service {
	// Return a concrete implementation of Service
	// This is a simplified example; in a real app, this would be more complex
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Login(c fiber.Ctx, v *validation.Login) (presenter.Tokens, *fiber.Error) {
	// Simplified login logic for demonstration
	fmt.Println(v.Email, 3311)
	user, err := s.userRepo.GetUserByEmail(c, v.Email)
	if err != nil {
		return presenter.Tokens{}, fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// In a real application, you would verify the password here
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password)) != nil {
		return presenter.Tokens{}, fiber.NewError(fiber.StatusBadRequest, "Invalid username or password")
	}
	// Generate tokens (this is just a placeholder, implement your token generation logic)
	tokens := presenter.Tokens{
		AccessToken:  "access_token_placeholder",
		RefreshToken: "refresh_token_placeholder",
	}
	return tokens, nil
}
