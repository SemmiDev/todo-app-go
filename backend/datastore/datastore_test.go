package datastore

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

// app is a global routes variable to be used in tests
var app *mux.Router

// testPostgreStore is a global postgre store to be used in tests
var testPostgreStore *DBStore

// testMapStore is a global map store to be used in tests
var testMapStore *MapStore

// note: array store does not support because index starts from 0 :(

// TestMain is the main function to be executed when running tests
func TestMain(m *testing.M) {
	// set up test database
	app = mux.NewRouter()
	// set up postgre store
	testPostgreStore = NewDBStore()
	// set up map store
	testMapStore = NewMapStore()

	store, storeType := New(ArrayDataStore)
	log.Printf("test running on [%s]\n", storeType)

	// uncomment the following line if you want to switch to other store
	store = testPostgreStore
	store = testMapStore

	// set up routes
	app.HandleFunc("/todo/completed", store.GetCompleted).Methods(http.MethodGet)
	app.HandleFunc("/todo/incomplete", store.GetIncomplete).Methods(http.MethodGet)
	app.HandleFunc("/add", store.CreateTodo).Methods(http.MethodPost)
	app.HandleFunc("/update/{id}", store.UpdateTodo).Methods(http.MethodPut)
	app.HandleFunc("/delete/{id}", store.DeleteTodo).Methods(http.MethodDelete)

	// run tests
	os.Exit(m.Run())
}

// truncateTable truncates the table
func truncateTable() {
	testPostgreStore.db.Exec(`TRUNCATE todo RESTART IDENTITY CASCADE`)
	testMapStore.key = 0 // reset key
}

// TestCreateTodo tests the creation of a todo
func TestCreateTodo(t *testing.T) {
	truncateTable()

	req, _ := http.NewRequest("POST", "/add", strings.NewReader("title=testTodo"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var todo model.TodoData
	err := json.Unmarshal(response.Body.Bytes(), &todo)
	if err != nil {
		t.Errorf("Cannot convert response body to json: %v", err)
	}
	require.Equal(t, int(1), todo.ID)
	require.Equal(t, "testTodo", todo.Title)
	require.Equal(t, false, todo.Status)
}

// TestGetIncompleteTodos tests the retrieval of incomplete todos
func TestGetIncompleteTodos(t *testing.T) {
	// create a todo
	TestCreateTodo(t)

	req, _ := http.NewRequest("GET", "/todo/incomplete", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var todo []model.TodoData
	err := json.Unmarshal(response.Body.Bytes(), &todo)
	if err != nil {
		t.Errorf("Cannot convert response body to json: %v", err)
	}
	require.Equal(t, 1, len(todo))
}

// TestUpdateTodo tests the update of a todo
func TestUpdateTodo(t *testing.T) {
	// create a todo
	TestCreateTodo(t)

	req, _ := http.NewRequest("PUT", "/update/1", strings.NewReader("status=true"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

// TestGetcompletedTodos tests the retrieval of completed todos
func TestGetcompletedTodos(t *testing.T) {
	// create and update a todo
	TestUpdateTodo(t)

	req, _ := http.NewRequest("GET", "/todo/completed", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	var todo []model.TodoData
	err := json.Unmarshal(response.Body.Bytes(), &todo)
	if err != nil {
		t.Errorf("Cannot convert response body to json: %v", err)
	}

	require.Equal(t, 1, len(todo))
	require.Equal(t, 1, todo[0].ID)
	require.Equal(t, "testTodo", todo[0].Title)
	require.Equal(t, true, todo[0].Status)
}

// TestDeleteTodo tests the deletion of a todo
func TestDeleteTodo(t *testing.T) {
	// create a todo
	TestCreateTodo(t)

	req, _ := http.NewRequest("DELETE", "/delete/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

// executeRequest executes a request and returns the response
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	return rr
}

// checkResponseCode checks the response code
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
