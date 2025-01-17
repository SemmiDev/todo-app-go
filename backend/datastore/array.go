package datastore

import (
	"context"
	"log"
	"sync"

	"github.com/Xanvial/todo-app-go/backend/entity"
)

// ArrayStore is a datastore that stores data in slice
type ArrayStore struct {
	// m is a mutex to ensure thread safety
	m sync.RWMutex
	// key always increment when creating new todo
	key int
	// data is a slice of todo data
	data []*entity.TodoData
}

// NewArrayStore creates a new ArrayStore
func NewArrayStore() *ArrayStore {
	log.Println("[Data Store] App Currently Using Array Data Store")
	return &ArrayStore{
		data: make([]*entity.TodoData, 0),
	}
}

// GetCompleted returns completed data
func (as *ArrayStore) GetCompleted(ctx context.Context) ([]*entity.TodoData, error) {
	as.m.RLock()
	defer as.m.RUnlock()

	completed := make([]*entity.TodoData, 0)
	for _, todo := range as.data {
		if todo.Status {
			completed = append(completed, todo.Clone())
		}
	}

	return completed, nil
}

// GetIncomplete returns incomplete data
func (as *ArrayStore) GetIncomplete(ctx context.Context) ([]*entity.TodoData, error) {
	as.m.RLock()
	defer as.m.RUnlock()

	// get incomplete data
	incomplete := make([]*entity.TodoData, 0)
	for _, todo := range as.data {
		if !todo.Status {
			incomplete = append(incomplete, todo.Clone())
		}
	}

	return incomplete, nil
}

// CreateTodo creates a new todo
func (as *ArrayStore) CreateTodo(ctx context.Context, title string) (*entity.TodoData, error) {
	as.m.Lock()
	defer as.m.Unlock()

	as.key += 1

	todo := &entity.TodoData{
		ID:     as.key,
		Title:  title,
		Status: false,
	}
	as.data = append(as.data, todo)

	return todo.Clone(), nil
}

// UpdateTodo updates a todo
func (as *ArrayStore) UpdateTodo(ctx context.Context, ID int, status bool) error {
	as.m.Lock()
	defer as.m.Unlock()

	for idx, d := range as.data {
		if d.ID == ID {
			as.data[idx].Status = status
		}
	}

	return nil
}

// DeleteTodo deletes a todo
func (as *ArrayStore) DeleteTodo(ctx context.Context, ID int) error {
	as.m.Lock()
	defer as.m.Unlock()

	for idx, d := range as.data {
		if d.ID == ID {
			as.data = append(as.data[:idx], as.data[idx+1:]...)
		}
	}

	return nil
}
