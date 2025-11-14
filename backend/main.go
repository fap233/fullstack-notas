package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewApiServer()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", server.HealthHandler)

	port := ":8080"

	log.Printf("Server is running on port %s", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Could not start server:", err)
	}
}
