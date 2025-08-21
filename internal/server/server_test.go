package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-linear-testmo-demo/internal/config"
)

func TestNewServer(t *testing.T) {
	cfg := &config.Config{}
	server := NewServer(cfg)

	if server == nil {
		t.Fatal("NewServer should not return nil")
	}

	if server.config == nil {
		t.Error("config should be initialized")
	}

	if server.authHandler == nil {
		t.Error("authHandler should be initialized")
	}

	if server.staticHandler == nil {
		t.Error("staticHandler should be initialized")
	}

	if server.mux == nil {
		t.Error("mux should be initialized")
	}
}

func TestServerRoutes(t *testing.T) {
	cfg := &config.Config{}
	server := NewServer(cfg)

	// Test health endpoint
	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	server.GetMux().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 for /health, got: %d", w.Code)
	}

	if w.Body.String() != "ok" {
		t.Errorf("Expected body 'ok' for /health, got: %s", w.Body.String())
	}
}

func TestServerGetMux(t *testing.T) {
	cfg := &config.Config{}
	server := NewServer(cfg)

	mux := server.GetMux()
	if mux == nil {
		t.Error("GetMux should not return nil")
	}
}
