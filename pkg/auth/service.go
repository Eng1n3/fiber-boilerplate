package auth

import "fmt"

type AuthService interface {
	Login(username, password string) (string, error)
}

func NewService() AuthService {
	// Return a concrete implementation of AuthService
	// This is a simplified example; in a real app, this would be more complex
	return &authService{}
}

type authService struct{}

func (s *authService) Login(username, password string) (string, error) {
	// Simplified login logic for demonstration
	if username == "admin" && password == "password" {
		return "fake-jwt-token", nil
	}
	return "", fmt.Errorf("invalid credentials")
}
