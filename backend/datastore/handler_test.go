package datastore

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/Xanvial/todo-app-go/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

var app *mux.Router
var testPostgreStore *DBStore
var testMapStore *MapStore

func newDataStoreTest(ds DataStore) DataStore {
	return ds
}

func TestMain(m *testing.M) {
	app = mux.NewRouter()
	testPostgreStore = NewDBStore()
	testMapStore = NewMapStore()

	// register data store (pq or map)
	var store DataStore
	store = testPostgreStore // comment if you want to use mapStore
	store = testMapStore

	app.HandleFunc("/todo/completed", store.GetCompleted).Methods(http.MethodGet)
	app.HandleFunc("/todo/incomplete", store.GetIncomplete).Methods(http.MethodGet)
	app.HandleFunc("/add", store.CreateTodo).Methods(http.MethodPost)
	app.HandleFunc("/update/{id}", store.UpdateTodo).Methods(http.MethodPut)
	app.HandleFunc("/delete/{id}", store.DeleteTodo).Methods(http.MethodDelete)

	os.Exit(m.Run())
}

func truncateTable() {
	testPostgreStore.db.Exec(`TRUNCATE todo RESTART IDENTITY CASCADE`)
	testMapStore.key = 0 // reset key
}

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

func TestUpdateTodo(t *testing.T) {
	// create a todo
	TestCreateTodo(t)

	req, _ := http.NewRequest("PUT", "/update/1", strings.NewReader("status=true"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

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

func TestDeleteTodo(t *testing.T) {
	// create a todo
	TestCreateTodo(t)

	req, _ := http.NewRequest("DELETE", "/delete/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
