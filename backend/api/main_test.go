package api

import (
	"embed"
	"log"
	"os"
	"testing"

	"github.com/Xanvial/todo-app-go/backend/datastore"
	"github.com/Xanvial/todo-app-go/backend/util"
)

func newTestServer(t *testing.T, store datastore.DataStore) *Server {
	// load the configuration
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	server := NewServer(config, embed.FS{}, store)
	return server
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
