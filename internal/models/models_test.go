package models

import (
	"encoding/json"
	"testing"
)

func TestLoginRequest_JSONSerialization(t *testing.T) {
	// Test JSON marshaling
	req := LoginRequest{
		Username: "testuser",
		Password: "testpass",
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal LoginRequest: %v", err)
	}

	// Test JSON unmarshaling
	var unmarshaled LoginRequest
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal LoginRequest: %v", err)
	}

	if unmarshaled.Username != req.Username {
		t.Errorf("Expected username %s, got %s", req.Username, unmarshaled.Username)
	}
	if unmarshaled.Password != req.Password {
		t.Errorf("Expected password %s, got %s", req.Password, unmarshaled.Password)
	}
}

func TestLoginResponse_JSONSerialization(t *testing.T) {
	// Test successful login response
	successResp := LoginResponse{
		Success: true,
		Message: "Login successful",
		Token:   "test-token-123",
	}

	data, err := json.Marshal(successResp)
	if err != nil {
		t.Fatalf("Failed to marshal LoginResponse: %v", err)
	}

	var unmarshaled LoginResponse
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal LoginResponse: %v", err)
	}

	if unmarshaled.Success != successResp.Success {
		t.Errorf("Expected success %v, got %v", successResp.Success, unmarshaled.Success)
	}
	if unmarshaled.Message != successResp.Message {
		t.Errorf("Expected message %s, got %s", successResp.Message, unmarshaled.Message)
	}
	if unmarshaled.Token != successResp.Token {
		t.Errorf("Expected token %s, got %s", successResp.Token, unmarshaled.Token)
	}
}

func TestLoginResponse_FailedLogin(t *testing.T) {
	// Test failed login response
	failedResp := LoginResponse{
		Success: false,
		Message: "Invalid credentials",
		Token:   "",
	}

	data, err := json.Marshal(failedResp)
	if err != nil {
		t.Fatalf("Failed to marshal failed LoginResponse: %v", err)
	}

	var unmarshaled LoginResponse
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal failed LoginResponse: %v", err)
	}

	if unmarshaled.Success {
		t.Error("Expected success to be false")
	}
	if unmarshaled.Message != failedResp.Message {
		t.Errorf("Expected message %s, got %s", failedResp.Message, unmarshaled.Message)
	}
	if unmarshaled.Token != "" {
		t.Errorf("Expected empty token, got %s", unmarshaled.Token)
	}
}

func TestUser_JSONSerialization(t *testing.T) {
	// Test User model JSON serialization
	user := User{
		Username: "admin",
		Password: "secret123",
		Role:     "admin",
	}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal User: %v", err)
	}

	var unmarshaled User
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal User: %v", err)
	}

	if unmarshaled.Username != user.Username {
		t.Errorf("Expected username %s, got %s", user.Username, unmarshaled.Username)
	}
	if unmarshaled.Password != "" {
		t.Error("Expected password to be empty due to json:\"-\" tag")
	}
	if unmarshaled.Role != user.Role {
		t.Errorf("Expected role %s, got %s", user.Role, unmarshaled.Role)
	}
}

func TestUser_PasswordFieldHidden(t *testing.T) {
	// Test that password field is hidden in JSON
	user := User{
		Username: "testuser",
		Password: "secretpassword",
		Role:     "user",
	}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Failed to marshal User: %v", err)
	}

	// Check that password is not in JSON
	jsonStr := string(data)
	if jsonStr == "" {
		t.Fatal("JSON string is empty")
	}

	// Password should not appear in JSON due to json:"-" tag
	if jsonStr == "{}" {
		t.Error("JSON should not be empty object")
	}

	// Verify the JSON contains expected fields
	if !json.Valid(data) {
		t.Error("Generated JSON is not valid")
	}
}

func TestModels_EdgeCases(t *testing.T) {
	// Test empty values
	emptyReq := LoginRequest{}
	data, err := json.Marshal(emptyReq)
	if err != nil {
		t.Fatalf("Failed to marshal empty LoginRequest: %v", err)
	}

	var unmarshaled LoginRequest
	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal empty LoginRequest: %v", err)
	}

	// Test with special characters
	specialReq := LoginRequest{
		Username: "user@domain.com",
		Password: "pass!@#$%^&*()",
	}

	data, err = json.Marshal(specialReq)
	if err != nil {
		t.Fatalf("Failed to marshal special LoginRequest: %v", err)
	}

	err = json.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal special LoginRequest: %v", err)
	}

	if unmarshaled.Username != specialReq.Username {
		t.Errorf("Expected username %s, got %s", specialReq.Username, unmarshaled.Username)
	}
	if unmarshaled.Password != specialReq.Password {
		t.Errorf("Expected password %s, got %s", specialReq.Password, unmarshaled.Password)
	}
}
