package api

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

// Server is a struct that holds the server information
type Server struct {
	router    *mux.Router
	dataStore datastore.DataStore
}

// NewServer is a function that creates a new server
func NewServer(dataStoreType datastore.DataStoreType, htmlData embed.FS) *Server {
	s := &Server{
		router:    mux.NewRouter(),
		dataStore: datastore.New(dataStoreType),
	}
	s.setupMiddleware()
	s.setupRouter(htmlData)

	return s
}

// ping is a function that returns a pong
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

// initRoutes is a function that initializes the routes
func (s *Server) setupRouter(htmlData embed.FS) {
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
func (s *Server) Start(addr string) {
	// Optional, CORS config, to make sure it can be called from everywhere
	headersOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("[Server] HTTP server is running at port %s", addr)
	err := http.ListenAndServe(addr, handlers.CORS(headersOk, methodsOk)(s.router))
	if err != nil {
		log.Fatal(err)
	}
}
