package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()

	// read config

	// should not do this on real application, save the configuration on separate config file
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", model.DBHost, model.DBPort),
		Database: model.DBName,
		User:     model.DBUser,
		Password: model.DBPassword,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	if newVersion != oldVersion {
		log.Println("migrated from version", oldVersion, "to", newVersion)
	} else {
		log.Println("version is", oldVersion)
	}
}

func usage() {
	log.Println(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}
