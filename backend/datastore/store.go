package datastore

import (
	"context"
	"log"

	"github.com/Xanvial/todo-app-go/backend/util"
	"github.com/Xanvial/todo-app-go/model"
)

// DataStoreType is the type of the datastore
type Type uint

// variants of DataStoreType
const (
	Array Type = iota
	Map
	Postgre
)

// New creates a new datastore
func New(config util.Config, datastore Type) DataStore {
	// switch the storeType and return the appropriate datastore
	switch datastore {
	case Array:
		log.Println("[Data Store] App Currently Using Array Data Store")
		return NewArrayStore()
	case Map:
		log.Println("[Data Store] App Currently Using Map Data Store")
		return NewMapStore()
	case Postgre:
		log.Println("[Data Store] App Currently Using Postgre SQL Data Store")
		return NewDBStore(config)
	default:
		log.Println("[Default] App Currently Using Map Data Store")
		return NewMapStore()
	}
}

// Datastore is the interface that wraps the basic Get, Put and Delete methods.
type DataStore interface {
	GetCompleted(ctx context.Context) ([]*model.TodoData, error)
	GetIncomplete(ctx context.Context) ([]*model.TodoData, error)
	CreateTodo(ctx context.Context, title string) (*model.TodoData, error)
	UpdateTodo(ctx context.Context, ID int, status bool) error
	DeleteTodo(ctx context.Context, ID int) error
}
