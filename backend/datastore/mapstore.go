package datastore

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
)

type MapStore struct {
	// mutual exclusion lock.
	m sync.RWMutex
	// key always increment when creating new todo.
	key int
	// data store with key as the id and *model.TodoData as the value.
	data map[int]*model.TodoData
}

func NewMapStore() *MapStore {
	newData := make(map[int]*model.TodoData, 0)
	return &MapStore{
		data: newData,
	}
}

func (ms *MapStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ms.m.Lock()
	defer ms.m.Unlock()

	title := r.FormValue("title")
	ms.key += 1

	todo := &model.TodoData{
		ID:     ms.key,
		Title:  title,
		Status: false,
	}
	ms.data[ms.key] = todo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(copy(todo))
}

func (ms *MapStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	ms.m.RLock()
	defer ms.m.RUnlock()

	// get completed data
	completed := []*model.TodoData{}
	for _, todo := range ms.data {
		if todo.Status {
			completed = append(completed, copy(todo))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

func (ms *MapStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {
	ms.m.RLock()
	defer ms.m.RUnlock()

	// get incompleted data
	incompleted := []*model.TodoData{}
	for _, todo := range ms.data {
		if !todo.Status {
			incompleted = append(incompleted, copy(todo))
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incompleted)
}

func (ms *MapStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ms.m.Lock()
	defer ms.m.Unlock()

	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	status, _ := strconv.ParseBool(r.FormValue("status"))

	for todoID, todo := range ms.data {
		if todoID == ID {
			todo.Status = status
		}
	}
}

func (ms *MapStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ms.m.Lock()
	defer ms.m.Unlock()

	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])
	delete(ms.data, ID)
}

func copy(todo *model.TodoData) *model.TodoData {
	return &model.TodoData{
		ID:     todo.ID,
		Title:  todo.Title,
		Status: todo.Status,
	}
}
