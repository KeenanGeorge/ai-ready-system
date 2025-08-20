package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

// setupServer configures the server routes
func setupServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
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
