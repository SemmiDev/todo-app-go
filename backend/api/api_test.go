package api

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"strings"
// 	"testing"

// 	"github.com/Xanvial/todo-app-go/backend/datastore"
// 	"github.com/Xanvial/todo-app-go/backend/util"
// 	"github.com/Xanvial/todo-app-go/model"
// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/require"
// )

// // app is a global routes variable to be used in tests
// var app *mux.Router

// // testPostgreStore is a global postgre store to be used in tests
// var testPostgreStore *datastore.DBStore

// // testMapStore is a global map store to be used in tests
// var testMapStore *datastore.MapStore

// // note: array store does not support because index starts from 0 :(

// // TestMain is the main function to be executed when running tests
// func TestMain(m *testing.M) {
// 	// load config
// 	config, err := util.LoadConfig("../..")
// 	if err != nil {
// 		log.Fatal("cannot load config:", err)
// 	}

// 	dataStore

// 	app = mux.NewRouter()
// 	server := Server{
// 		dataStore: datastore.New(config, datastore.Postgre),
// 		router:    app,
// 	}

// 	// create new router
// 	testPostgreStore = datastore.NewDBStore(config)
// 	testMapStore = datastore.NewMapStore()

// 	// set up routes
// 	app.HandleFunc("/todo/completed", server.getCompleted).Methods(http.MethodGet)
// 	app.HandleFunc("/todo/incomplete", server.getIncomplete).Methods(http.MethodGet)
// 	app.HandleFunc("/add", server.createTodo).Methods(http.MethodPost)
// 	app.HandleFunc("/update/{id}", server.updateTodo).Methods(http.MethodPut)
// 	app.HandleFunc("/delete/{id}", server.deleteTodo).Methods(http.MethodDelete)

// 	app = server.router

// 	// run tests
// 	os.Exit(m.Run())
// }

// // truncateTable truncates the table
// func truncateTable() {
// 	testPostgreStore.GetDB().Exec(`TRUNCATE todo RESTART IDENTITY CASCADE`)
// 	testMapStore.SetKey(0) // reset key
// }

// // TestCreateTodo tests the creation of a todo
// func TestCreateTodo(t *testing.T) {
// 	truncateTable()

// 	req, _ := http.NewRequest("POST", "/add", strings.NewReader("title=testTodo"))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var todo model.TodoData
// 	err := json.Unmarshal(response.Body.Bytes(), &todo)
// 	if err != nil {
// 		t.Errorf("Cannot convert response body to json: %v", err)
// 	}
// 	require.Equal(t, int(1), todo.ID)
// 	require.Equal(t, "testTodo", todo.Title)
// 	require.Equal(t, false, todo.Status)
// }

// // TestCreateTodo tests the creation of a todo: failed
// func TestCreateTodoFailed(t *testing.T) {
// 	truncateTable()

// 	req, _ := http.NewRequest("POST", "/add", strings.NewReader("XXX=testTodo"))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusBadRequest, response.Code)
// }

// // TestGetIncompleteTodos tests the retrieval of incomplete todos
// func TestGetIncompleteTodos(t *testing.T) {
// 	// create a todo
// 	TestCreateTodo(t)

// 	req, _ := http.NewRequest("GET", "/todo/incomplete", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var todo []model.TodoData
// 	err := json.Unmarshal(response.Body.Bytes(), &todo)
// 	if err != nil {
// 		t.Errorf("Cannot convert response body to json: %v", err)
// 	}
// 	require.Equal(t, 1, len(todo))
// }

// // TestUpdateTodo tests the update of a todo
// func TestUpdateTodo(t *testing.T) {
// 	// create a todo
// 	TestCreateTodo(t)

// 	req, _ := http.NewRequest("PUT", "/update/1", strings.NewReader("status=true"))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)
// }

// // TestUpdateTodo tests the update of a todo: failed
// func TestUpdateTodoFailed(t *testing.T) {
// 	// create a todo
// 	TestCreateTodo(t)

// 	req, _ := http.NewRequest("PUT", "/update/XXX", strings.NewReader("status=true"))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusBadRequest, response.Code)
// }

// // TestUpdateTodo tests the update of a todo: failed
// func TestUpdateTodoFailedFormValue(t *testing.T) {
// 	// create a todo
// 	TestCreateTodo(t)

// 	req, _ := http.NewRequest("PUT", "/update/XXX", strings.NewReader("status=XXX"))
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusBadRequest, response.Code)
// }

// // TestGetcompletedTodos tests the retrieval of completed todos
// func TestGetcompletedTodos(t *testing.T) {
// 	// create and update a todo
// 	TestUpdateTodo(t)

// 	req, _ := http.NewRequest("GET", "/todo/completed", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	var todo []model.TodoData
// 	err := json.Unmarshal(response.Body.Bytes(), &todo)
// 	if err != nil {
// 		t.Errorf("Cannot convert response body to json: %v", err)
// 	}

// 	require.Equal(t, 1, len(todo))
// 	require.Equal(t, 1, todo[0].ID)
// 	require.Equal(t, "testTodo", todo[0].Title)
// 	require.Equal(t, true, todo[0].Status)
// }

// // TestDeleteTodo tests the deletion of a todo
// func TestDeleteTodo(t *testing.T) {
// 	// create a todo
// 	TestCreateTodo(t)

// 	req, _ := http.NewRequest("DELETE", "/delete/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)
// }

// // executeRequest executes a request and returns the response
// func executeRequest(req *http.Request) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	app.ServeHTTP(rr, req)

// 	return rr
// }

// // checkResponseCode checks the response code
// func checkResponseCode(t *testing.T, expected, actual int) {
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
// }
