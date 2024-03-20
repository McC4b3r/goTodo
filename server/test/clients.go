package test

import (
	"context"
	"time"

	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
	"google.golang.org/grpc"
)

var ApiClient = initializeApiClient(5 * time.Second)

func initializeApiClient(timeout time.Duration) go_appv1.ApiClient {
	connectTo := "127.0.0.1:6000"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	conn, err := grpc.DialContext(ctx, connectTo, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return go_appv1.NewApiClient(conn)
}

// begin todo client funcs
func UpsertTodos(todos []*go_appv1.Todo) ([]*go_appv1.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.UpsertTodos(ctx, &go_appv1.UpsertTodosRequest{Todos: todos})
	if err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else {
		return response.Todos, err
	}
}

func DeleteTodos(ids []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ApiClient.DeleteTodos(ctx, &go_appv1.DeleteTodosRequest{Ids: ids})
	return err
}

func ListTodos(limit, offset int, orderBy string) ([]*go_appv1.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.ListTodos(ctx, &go_appv1.ListTodosRequest{Limit: int32(limit), Offset: int32(offset), OrderBy: orderBy})
	if err != nil {
		return nil, err
	}
	return response.Todos, err
}

func GetTodosById(ids []string) ([]*go_appv1.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response, err := ApiClient.GetTodos(ctx, &go_appv1.GetTodosRequest{Ids: ids})
	return response.Todos, err
}
