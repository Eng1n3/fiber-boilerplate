package auth

import (
	"errors"
	"fiber-boilerplate/api/presenter"
	"fiber-boilerplate/pkg/entities"
	"fiber-boilerplate/pkg/user"
	"fiber-boilerplate/pkg/validation"
	"fiber-boilerplate/utils"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrInvalidCredentials        = errors.New("invalid credentials")
	ErrInvalidToken              = errors.New("invalid token")
	ErrExpiredToken              = errors.New("token has expired")
	ErrEmailInUse                = errors.New("email already in use")
)

type Service interface {
	Register(c fiber.Ctx, v *validation.Register) *fiber.Error
	Login(c fiber.Ctx, v *validation.Login) (presenter.Tokens, *fiber.Error)
}

type service struct {
	log                    *logrus.Logger
	userRepo               user.Repository
	jwtPrivateKey          []byte
	jwtRefreshPrivateKey   []byte
	jwtTTLInSeconds        time.Duration
	jwtRefreshTTLInSeconds time.Duration
}

func NewService(userRepo user.Repository, jwtPrivateKey string, jwtRefreshPrivateKey string, jwtTTLInSeconds time.Duration, jwtRefreshTTLInSeconds time.Duration) Service {
	// Return a concrete implementation of Service
	// This is a simplified example; in a real app, this would be more complex
	return &service{
		log:                    utils.Log,
		userRepo:               userRepo,
		jwtPrivateKey:          []byte(jwtPrivateKey),
		jwtRefreshPrivateKey:   []byte(jwtRefreshPrivateKey),
		jwtTTLInSeconds:        jwtTTLInSeconds,
		jwtRefreshTTLInSeconds: jwtRefreshTTLInSeconds,
	}
}

func (s *service) Register(c fiber.Ctx, v *validation.Register) *fiber.Error {
	// Check if email is already in use
	existingUser, err := s.userRepo.GetUserByEmail(c, v.Email)
	if err == nil && existingUser != nil {
		s.log.Errorf("Registration attempt with already used email: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, ErrEmailInUse.Error())
	}

	// Hash the password
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(v.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to hash password")
	}

	// Create a new user entity
	user := &entities.User{
		Username: v.Username,
		Email:    v.Email,
		Password: string(hashedPassword),
	}

	// Save the user to the database
	err = s.userRepo.CreateUser(c, user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create user")
	}

	return nil
}

// generateAccessToken creates a new JWT access token
func (s *service) generateRefreshToken(user *entities.User) (string, error) {
	// Set the expiration time
	expirationTime := time.Now().Add(s.jwtRefreshTTLInSeconds)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"sub":      user.ID.String(), // subject (user ID)
		"username": user.Username,    // custom claim
		"email":    user.Email,       // custom claim
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString(s.jwtRefreshPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// generateAccessToken creates a new JWT access token
func (s *service) generateAccessToken(user *entities.User) (string, error) {
	// Set the expiration time
	expirationTime := time.Now().Add(s.jwtTTLInSeconds)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"sub":      user.ID.String(), // subject (user ID)
		"username": user.Username,    // custom claim
		"email":    user.Email,       // custom claim
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
	}

	// Create the token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString(s.jwtPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken verifies a JWT token and returns the claims
func (s *service) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, ErrInvalidToken.Error())
		}
		return s.jwtPrivateKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fiber.NewError(fiber.StatusUnauthorized, ErrExpiredToken.Error())
		}
		return nil, fiber.NewError(fiber.StatusUnauthorized, ErrInvalidToken.Error())
	}

	// Extract and validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fiber.NewError(fiber.StatusUnauthorized, ErrInvalidToken.Error())
}

func (s *service) Login(c fiber.Ctx, v *validation.Login) (presenter.Tokens, *fiber.Error) {
	// Simplified login logic for demonstration
	user, err := s.userRepo.GetUserByEmail(c, v.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.log.Errorf("Login attempt with non-existent email: %v", err)
		return presenter.Tokens{}, fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	// In a real application, you would verify the password here
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(v.Password)) != nil {
		return presenter.Tokens{}, fiber.NewError(fiber.StatusBadRequest, ErrInvalidUsernameOrPassword.Error())
	}

	accessToken, accessErr := s.generateAccessToken(user)
	if accessErr != nil {
		return presenter.Tokens{}, fiber.NewError(fiber.StatusInternalServerError, accessErr.Error())
	}

	refreshToken, refreshErr := s.generateRefreshToken(user)
	if refreshErr != nil {
		return presenter.Tokens{}, fiber.NewError(fiber.StatusInternalServerError, refreshErr.Error())
	}

	// Generate tokens (this is just a placeholder, implement your token generation logic)
	tokens := presenter.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokens, nil
}
