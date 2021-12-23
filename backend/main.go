package main

import (
	"embed"
	"log"

	"github.com/Xanvial/todo-app-go/backend/api"
	"github.com/Xanvial/todo-app-go/backend/datastore"
	"github.com/Xanvial/todo-app-go/backend/util"
)

//go:embed webstatic/*
var htmlData embed.FS

func main() {
	// load the configuration
	config, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// setup the datastore
	dataStore := datastore.NewDBStore(config)

	// setup the server
	server := api.NewServer(config, htmlData, dataStore)

	// start the server with graceful shutdown
	server.StartWithGraceful()
}
