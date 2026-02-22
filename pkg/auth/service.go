package auth

import (
	"errors"
	"fiber-boilerplate/api/presenter"
	"fiber-boilerplate/pkg/entities"
	"fiber-boilerplate/pkg/user"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(c fiber.Ctx, u entities.User) (presenter.Tokens, error)
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

func (s *service) Login(c fiber.Ctx, u entities.User) (presenter.Tokens, error) {
	// Simplified login logic for demonstration
	user, err := s.userRepo.GetUserByEmail(c, u.Email)
	if err != nil {
		return presenter.Tokens{}, err
	}

	// In a real application, you would verify the password here
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)) != nil {
		return presenter.Tokens{}, errors.New("invalid credentials")
	}
	// Generate tokens (this is just a placeholder, implement your token generation logic)
	tokens := presenter.Tokens{
		AccessToken:  "access_token_placeholder",
		RefreshToken: "refresh_token_placeholder",
	}
	return tokens, nil
}
