package handlers

import (
	"net/http"
	"strings"
)

// StaticHandler handles serving static files
type StaticHandler struct{}

// NewStaticHandler creates a new instance of StaticHandler
func NewStaticHandler() *StaticHandler {
	return &StaticHandler{}
}

// Serve handles serving static files
func (h *StaticHandler) Serve(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// Serve login page at root
	if path == "/" {
		http.ServeFile(w, r, "static/login.html")
		return
	}

	// Serve other static files
	filePath := "static/" + strings.TrimPrefix(path, "/")
	http.ServeFile(w, r, filePath)
}
