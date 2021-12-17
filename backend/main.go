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
	config, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := api.NewServer(config, htmlData, datastore.Postgre)
	server.Start(config.AppPort)
}
