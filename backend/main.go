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

func main() {
	// create a new router
	router := mux.NewRouter()

	// create the endpoint for the ping
	router.HandleFunc("/ping", ping).Methods(http.MethodGet)

	// define or choose type of data store
	dataStore, storeType := datastore.New(
		datastore.PostgreDataStore,
	)
	fmt.Printf("Currently Using [%s]\n", storeType)

	// get completed todo "/todo/completed"
	router.HandleFunc("/todo/completed", dataStore.GetCompleted).Methods(http.MethodGet)
	// get incomplete todo "/todo/incomplete"
	router.HandleFunc("/todo/incomplete", dataStore.GetIncomplete).Methods(http.MethodGet)
	// add todo
	router.HandleFunc("/add", dataStore.CreateTodo).Methods(http.MethodPost)
	// update todo status
	router.HandleFunc("/update/{id}", dataStore.UpdateTodo).Methods(http.MethodPut)
	// delete todo
	router.HandleFunc("/delete/{id}", dataStore.DeleteTodo).Methods(http.MethodDelete)

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

	// Optional, CORS config, to make sure it can be called from everywhere
	headersOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Println("[Server] HTTP server is running at port 8080")
	err = http.ListenAndServe(":8080", handlers.CORS(headersOk, methodsOk)(router))
	if err != nil {
		log.Fatal(err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
