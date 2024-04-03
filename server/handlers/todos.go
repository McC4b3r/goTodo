package handlers

import (
	"context"
	"fmt"

	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
	"github.com/mcc4b3r/go_app/server/storage"
)

func (a ApiServer) UpsertTodos(ctx context.Context, request *go_appv1.UpsertTodosRequest) (*go_appv1.Todos, error) {
	fmt.Println("hits api handler")
	upsertedTodos, err := storage.Storage.UpsertTodos(ctx, request)
	return &go_appv1.Todos{Data: upsertedTodos}, err
}
func (a ApiServer) DeleteTodos(ctx context.Context, request *go_appv1.DeleteTodosRequest) (*go_appv1.DeleteTodosResponse, error) {
	return &go_appv1.DeleteTodosResponse{}, storage.Storage.DeleteTodos(ctx, request.Ids)
}
func (a ApiServer) ListTodos(ctx context.Context, request *go_appv1.ListTodosRequest) (*go_appv1.Todos, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	todos, err := storage.Storage.ListTodos(ctx, request)
	return &go_appv1.Todos{Data: todos}, err
}
func (a ApiServer) GetTodos(ctx context.Context, request *go_appv1.GetTodosRequest) (*go_appv1.Todos, error) {
	todos, err := storage.Storage.GetTodos(ctx, request)
	return &go_appv1.Todos{Data: todos}, err
}
