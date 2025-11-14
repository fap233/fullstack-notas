package main

import (
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	// db *sql.DB
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (s *ApiServer) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"status": "ok",
		"message": "Server is running...",
	}

	json.NewEncoder(w).Encode(response)
}

