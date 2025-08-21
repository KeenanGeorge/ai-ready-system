package services

import (
	"testing"
)

func TestNewAuthService(t *testing.T) {
	service := NewAuthService()
	if service == nil {
		t.Fatal("NewAuthService should not return nil")
	}
}

func TestAuthenticateUser_ValidCredentials(t *testing.T) {
	service := NewAuthService()

	// Test valid admin credentials
	response, err := service.AuthenticateUser("admin", "admin123")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !response.Success {
		t.Errorf("Expected success=true, got: %v", response.Success)
	}

	if response.Message != "Login successful" {
		t.Errorf("Expected message 'Login successful', got: %s", response.Message)
	}

	if response.Token == "" {
		t.Error("Expected token to be generated")
	}

	// Test valid user credentials
	response, err = service.AuthenticateUser("user", "user123")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !response.Success {
		t.Errorf("Expected success=true, got: %v", response.Success)
	}
}

func TestAuthenticateUser_InvalidCredentials(t *testing.T) {
	service := NewAuthService()

	// Test invalid username
	response, err := service.AuthenticateUser("invalid", "password")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if response.Success {
		t.Errorf("Expected success=false, got: %v", response.Success)
	}

	if response.Message != "Invalid username or password" {
		t.Errorf("Expected message 'Invalid username or password', got: %s", response.Message)
	}

	// Test invalid password
	response, err = service.AuthenticateUser("admin", "wrongpassword")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if response.Success {
		t.Errorf("Expected success=false, got: %v", response.Success)
	}
}

func TestAuthenticateUser_EmptyCredentials(t *testing.T) {
	service := NewAuthService()

	// Test empty username
	_, err := service.AuthenticateUser("", "password")
	if err == nil {
		t.Error("Expected error for empty username")
	}

	// Test empty password
	_, err = service.AuthenticateUser("admin", "")
	if err == nil {
		t.Error("Expected error for empty password")
	}
}

func TestValidateToken(t *testing.T) {
	service := NewAuthService()

	// Test valid token format
	valid, err := service.ValidateToken("dummy-token-admin-123")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if !valid {
		t.Error("Expected valid=true for valid token format")
	}

	// Test invalid token
	valid, err = service.ValidateToken("invalid-token")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if valid {
		t.Error("Expected valid=false for invalid token")
	}

	// Test empty token
	_, err = service.ValidateToken("")
	if err == nil {
		t.Error("Expected error for empty token")
	}
}

func TestGenerateToken(t *testing.T) {
	service := NewAuthService()

	// Test token generation
	token, err := service.GenerateToken("admin")
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if token == "" {
		t.Error("Expected token to be generated")
	}

	// Test empty username
	_, err = service.GenerateToken("")
	if err == nil {
		t.Error("Expected error for empty username")
	}
}
