package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	mockdb "github.com/Xanvial/todo-app-go/backend/datastore/mock"
	"github.com/Xanvial/todo-app-go/backend/util"
	"github.com/Xanvial/todo-app-go/model"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func requireBodyMatchTodo(t *testing.T, body *bytes.Buffer, todo *model.TodoData) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotTodo model.TodoData
	err = json.Unmarshal(data, &gotTodo)

	require.NoError(t, err)
	require.Equal(t, todo.ID, gotTodo.ID)
	require.Equal(t, todo.Title, gotTodo.Title)
	require.Equal(t, todo.Status, gotTodo.Status)
}

func requireBodyMatchTodos(t *testing.T, body *bytes.Buffer) int {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var todos []*model.TodoData
	err = json.Unmarshal(data, &todos)

	return len(todos)
}

func TestCreateTodoAPI(t *testing.T) {
	todo := randomTodo(t, true)
	testCases := []struct {
		name          string
		title         string
		FormValue     string
		buildStubs    func(store *mockdb.MockDataStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "Create a todo successfully",
			title:     todo.Title,
			FormValue: "title",
			buildStubs: func(store *mockdb.MockDataStore) {
				store.EXPECT().
					CreateTodo(gomock.Any(), gomock.Eq(todo.Title)).
					Return(todo, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchTodo(t, recorder.Body, todo)
			},
		},
		{
			name:      "create a todo failed bad request",
			title:     todo.Title,
			FormValue: "xxx",
			buildStubs: func(store *mockdb.MockDataStore) {
				// skip because we don't expect any call to the store
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},

		{
			name:      "create a todo failed internal server error",
			title:     todo.Title,
			FormValue: "title",
			buildStubs: func(store *mockdb.MockDataStore) {
				store.EXPECT().
					CreateTodo(gomock.Any(), gomock.Eq(todo.Title)).
					Return(nil, errors.New("internal server errror"))
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
				require.Equal(t, fmt.Sprintln("Internal Server Error"), recorder.Body.String())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockDataStore(ctrl)
			tc.buildStubs(store)

			api := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			formValue := fmt.Sprintf("%s=%s", tc.FormValue, tc.title)
			req, err := http.NewRequest(http.MethodPost, "/add", strings.NewReader(formValue))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			api.createTodo(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetCompleteTodoAPI(t *testing.T) {
	todos := randomTodos(t)
	completedTodos, _ := filterTodos(t, todos)

	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockDataStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "get all todos completed successfully",
			buildStubs: func(store *mockdb.MockDataStore) {
				store.EXPECT().
					GetCompleted(gomock.Any()).
					Return(completedTodos, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				lenght := requireBodyMatchTodos(t, recorder.Body)
				require.Equal(t, lenght, len(completedTodos))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockDataStore(ctrl)
			tc.buildStubs(store)

			api := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, "/todo/completed", nil)
			require.NoError(t, err)

			api.getCompleted(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestGetIncompleteTodoAPI(t *testing.T) {
	todos := randomTodos(t)
	_, incompleteTodos := filterTodos(t, todos)

	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockDataStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "get all todos incomplete successfully",
			buildStubs: func(store *mockdb.MockDataStore) {
				store.EXPECT().
					GetCompleted(gomock.Any()).
					Return(incompleteTodos, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				lenght := requireBodyMatchTodos(t, recorder.Body)
				require.Equal(t, lenght, len(incompleteTodos))
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockDataStore(ctrl)
			tc.buildStubs(store)

			api := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			req, err := http.NewRequest(http.MethodGet, "/todo/incompleted", nil)
			require.NoError(t, err)

			api.getCompleted(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestDeleteTodoAPI(t *testing.T) {
	todo := randomTodo(t, true)

	testCases := []struct {
		name          string
		ID            int
		buildStubs    func(store *mockdb.MockDataStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "delete todo successfully",
			ID:   todo.ID,
			buildStubs: func(store *mockdb.MockDataStore) {
				store.EXPECT().
					DeleteTodo(gomock.Any(), gomock.Eq(todo.ID)).
					Return(nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockDataStore(ctrl)
			tc.buildStubs(store)

			api := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			vars := map[string]string{
				"id": strconv.Itoa(tc.ID),
			}
			url := fmt.Sprintf("/delete/%s", vars["id"])
			req, err := http.NewRequest(http.MethodDelete, url, nil)
			require.NoError(t, err)
			req = mux.SetURLVars(req, vars)

			api.deleteTodo(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func TestUpdateTodoAPI(t *testing.T) {
	todo := randomTodo(t, true)

	testCases := []struct {
		name          string
		ID            int
		status        bool
		formValue     string
		buildStubs    func(store *mockdb.MockDataStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "update todo successfully",
			ID:        todo.ID,
			status:    todo.Status,
			formValue: "status",
			buildStubs: func(store *mockdb.MockDataStore) {
				store.EXPECT().
					UpdateTodo(gomock.Any(), gomock.Eq(todo.ID), gomock.Eq(todo.Status)).
					Return(nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockDataStore(ctrl)
			tc.buildStubs(store)

			api := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			vars := map[string]string{
				"id": strconv.Itoa(tc.ID),
			}
			url := fmt.Sprintf("/update/%s", vars["id"])
			formValue := fmt.Sprintf("%s=%v", tc.formValue, tc.status)
			req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(formValue))
			require.NoError(t, err)
			req = mux.SetURLVars(req, vars)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			api.updateTodo(recorder, req)
			tc.checkResponse(recorder)
		})
	}
}

func randomTodo(t *testing.T, status bool) *model.TodoData {
	todo := model.TodoData{
		ID:     util.RandomInt(1, 10),
		Title:  util.RandomString(10),
		Status: status,
	}

	return &todo
}

func randomTodos(t *testing.T) []*model.TodoData {
	todos := make([]*model.TodoData, 0)
	for i := 1; i <= 10; i++ {
		todos = append(todos, &model.TodoData{
			ID:     i,
			Title:  util.RandomString(10),
			Status: util.RandomBool(),
		})
	}

	return todos
}

func filterTodos(t *testing.T, todos []*model.TodoData) (completed []*model.TodoData, incomplete []*model.TodoData) {
	for _, todo := range todos {
		if todo.Status {
			completed = append(completed, todo)
		} else {
			incomplete = append(incomplete, todo)
		}
	}
	return
}
