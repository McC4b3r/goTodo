package handlers

import (
	"context"
	"errors"
	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
	"github.com/mcc4b3r/go_app/server/storage"
)

type ApiServer struct{}

func (a ApiServer) Healthy(ctx context.Context, empty *go_appv1.Empty) (*go_appv1.Empty, error) {
	return &go_appv1.Empty{}, nil
}

func (a ApiServer) Ready(ctx context.Context, empty *go_appv1.Empty) (*go_appv1.Empty, error) {
	if storage.Storage.Ready() {
		return &go_appv1.Empty{}, nil
	}
	return nil, errors.New("")
}
