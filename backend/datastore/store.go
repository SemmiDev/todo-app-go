package datastore

import "net/http"

// DataStoreType is the type of the datastore
type DataStoreType uint

// variants of DataStoreType
const (
	ArrayDataStore DataStoreType = iota + 1
	MapDataStore
	PostgreDataStore
)

// New creates a new datastore
func New(storeType DataStoreType) (DataStore, string) {
	// switch the storeType and return the appropriate datastore
	switch storeType {
	case ArrayDataStore:
		return NewArrayStore(), "Array Data Store"
	case MapDataStore:
		return NewMapStore(), "Map Data Store"
	case PostgreDataStore:
		return NewDBStore(), "Postgre SQL Data Store"
	default:
		return NewMapStore(), "Array Data Store"
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
