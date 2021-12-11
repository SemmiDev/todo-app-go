package datastore

import "net/http"

// Datastore is the interface that wraps the basic Get, Put and Delete methods.
type DataStore interface {
	GetCompleted(w http.ResponseWriter, r *http.Request)
	GetIncomplete(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}
