package datastore

import (
	"context"
	"log"
	"sync"

	"github.com/Xanvial/todo-app-go/backend/entity"
)

// MapStore stores todos in map
type MapStore struct {
	// m is a mutex to ensure thread safety
	m sync.RWMutex
	// key always increment when creating new todo
	key int
	// data store with key as the id and *entity.TodoData as the value
	data map[int]*entity.TodoData
}

// SetKey for set key to zero (for testing)
func (ms *MapStore) SetKey(i int) {
	ms.key = 0
}

// NewMapStore creates a new map store
func NewMapStore() *MapStore {
	log.Println("[Data Store] App Currently Using Map Data Store")
	return &MapStore{
		key:  0,
		data: make(map[int]*entity.TodoData),
	}
}

// GetCompleted get todos that are completed
func (ms *MapStore) GetCompleted(ctx context.Context) ([]*entity.TodoData, error) {
	ms.m.RLock()
	defer ms.m.RUnlock()

	completed := []*entity.TodoData{}
	for _, todo := range ms.data {
		if todo.Status {
			completed = append(completed, todo.Clone())
		}
	}

	return completed, nil
}

// GetIncomplete get todos that are incomplete
func (ms *MapStore) GetIncomplete(ctx context.Context) ([]*entity.TodoData, error) {
	ms.m.RLock()
	defer ms.m.RUnlock()

	incompleted := []*entity.TodoData{}
	for _, todo := range ms.data {
		if !todo.Status {
			incompleted = append(incompleted, todo.Clone())
		}
	}

	return incompleted, nil
}

// CreateTodo saves the todo to the map store
func (ms *MapStore) CreateTodo(ctx context.Context, title string) (*entity.TodoData, error) {
	ms.m.Lock()
	defer ms.m.Unlock()

	ms.key += 1

	todo := &entity.TodoData{
		ID:     ms.key,
		Title:  title,
		Status: false,
	}
	ms.data[ms.key] = todo

	return todo.Clone(), nil
}

// UpdateTodo updates the todo with the given id
func (ms *MapStore) UpdateTodo(ctx context.Context, ID int, status bool) error {
	ms.m.Lock()
	defer ms.m.Unlock()

	for todoID, todo := range ms.data {
		if todoID == ID {
			todo.Status = status
		}
	}
	return nil
}

// DeleteTodo deletes the todo with the given id
func (ms *MapStore) DeleteTodo(ctx context.Context, ID int) error {
	ms.m.Lock()
	defer ms.m.Unlock()

	delete(ms.data, ID)
	return nil
}
