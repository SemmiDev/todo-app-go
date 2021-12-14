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
	s := api.NewServer(datastore.Map, htmlData)
	// run the server on port in config
	s.Start(model.AppPort)
}
