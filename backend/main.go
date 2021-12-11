package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/Xanvial/todo-app-go/backend/datastore"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// embed the static web to binary file
//go:embed webstatic/*
var htmlData embed.FS

// Server is a struct that holds the server information
type Server struct {
	router    *mux.Router
	dataStore datastore.DataStore
}

// NewServer is a function that creates a new server
func NewServer(dataStoreType datastore.DataStoreType) *Server {
	return &Server{
		router:    mux.NewRouter(),
		dataStore: datastore.New(dataStoreType),
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

// initRoutes is a function that initializes the routes
func (s *Server) initRoutes() {
	// create the endpoint for the ping
	s.router.HandleFunc("/ping", ping).Methods(http.MethodGet)
	// get completed todo "/todo/completed"
	s.router.HandleFunc("/todo/completed", s.dataStore.GetCompleted).Methods(http.MethodGet)
	// get incomplete todo "/todo/incomplete"
	s.router.HandleFunc("/todo/incomplete", s.dataStore.GetIncomplete).Methods(http.MethodGet)
	// add todo
	s.router.HandleFunc("/add", s.dataStore.CreateTodo).Methods(http.MethodPost)
	// update todo status
	s.router.HandleFunc("/update/{id}", s.dataStore.UpdateTodo).Methods(http.MethodPut)
	// delete todo
	s.router.HandleFunc("/delete/{id}", s.dataStore.DeleteTodo).Methods(http.MethodDelete)

	// server static resource last
	// this assumes main.go is called from root project,
	// change this accordingly, if it's called elsewhere
	serverRoot, err := fs.Sub(htmlData, "webstatic")
	if err != nil {
		log.Fatal(err)
	}

	s.router.PathPrefix("/").Handler(http.FileServer(http.FS(serverRoot)))
	// if current go doesn't support embed, uncomment this and use instead of embedded implementation above
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("webstatic")))
}

// Run is a function that runs the server
func (s *Server) Run(addr string) {
	// Optional, CORS config, to make sure it can be called from everywhere
	headersOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("[Server] HTTP server is running at port %s", addr)
	err := http.ListenAndServe(addr, handlers.CORS(headersOk, methodsOk)(s.router))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// define the data store type & create the server
	s := NewServer(datastore.PostgreDataStore)
	// init the routes
	s.initRoutes()
	// run the server on port 8080
	s.Run(":8080")
}
