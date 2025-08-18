package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	fmt.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
