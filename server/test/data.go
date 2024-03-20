package test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/emirpasic/gods/sets/hashset"
	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

// begin Todo test data funcs
type TestTodoData struct {
	Todos   []*go_appv1.Todo
	TodoIds *hashset.Set
}

var LoadedTodoTestData = TestTodoData{}

func loadTestTodoData(t *testing.T) {
	deleteAllTodoData(t)
	loadAllTodoData(t)
}

func loadAllTodoData(t *testing.T) {
	loadTodos(t)
}

func deleteAllTodoData(t *testing.T) {
	deleteTodos(t)
}

func deleteTodos(t *testing.T) {
	todos, err := ListTodos(1000, 0, "")
	require.NoError(t, err)
	ids := lo.Map(todos, func(item *go_appv1.Todo, index int) string {
		return lo.FromPtr(item.Id)
	})
	require.NoError(t, DeleteTodos(ids))
}

func loadTodos(t *testing.T) {
	todos := CreateRandomNumTodos(t)
	LoadedTodoTestData.Todos = todos
	LoadedTodoTestData.TodoIds = hashset.New()
	for _, todo := range todos {
		LoadedTodoTestData.TodoIds.Add(*todo.Id)
	}
}

func CreateRandomNumTodos(t *testing.T) []*go_appv1.Todo {
	return CreateTodos(t, gofakeit.Number(5, 10))
}

func CreateTodos(t *testing.T, num int) []*go_appv1.Todo {
	var err error
	todos := []*go_appv1.Todo{}
	for i := 0; i < num; i++ {
		todo := createRandomTodoProto(t)
		todos = append(todos, todo)
	}
	todos, err = UpsertTodos(todos)
	require.NoError(t, err)
	return todos
}

func createRandomTodoProto(t *testing.T) *go_appv1.Todo {
	todo := &go_appv1.Todo{}
	err := gofakeit.Struct(todo)
	require.NoError(t, err)
	return todo
}

func randomizeTodos(t *testing.T, todos []*go_appv1.Todo) []*go_appv1.Todo {
	randomized := []*go_appv1.Todo{}
	for _, todo := range todos {
		random := createRandomTodoProto(t)
		random.Id = todo.Id
		randomized = append(randomized, random)
	}
	return randomized
}
