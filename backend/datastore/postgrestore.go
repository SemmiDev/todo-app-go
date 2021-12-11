package datastore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// DBStore is a struct that contains a database abstraction
type DBStore struct {
	db *sql.DB
}

// NewDBStore creates a new DBStore
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

	// setup the connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DBStore{
		db: db,
	}
}

// GetCompleted returns all completed todos
func (ds *DBStore) GetCompleted(w http.ResponseWriter, r *http.Request) {
	var completed []*model.TodoData

	query := `
		SELECT id, title, status
		FROM todo
		WHERE status = true
	`

	rows, err := ds.db.QueryContext(r.Context(), query)
	if err != nil {
		log.Println("error on getting todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data model.TodoData
		if err := rows.Scan(&data.ID, &data.Title, &data.Status); err != nil {
			log.Println("error on getting todo:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		completed = append(completed, &data)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

// GetIncomplete returns all incomplete todos
func (ds *DBStore) GetIncomplete(w http.ResponseWriter, r *http.Request) {
	var completed []*model.TodoData

	query := `
		SELECT id, title, status
		FROM todo
		WHERE status = false
	`

	rows, err := ds.db.QueryContext(r.Context(), query)
	if err != nil {
		log.Println("error on getting todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data model.TodoData
		if err := rows.Scan(&data.ID, &data.Title, &data.Status); err != nil {
			log.Println("error on getting todo:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		completed = append(completed, &data)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(completed)
}

// CreateTodo creates a new todo with the given title
func (ds *DBStore) CreateTodo(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	query := `
		INSERT INTO todo (title) 
		VALUES ($1) 
		RETURNING id, title, status
	`

	row := ds.db.QueryRowContext(r.Context(), query, title)
	var todo model.TodoData
	if err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Status,
	); err != nil {
		log.Println("error on creating todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todo)
}

// UpdateTodo updates a todo with the given ID
func (ds *DBStore) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("error on deleting todo:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	status, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		log.Println("error on updating todo:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	query := `UPDATE todo SET status = $2 WHERE id = $1`

	_, err = ds.db.ExecContext(r.Context(), query, ID, status)

	if err != nil {
		log.Println("error on updating todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// DeleteTodo deletes a todo with the given ID
func (ds *DBStore) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("error on deleting todo:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM todo WHERE id = $1`

	_, err = ds.db.ExecContext(r.Context(), query, ID)
	if err != nil {
		log.Println("error on updating todo:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
