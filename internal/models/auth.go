package models

// LoginRequest represents the login request structure
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the login response structure
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// User represents a user in the system
type User struct {
	Username string `json:"username"`
	Password string `json:"-"` // "-" means this field won't be serialized to JSON
	Role     string `json:"role"`
}
