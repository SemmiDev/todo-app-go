package datastore

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
)

// ArrayStore is a datastore that stores data in slice
type ArrayStore struct {
	// m is a mutex to ensure thread safety
	m sync.RWMutex
	// data is a slice of todo data
	data []*model.TodoData
}

// NewArrayStore creates a new ArrayStore
func NewArrayStore() *ArrayStore {
	newData := make([]*model.TodoData, 0)

	return &ArrayStore{
		data: newData,
	}
}

// GetCompleted returns completed data
func (as *ArrayStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	as.m.RLock()
	defer as.m.RUnlock()

	completed := make([]*model.TodoData, 0)
	for _, d := range as.data {
		if d.Status {
			completed = append(completed, copy(d))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

// GetIncomplete returns incomplete data
func (as *ArrayStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {
	as.m.RLock()
	defer as.m.RUnlock()

	// get incomplete data
	incomplete := make([]*model.TodoData, 0)
	for _, d := range as.data {
		if !d.Status {
			incomplete = append(incomplete, copy(d))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incomplete)
}

// CreateTodo creates a new todo
func (as *ArrayStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	as.m.Lock()
	defer as.m.Unlock()

	title := r.FormValue("title")
	if title == "" {
		log.Println("form title is empty")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	as.data = append(as.data, &model.TodoData{
		Title: title,
	})
}

// UpdateTodo updates a todo
func (as *ArrayStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	as.m.Lock()
	defer as.m.Unlock()

	vars := mux.Vars(r)
	title := vars["title"]
	status, _ := strconv.ParseBool(r.FormValue("status"))

	for idx, d := range as.data {
		if d.Title == title {
			as.data[idx].Status = status
		}
	}
}

// DeleteTodo deletes a todo
func (as *ArrayStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	as.m.Lock()
	defer as.m.Unlock()

	vars := mux.Vars(r)
	title := vars["title"]

	for idx, d := range as.data {
		if d.Title == title {
			as.data = append(as.data[:idx], as.data[idx+1:]...)
		}
	}
}
