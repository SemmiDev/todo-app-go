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

// MapStore stores todos in map
type MapStore struct {
	// m is a mutex to ensure thread safety
	m sync.RWMutex
	// key always increment when creating new todo
	key int
	// data store with key as the id and *model.TodoData as the value
	data map[int]*model.TodoData
}

// SetKey for set key to zero (for testing)
func (ms *MapStore) SetKey(i int) {
	ms.key = 0
}

// NewMapStore creates a new map store
func NewMapStore() *MapStore {
	return &MapStore{
		key:  0,
		data: make(map[int]*model.TodoData),
	}
}

// CreateTodo saves the todo to the map store
func (ms *MapStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ms.m.Lock()
	defer ms.m.Unlock()

	title := r.FormValue("title")
	if title == "" {
		log.Println("form title is empty")
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	ms.key += 1

	todo := &model.TodoData{
		ID:     ms.key,
		Title:  title,
		Status: false,
	}
	ms.data[ms.key] = todo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo.Clone())
}

// GetCompleted get todos that are completed
func (ms *MapStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	ms.m.RLock()
	defer ms.m.RUnlock()

	completed := []*model.TodoData{}
	for _, todo := range ms.data {
		if todo.Status {
			completed = append(completed, todo.Clone())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

// GetIncomplete get todos that are incomplete
func (ms *MapStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {
	ms.m.RLock()
	defer ms.m.RUnlock()

	incompleted := []*model.TodoData{}
	for _, todo := range ms.data {
		if !todo.Status {
			incompleted = append(incompleted, todo.Clone())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incompleted)
}

// UpdateTodo updates the todo with the given id
func (ms *MapStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ms.m.Lock()
	defer ms.m.Unlock()

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("error on deleting todo:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	status, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		log.Println("error on updating todo:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for todoID, todo := range ms.data {
		if todoID == ID {
			todo.Status = status
		}
	}
}

// DeleteTodo deletes the todo with the given id
func (ms *MapStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ms.m.Lock()
	defer ms.m.Unlock()

	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("error on deleting todo:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	delete(ms.data, ID)
}
