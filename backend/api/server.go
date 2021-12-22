package api

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Xanvial/todo-app-go/backend/datastore"
	"github.com/Xanvial/todo-app-go/backend/util"
	"github.com/gorilla/mux"
)

// Server is a struct that holds the server information
type Server struct {
	config    util.Config
	dataStore datastore.DataStore
	router    *mux.Router
}

// NewServer is a function that creates a new server
func NewServer(
	config util.Config,
	static embed.FS,
	dataStore datastore.DataStore) *Server {

	server := &Server{
		config:    config,
		dataStore: dataStore,
	}

	server.setupRouter(static)

	server.setupLoggingMiddleware()
	server.setupRecoveryMiddleware()
	server.setupCORS()

	return server
}

// ping is a function that returns a pong
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

// initRoutes is a function that initializes the routes
func (s *Server) setupRouter(htmlData embed.FS) {
	router := mux.NewRouter()

	// create the endpoint for the ping
	router.HandleFunc("/ping", ping).Methods(http.MethodGet)
	// get completed todo "/todo/completed"
	router.HandleFunc("/todo/completed", s.getCompleted).Methods(http.MethodGet)
	// get incomplete todo "/todo/incomplete"
	router.HandleFunc("/todo/incomplete", s.getIncomplete).Methods(http.MethodGet)
	// add todo
	router.HandleFunc("/add", s.createTodo).Methods(http.MethodPost)
	// update todo status
	router.HandleFunc("/update/{id}", s.updateTodo).Methods(http.MethodPut)
	// delete todo
	router.HandleFunc("/delete/{id}", s.deleteTodo).Methods(http.MethodDelete)

	// server static resource last
	// this assumes main.go is called from root project,
	// change this accordingly, if it's called elsewhere
	serverRoot, err := fs.Sub(htmlData, "webstatic")
	if err != nil {
		log.Fatal(err)
	}

	router.PathPrefix("/").Handler(http.FileServer(http.FS(serverRoot)))
	// if current go doesn't support embed, uncomment this and use instead of embedded implementation above
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("webstatic")))

	s.router = router
}

// Run is a function that runs the server
func (s *Server) Start() {
	log.Printf("[Server] HTTP server is running at port %s", s.config.AppPort)
	err := http.ListenAndServe(s.config.AppPort, s.router)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) StartWithGraceful() {
	log.Printf("[Server] HTTP server is running at port %s with gracefull shutdown", s.config.AppPort)
	var wait time.Duration

	srv := &http.Server{
		Addr:         s.config.AppPort,
		WriteTimeout: s.config.WriteTimeout,
		ReadTimeout:  s.config.ReadTimeout,
		IdleTimeout:  s.config.IdleTimeout,
		Handler:      s.router,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
