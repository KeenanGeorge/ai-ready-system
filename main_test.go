package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-linear-testmo-demo/internal/config"
	"example.com/go-linear-testmo-demo/internal/models"
	"example.com/go-linear-testmo-demo/internal/server"
)

func TestMain_ConfigurationLoading(t *testing.T) {
	// Test that configuration can be loaded
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	if cfg == nil {
		t.Fatal("Configuration should not be nil")
	}

	// Verify default values
	if cfg.Server.Port != "8080" {
		t.Errorf("Expected default port 8080, got: %s", cfg.Server.Port)
	}
}

func TestMain_ServerCreation(t *testing.T) {
	// Test that server can be created
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	srv := server.NewServer(cfg)
	if srv == nil {
		t.Fatal("Server should not be nil")
	}
}

func TestMain_EndToEndLoginFlow(t *testing.T) {
	// Test the complete login flow through the server
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	srv := server.NewServer(cfg)
	mux := srv.GetMux()

	// Test valid login
	loginReq := models.LoginRequest{
		Username: "admin",
		Password: "admin123",
	}

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	// Verify response
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

func TestMain_HealthEndpoint(t *testing.T) {
	// Test health endpoint
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	srv := server.NewServer(cfg)
	mux := srv.GetMux()

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got: %d", w.Code)
	}

	if w.Body.String() != "ok" {
		t.Errorf("Expected body 'ok', got: %s", w.Body.String())
	}
}

func TestMain_InvalidLogin(t *testing.T) {
	// Test invalid login
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Failed to load configuration: %v", err)
	}

	srv := server.NewServer(cfg)
	mux := srv.GetMux()

	loginReq := models.LoginRequest{
		Username: "admin",
		Password: "wrongpassword",
	}

	body, _ := json.Marshal(loginReq)
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	// Verify response
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
