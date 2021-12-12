package main

import (
	"embed"

	"github.com/Xanvial/todo-app-go/backend/api"
	"github.com/Xanvial/todo-app-go/backend/datastore"
	"github.com/Xanvial/todo-app-go/model"
)

// embed the static web to binary file
//go:embed webstatic/*
var htmlData embed.FS

func main() {
	// define the data store type & create the server
	s := api.NewServer(
		// use the map/arrary/db data store
		datastore.PostgreDataStore,
		// use the embedded html files
		htmlData,
	)

	// run the server on port in config
	s.Start(model.AppPort)
}
