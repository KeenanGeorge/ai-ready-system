package main

import (
	"log"

	"example.com/go-linear-testmo-demo/internal/config"
	"example.com/go-linear-testmo-demo/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create and start server
	srv := server.NewServer(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
