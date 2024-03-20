package storage

import (
	"context"

	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
)

var Storage StorageInterface

type StorageInterface interface {
	Initialize() (shutdown func(), err error)
	Ready() bool
	UpsertHellos(ctx context.Context, request *go_appv1.UpsertHellosRequest) ([]*go_appv1.Hello, error)
	DeleteHellos(ctx context.Context, ids []string) error
	ListHellos(ctx context.Context, request *go_appv1.ListRequest) ([]*go_appv1.Hello, error)
	GetHellos(ctx context.Context, request *go_appv1.GetRequest) ([]*go_appv1.Hello, error)
	GetHello(ctx context.Context, id string) (*go_appv1.Hello, error)
	UpsertTodos(ctx context.Context, request *go_appv1.UpsertTodosRequest) ([]*go_appv1.Todo, error)
	DeleteTodos(ctx context.Context, ids []string) error
	ListTodos(ctx context.Context, request *go_appv1.ListTodosRequest) ([]*go_appv1.Todo, error)
	GetTodos(ctx context.Context, request *go_appv1.GetTodosRequest) ([]*go_appv1.Todo, error)
}
