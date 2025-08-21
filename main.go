package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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

// Hardcoded credentials for testing
var validCredentials = map[string]string{
	"admin": "admin123",
	"user":  "user123",
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if credentials are valid
	if password, exists := validCredentials[loginReq.Username]; exists && password == loginReq.Password {
		response := LoginResponse{
			Success: true,
			Message: "Login successful",
			Token:   "dummy-token-" + loginReq.Username,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		response := LoginResponse{
			Success: false,
			Message: "Invalid username or password",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
	}
}

// setupServer configures the server routes
func setupServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/login", loginHandler)

	// Serve static files for the frontend
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "static/login.html")
		} else {
			http.ServeFile(w, r, "static/"+strings.TrimPrefix(r.URL.Path, "/"))
		}
	})

	return mux
}

// startServer starts the server on the specified port
func startServer(port string) error {
	mux := setupServer()
	fmt.Printf("server listening on %s\n", port)
	return http.ListenAndServe(port, mux)
}

func main() {
	if err := startServer(":8080"); err != nil {
		panic(err)
	}
}
