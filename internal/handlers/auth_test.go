package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-linear-testmo-demo/internal/models"
	"example.com/go-linear-testmo-demo/internal/services"
)

func TestNewAuthHandler(t *testing.T) {
	authService := services.NewAuthService()
	handler := NewAuthHandler(authService)

	if handler == nil {
		t.Fatal("NewAuthHandler should not return nil")
	}

	if handler.authService == nil {
		t.Error("authService should be initialized")
	}
}

func TestLoginHandler_ValidCredentials(t *testing.T) {
	authService := services.NewAuthService()
	handler := NewAuthHandler(authService)

	// Create valid login request
	loginReq := models.LoginRequest{
		Username: "admin",
		Password: "admin123",
	}

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.Login(w, req)

	// Check response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got: %d", w.Code)
	}

	var response models.LoginResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
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
}

func TestLoginHandler_InvalidCredentials(t *testing.T) {
	authService := services.NewAuthService()
	handler := NewAuthHandler(authService)

	// Create invalid login request
	loginReq := models.LoginRequest{
		Username: "admin",
		Password: "wrongpassword",
	}

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.Login(w, req)

	// Check response
	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got: %d", w.Code)
	}

	var response models.LoginResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.Success {
		t.Errorf("Expected success=false, got: %v", response.Success)
	}

	if response.Message != "Invalid username or password" {
		t.Errorf("Expected message 'Invalid username or password', got: %s", response.Message)
	}
}

func TestLoginHandler_WrongMethod(t *testing.T) {
	authService := services.NewAuthService()
	handler := NewAuthHandler(authService)

	// Test GET method (should not be allowed)
	req := httptest.NewRequest("GET", "/api/login", nil)
	w := httptest.NewRecorder()
	handler.Login(w, req)

	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got: %d", w.Code)
	}
}

func TestLoginHandler_InvalidJSON(t *testing.T) {
	authService := services.NewAuthService()
	handler := NewAuthHandler(authService)

	// Test invalid JSON
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.Login(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got: %d", w.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	authService := services.NewAuthService()
	handler := NewAuthHandler(authService)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	handler.Health(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got: %d", w.Code)
	}

	if w.Body.String() != "ok" {
		t.Errorf("Expected body 'ok', got: %s", w.Body.String())
	}
}
