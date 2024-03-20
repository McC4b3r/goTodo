package test

import (
	"fmt"

	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/testing/protocmp"
)

func (s *Go_appSuite) TestUpdateTodos() {
	todos := CreateRandomNumTodos(s.T())
	fmt.Println("Todos:")
	for i, todo := range todos {
		fmt.Printf("Todo %d:\n", i+1)
		// Safe dereferencing for *string fields
		id := ""
		if todo.Id != nil {
			id = *todo.Id
		}
		todoName := ""
		if todo.TodoName != nil {
			todoName = *todo.TodoName
		}
		// Adjusted print statements
		fmt.Printf("  ID:        %s\n", id)
		fmt.Printf("  Type:      %s\n", todo.TodoType.String())
		fmt.Printf("  Name:      %s\n", todoName)
		fmt.Printf("  Priority:  %d\n", todo.Priority)
		fmt.Printf("  Completed: %t\n", todo.Completed)
		fmt.Println()
	}
	randomizedTodos := randomizeTodos(s.T(), todos)
	updatedTodos, err := UpsertTodos(randomizedTodos)
	require.NoError(s.T(), err)

	assertProtoEqualitySortById(s.T(), randomizedTodos, updatedTodos,
		protocmp.IgnoreFields(&go_appv1.Todo{}, "id", "todo_type", "todo_name"),
	)
}

func (s *Go_appSuite) TestListTodos() {
	Todos, err := ListTodos(1000, 0, "")
	require.NoError(s.T(), err)
	Todos = lo.Filter(Todos, func(item *go_appv1.Todo, index int) bool {
		return LoadedTodoTestData.TodoIds.Contains(lo.FromPtr(item.Id))
	})
	assertProtoEqualitySortById(s.T(), LoadedTodoTestData.Todos, Todos)
}

func (s *Go_appSuite) TestGetTodosById() {
	ids := []string{}
	for _, id := range LoadedTodoTestData.TodoIds.Values() {
		ids = append(ids, id.(string))
	}
	Todos, err := GetTodosById(ids)
	require.NoError(s.T(), err)
	assertProtoEqualitySortById(s.T(), LoadedTodoTestData.Todos, Todos)
}

func (s *Go_appSuite) TestDeleteTodos() {
	Todos := CreateRandomNumTodos(s.T())
	ids := lo.Map(Todos, func(item *go_appv1.Todo, index int) string {
		return lo.FromPtr(item.Id)
	})
	err := DeleteTodos(ids)
	require.NoError(s.T(), err)
	Todos, err = GetTodosById(ids)
	require.NoError(s.T(), err)
	require.Len(s.T(), Todos, 0)
}
