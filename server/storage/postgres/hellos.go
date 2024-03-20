package postgres

import (
	"context"
	"fmt"

	"github.com/joomcode/errorx"
	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
)

func (p PostgresStorage) UpsertHellos(ctx context.Context, request *go_appv1.UpsertHellosRequest) ([]*go_appv1.Hello, error) {
	helloProtos := go_appv1.HelloProtos(request.Hellos)
	_, err := helloProtos.Upsert(ctx, db)
	return helloProtos, err
}

func (p PostgresStorage) DeleteHellos(ctx context.Context, ids []string) error {
	return go_appv1.DeleteHelloGormModels(ctx, db, ids)
}

func (p PostgresStorage) ListHellos(ctx context.Context, request *go_appv1.ListRequest) ([]*go_appv1.Hello, error) {
	protos := go_appv1.HelloProtos{}
	var orderBy string
	if orderBy = request.OrderBy; orderBy == "" {
		orderBy = "created_at"
	}
	err := protos.List(ctx, db, int(request.Limit), int(request.Offset), orderBy)
	if err != nil {
		return nil, fmt.Errorf("error listing Hellos: %w", err)
	}
	return protos, nil
}

func (p PostgresStorage) GetHellos(ctx context.Context, request *go_appv1.GetRequest) ([]*go_appv1.Hello, error) {
	protos := go_appv1.HelloProtos{}
	err := protos.GetByIds(ctx, db, request.Ids)
	return protos, err
}

func (p PostgresStorage) GetHello(ctx context.Context, id string) (*go_appv1.Hello, error) {
	var err error
	protos := go_appv1.HelloProtos{}
	if err = protos.GetByIds(ctx, db, []string{id}); err != nil {
		return nil, err
	}
	if len(protos) == 0 {
		return nil, errorx.IllegalArgument.New("hello with id %s not found", id)
	}
	return protos[0], nil
}
