package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mockdb "github.com/Xanvial/todo-app-go/backend/datastore/mock"
	"github.com/Xanvial/todo-app-go/backend/util"
	"github.com/Xanvial/todo-app-go/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, todo *model.TodoData) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotTodo model.TodoData
	err = json.Unmarshal(data, &gotTodo)

	require.NoError(t, err)
	require.Equal(t, todo.ID, gotTodo.ID)
	require.Equal(t, todo.Title, gotTodo.Title)
	require.Equal(t, todo.Status, gotTodo.Status)
}

func TestCreateTodoAPI(t *testing.T) {
	todo := randomTodo(t)
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
				requireBodyMatchUser(t, recorder.Body, todo)
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

func randomTodo(t *testing.T) *model.TodoData {
	todo := model.TodoData{
		ID:     util.RandomInt(1, 10),
		Title:  util.RandomString(10),
		Status: false,
	}

	return &todo
}
