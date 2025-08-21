package server

import (
	"fmt"
	"net/http"

	"example.com/go-linear-testmo-demo/internal/config"
	"example.com/go-linear-testmo-demo/internal/handlers"
	"example.com/go-linear-testmo-demo/internal/services"
)

// Server represents the HTTP server
type Server struct {
	config        *config.Config
	authHandler   *handlers.AuthHandler
	staticHandler *handlers.StaticHandler
	mux           *http.ServeMux
}

// NewServer creates a new instance of Server
func NewServer(cfg *config.Config) *Server {
	// Initialize services
	authService := services.NewAuthService()

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	staticHandler := handlers.NewStaticHandler()

	// Create server
	server := &Server{
		config:        cfg,
		authHandler:   authHandler,
		staticHandler: staticHandler,
		mux:           http.NewServeMux(),
	}

	// Setup routes
	server.setupRoutes()

	return server
}

// setupRoutes configures all server routes
func (s *Server) setupRoutes() {
	// Health endpoint
	s.mux.HandleFunc("/health", s.authHandler.Health)

	// Login endpoint
	s.mux.HandleFunc("/api/login", s.authHandler.Login)

	// Static files (catch-all for frontend)
	s.mux.HandleFunc("/", s.staticHandler.Serve)
}

// Start starts the server
func (s *Server) Start() error {
	address := s.config.GetServerAddress()
	fmt.Printf("Server listening on %s\n", address)
	return http.ListenAndServe(address, s.mux)
}

// GetMux returns the underlying ServeMux (useful for testing)
func (s *Server) GetMux() *http.ServeMux {
	return s.mux
}
