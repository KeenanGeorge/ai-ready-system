package config

import (
	"os"
	"testing"
)

func TestLoad_DefaultValues(t *testing.T) {
	// Clear any existing environment variables
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("SERVER_HOST")
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("TOKEN_TTL")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected default port 8080, got: %s", cfg.Server.Port)
	}

	if cfg.Server.Host != "localhost" {
		t.Errorf("Expected default host localhost, got: %s", cfg.Server.Host)
	}

	if cfg.Auth.JWTSecret != "default-secret-key" {
		t.Errorf("Expected default JWT secret, got: %s", cfg.Auth.JWTSecret)
	}

	if cfg.Auth.TokenTTL != 60 {
		t.Errorf("Expected default token TTL 60, got: %d", cfg.Auth.TokenTTL)
	}

	if cfg.Database.Host != "localhost" {
		t.Errorf("Expected default DB host localhost, got: %s", cfg.Database.Host)
	}

	if cfg.Database.Port != 5432 {
		t.Errorf("Expected default DB port 5432, got: %d", cfg.Database.Port)
	}
}

func TestLoad_EnvironmentVariables(t *testing.T) {
	// Set environment variables
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("SERVER_HOST", "0.0.0.0")
	os.Setenv("JWT_SECRET", "custom-secret")
	os.Setenv("TOKEN_TTL", "120")
	os.Setenv("DB_HOST", "custom-db-host")
	os.Setenv("DB_PORT", "3306")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if cfg.Server.Port != "9090" {
		t.Errorf("Expected port 9090, got: %s", cfg.Server.Port)
	}

	if cfg.Server.Host != "0.0.0.0" {
		t.Errorf("Expected host 0.0.0.0, got: %s", cfg.Server.Host)
	}

	if cfg.Auth.JWTSecret != "custom-secret" {
		t.Errorf("Expected JWT secret 'custom-secret', got: %s", cfg.Auth.JWTSecret)
	}

	if cfg.Auth.TokenTTL != 120 {
		t.Errorf("Expected token TTL 120, got: %d", cfg.Auth.TokenTTL)
	}

	if cfg.Database.Host != "custom-db-host" {
		t.Errorf("Expected DB host 'custom-db-host', got: %s", cfg.Database.Host)
	}

	if cfg.Database.Port != 3306 {
		t.Errorf("Expected DB port 3306, got: %d", cfg.Database.Port)
	}

	// Clean up
	os.Unsetenv("SERVER_PORT")
	os.Unsetenv("SERVER_HOST")
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("TOKEN_TTL")
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
}

func TestGetServerAddress(t *testing.T) {
	cfg := &Config{
		Server: ServerConfig{
			Host: "localhost",
			Port: "8080",
		},
	}

	address := cfg.GetServerAddress()
	expected := "localhost:8080"

	if address != expected {
		t.Errorf("Expected server address '%s', got: '%s'", expected, address)
	}
}

func TestGetEnvAsInt_InvalidValue(t *testing.T) {
	// Set invalid integer value
	os.Setenv("TEST_INT", "invalid")

	value := getEnvAsInt("TEST_INT", 42)
	if value != 42 {
		t.Errorf("Expected default value 42 for invalid env var, got: %d", value)
	}

	os.Unsetenv("TEST_INT")
}
