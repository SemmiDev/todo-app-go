package datastore

import (
	"log"
	"net/http"
)

// DataStoreType is the type of the datastore
type DataStoreType uint

// variants of DataStoreType
const (
	ArrayDataStore DataStoreType = iota + 1
	MapDataStore
	PostgreDataStore
)

// New creates a new datastore
func New(storeType DataStoreType) DataStore {
	// switch the storeType and return the appropriate datastore
	switch storeType {
	case ArrayDataStore:
		log.Println("Currently Using Array Data Store")
		return NewArrayStore()
	case MapDataStore:
		log.Println("Currently Using Map Data Store")
		return NewMapStore()
	case PostgreDataStore:
		log.Println("Currently Using Postgre SQL Data Store")
		return NewDBStore()
	default:
		log.Println("Currently Using Map Data Store")
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
