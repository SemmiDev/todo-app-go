package datastore

import (
	"context"

	"github.com/Xanvial/todo-app-go/backend/entity"
)

// Datastore is the interface that wraps the basic Get, Put and Delete methods.
type DataStore interface {
	GetCompleted(ctx context.Context) ([]*entity.TodoData, error)
	GetIncomplete(ctx context.Context) ([]*entity.TodoData, error)
	CreateTodo(ctx context.Context, title string) (*entity.TodoData, error)
	UpdateTodo(ctx context.Context, ID int, status bool) error
	DeleteTodo(ctx context.Context, ID int) error
}
