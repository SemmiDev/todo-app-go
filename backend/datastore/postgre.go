package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Xanvial/todo-app-go/backend/util"
	"github.com/Xanvial/todo-app-go/model"
	_ "github.com/lib/pq"
)

// DBStore is a struct that contains a database abstraction
type DBStore struct {
	db *sql.DB
}

// GetDB return db abstraction for using in api test
func (ds *DBStore) GetDB() *sql.DB {
	return ds.db
}

// NewDBStore creates a new DBStore
func NewDBStore(config util.Config) *DBStore {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[DB] Successfully connected!")

	// setup the connection pool
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return &DBStore{
		db: db,
	}
}

// GetCompleted returns all completed todos
func (ds *DBStore) GetCompleted(ctx context.Context) ([]*model.TodoData, error) {
	var completed []*model.TodoData

	query := `
		SELECT id, title, status
		FROM todo
		WHERE status = true
	`

	rows, err := ds.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.TodoData
		if err := rows.Scan(&data.ID, &data.Title, &data.Status); err != nil {
			return nil, err
		}
		completed = append(completed, &data)
	}

	return completed, nil
}

// GetIncomplete returns all incomplete todos
func (ds *DBStore) GetIncomplete(ctx context.Context) ([]*model.TodoData, error) {
	var incomplete []*model.TodoData

	query := `
		SELECT id, title, status
		FROM todo
		WHERE status = false
	`

	rows, err := ds.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var data model.TodoData
		if err := rows.Scan(&data.ID, &data.Title, &data.Status); err != nil {
			return nil, err
		}
		incomplete = append(incomplete, &data)
	}

	return incomplete, nil
}

// CreateTodo creates a new todo with the given title
func (ds *DBStore) CreateTodo(ctx context.Context, title string) (*model.TodoData, error) {
	query := `
		INSERT INTO todo (title) 
		VALUES ($1) 
		RETURNING id, title, status
	`

	row := ds.db.QueryRowContext(ctx, query, title)
	var todo model.TodoData
	if err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Status,
	); err != nil {
		return nil, err
	}

	return &todo, nil
}

// UpdateTodo updates a todo with the given ID
func (ds *DBStore) UpdateTodo(ctx context.Context, ID int, status bool) error {
	query := `UPDATE todo SET status = $2 WHERE id = $1`

	_, err := ds.db.ExecContext(ctx, query, ID, status)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo deletes a todo with the given ID
func (ds *DBStore) DeleteTodo(ctx context.Context, ID int) error {
	query := `DELETE FROM todo WHERE id = $1`

	_, err := ds.db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
	}
	return nil
}
