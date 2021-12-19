package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// getCompleted is a handler that returns completed todos
func (s *Server) getCompleted(w http.ResponseWriter, r *http.Request) {
	completed, err := s.dataStore.GetCompleted(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

// getIncomplete is a handler that returns incomplete todos
func (s *Server) getIncomplete(w http.ResponseWriter, r *http.Request) {
	incomplete, err := s.dataStore.GetIncomplete(r.Context())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incomplete)
}

// addTodo is a handler that adds a todo
func (s *Server) createTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	todo, err := s.dataStore.CreateTodo(r.Context(), title)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// deleteTodo is a handler that update a todo
func (s *Server) updateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	status, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = s.dataStore.UpdateTodo(r.Context(), ID, status)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// deleteTodo is a handler that deletes a todo
func (s *Server) deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = s.dataStore.DeleteTodo(r.Context(), ID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
