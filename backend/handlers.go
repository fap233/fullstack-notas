package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
	"github.com/google/uuid"
)

type ApiServer struct {
	store map[string]Note
	mu sync.RWMutex
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		store: make(map[string]Note),
	}
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


func (s *ApiServer) CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if note.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	note.ID = uuid.New().String()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	note.UserID = "user-temp-123"

	s.mu.Lock()
	s.store[note.ID] = note
	s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func (s *ApiServer) GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	notes := make([]Note, 0, len(s.store))

	for _, note := range s.store {
		notes = append(notes, note)
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (s *ApiServer) UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing note ID", http.StatusBadRequest)
		return
	}

	var updatedNote Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	existingNote, ok := s.store[id]
	if !ok {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	if updatedNote.Title != "" {
		existingNote.Title = updatedNote.Title
	}
	if updatedNote.Content != "" {
		existingNote.Content = updatedNote.Content
	}
	existingNote.UpdatedAt = time.Now()

	s.store[id] = existingNote

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingNote)
}

func (s *ApiServer) DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing note ID", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.store[id]; !ok {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}
	delete(s.store, id)
	
	w.WriteHeader(http.StatusNoContent)
}
