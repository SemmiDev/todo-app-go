package datastore

import (
	"log"
	"net/http"
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
func New(datastore Type) DataStore {
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
		return NewDBStore()
	default:
		log.Println("[Default] App Currently Using Map Data Store")
		return NewMapStore()
	}
}

// Datastore is the interface that wraps the basic Get, Put and Delete methods.
type DataStore interface {
	GetCompleted(w http.ResponseWriter, r *http.Request)
	GetIncomplete(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}
