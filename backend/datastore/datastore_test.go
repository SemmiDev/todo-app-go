package datastore

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Xanvial/todo-app-go/backend/entity"
	"github.com/Xanvial/todo-app-go/backend/util"
	"github.com/stretchr/testify/require"
)

var (
	dbStoreTest *DBStore
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dbStoreTest = NewDBStore(config)

	os.Exit(m.Run())
}

func createRandomTodo(t *testing.T) *entity.TodoData {
	title := util.RandomString(10)
	ctx := context.Background()

	todo, err := dbStoreTest.CreateTodo(ctx, title)
	require.NoError(t, err)
	require.NotEmpty(t, todo)
	require.Equal(t, title, todo.Title)

	return todo
}

func createRandomTodoAndTrue(t *testing.T) *entity.TodoData {
	ctx := context.Background()

	todo := createRandomTodo(t)
	err := dbStoreTest.UpdateTodo(ctx, todo.ID, true)
	require.NoError(t, err)
	require.Nil(t, err)

	return todo
}

func TestCreateTodo(t *testing.T) {
	createRandomTodo(t)
	dbStoreTest.truncateTable()
}

func TestGetIncomplete(t *testing.T) {
	todo1 := createRandomTodo(t)
	ctx := context.Background()

	todo2, err := dbStoreTest.GetIncomplete(ctx)
	require.NoError(t, err)
	require.NotNil(t, todo2)
	require.Equal(t, todo1.Title, todo2[0].Title)

	dbStoreTest.truncateTable()

	n := 10
	for i := 1; i <= n; i++ {
		_ = createRandomTodo(t)
	}
	todo2, err = dbStoreTest.GetIncomplete(ctx)
	require.NoError(t, err)
	require.NotNil(t, todo2)
	require.Equal(t, n, len(todo2))
	dbStoreTest.truncateTable()
}

func TestGetCompleted(t *testing.T) {
	todo1 := createRandomTodoAndTrue(t)
	ctx := context.Background()

	todo2, err := dbStoreTest.GetCompleted(ctx)
	require.NoError(t, err)
	require.NotNil(t, todo2)
	require.Equal(t, todo1.Title, todo2[0].Title)

	dbStoreTest.truncateTable()
}

func TestUpdateTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	ctx := context.Background()

	err := dbStoreTest.UpdateTodo(ctx, todo1.ID, true)
	require.NoError(t, err)
	require.Nil(t, err)

	err = dbStoreTest.UpdateTodo(ctx, todo1.ID, false)
	require.NoError(t, err)
	require.Nil(t, err)

	dbStoreTest.truncateTable()
}

func TestDeleteTodo(t *testing.T) {
	todo1 := createRandomTodo(t)
	ctx := context.Background()

	err := dbStoreTest.DeleteTodo(ctx, todo1.ID)
	require.NoError(t, err)
	require.Nil(t, err)

	todos, err := dbStoreTest.GetIncomplete(ctx)
	require.NoError(t, err)
	require.Nil(t, todos)

	dbStoreTest.truncateTable()
}
