package services

import (
	"errors"
	"fmt"
	"time"

	"example.com/go-linear-testmo-demo/internal/models"
)

// AuthService defines the interface for authentication operations
type AuthService interface {
	AuthenticateUser(username, password string) (*models.LoginResponse, error)
	ValidateToken(token string) (bool, error)
	GenerateToken(username string) (string, error)
}

// authService implements the AuthService interface
type authService struct {
	users map[string]models.User
}

// NewAuthService creates a new instance of AuthService
func NewAuthService() AuthService {
	// Initialize with hardcoded users for now
	// In production, this would come from a database
	users := map[string]models.User{
		"admin": {
			Username: "admin",
			Password: "admin123",
			Role:     "admin",
		},
		"user": {
			Username: "user",
			Password: "user123",
			Role:     "user",
		},
	}

	return &authService{
		users: users,
	}
}

// AuthenticateUser validates user credentials and returns login response
func (s *authService) AuthenticateUser(username, password string) (*models.LoginResponse, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password are required")
	}

	user, exists := s.users[username]
	if !exists {
		return &models.LoginResponse{
			Success: false,
			Message: "Invalid username or password",
		}, nil
	}

	if user.Password != password {
		return &models.LoginResponse{
			Success: false,
			Message: "Invalid username or password",
		}, nil
	}

	// Generate token
	token, err := s.GenerateToken(username)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &models.LoginResponse{
		Success: true,
		Message: "Login successful",
		Token:   token,
	}, nil
}

// ValidateToken validates if a token is valid
func (s *authService) ValidateToken(token string) (bool, error) {
	if token == "" {
		return false, errors.New("token is required")
	}

	// For now, just check if token follows expected format
	// In production, this would validate JWT tokens
	if len(token) > 0 && token[:12] == "dummy-token-" {
		return true, nil
	}

	return false, nil
}

// GenerateToken generates a new token for a user
func (s *authService) GenerateToken(username string) (string, error) {
	if username == "" {
		return "", errors.New("username is required")
	}

	// For now, generate a simple token
	// In production, this would create a proper JWT token
	timestamp := time.Now().Unix()
	token := fmt.Sprintf("dummy-token-%s-%d", username, timestamp)

	return token, nil
}
