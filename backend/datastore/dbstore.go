package datastore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Xanvial/todo-app-go/model"
	_ "github.com/lib/pq"
)

type DBStore struct {
	db *sql.DB
}

func NewDBStore() *DBStore {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		model.DBHost, model.DBPort, model.DBUser, model.DBPassword, model.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Successfully connected!")

	return &DBStore{
		db: db,
	}
}

func (ds *DBStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	var completed []model.TodoData

	query := `
		SELECT id, title, status
		FROM todo
		WHERE status = true
	`

	rows, err := ds.db.Query(query)
	if err != nil {
		log.Println("error on getting todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		var data model.TodoData
		if err := rows.Scan(&data.ID, &data.Title, &data.Status); err != nil {
			log.Println("error on getting todo:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		completed = append(completed, data)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

func (ds *DBStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {
}

func (ds *DBStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
}

func (ds *DBStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
}

func (ds *DBStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
}
