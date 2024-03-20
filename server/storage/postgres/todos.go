package postgres

import (
	"context"
	"fmt"

	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
)

func (p PostgresStorage) UpsertTodos(ctx context.Context, request *go_appv1.UpsertTodosRequest) ([]*go_appv1.Todo, error) {
	todoProtos := go_appv1.TodoProtos(request.Todos)
	_, err := todoProtos.Upsert(ctx, db)
	return todoProtos, err
}

func (p PostgresStorage) DeleteTodos(ctx context.Context, ids []string) error {
	return go_appv1.DeleteTodoGormModels(ctx, db, ids)
}

func (p PostgresStorage) ListTodos(ctx context.Context, request *go_appv1.ListTodosRequest) ([]*go_appv1.Todo, error) {
	protos := go_appv1.TodoProtos{}
	var orderBy string
	err := protos.List(ctx, db, int(request.Limit), int(request.Offset), orderBy)
	if err != nil {
		return nil, fmt.Errorf("error listing Hellos: %w", err)
	}
	return protos, nil
}
func (p PostgresStorage) GetTodos(ctx context.Context, request *go_appv1.GetTodosRequest) ([]*go_appv1.Todo, error) {
	protos := go_appv1.TodoProtos{}
	err := protos.GetByIds(ctx, db, request.Ids)
	return protos, err
}
