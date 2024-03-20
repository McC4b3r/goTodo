package handlers

import (
	"context"
	go_appv1 "github.com/mcc4b3r/go_app/protos/gen/go/go_app/v1"
	"github.com/mcc4b3r/go_app/server/storage"
)

func (a ApiServer) UpsertHellos(ctx context.Context, request *go_appv1.UpsertHellosRequest) (*go_appv1.Hellos, error) {
	upsertedHellos, err := storage.Storage.UpsertHellos(ctx, request)
	return &go_appv1.Hellos{Hellos: upsertedHellos}, err
}

func (a ApiServer) DeleteHellos(ctx context.Context, request *go_appv1.DeleteRequest) (*go_appv1.DeleteResponse, error) {
	return &go_appv1.DeleteResponse{}, storage.Storage.DeleteHellos(ctx, request.Ids)
}

func (a ApiServer) ListHellos(ctx context.Context, request *go_appv1.ListRequest) (*go_appv1.Hellos, error) {
	if request.Limit == 0 {
		request.Limit = 100
	}
	hellos, err := storage.Storage.ListHellos(ctx, request)
	return &go_appv1.Hellos{Hellos: hellos}, err
}

func (a ApiServer) GetHellos(ctx context.Context, request *go_appv1.GetRequest) (*go_appv1.Hellos, error) {
	hellos, err := storage.Storage.GetHellos(ctx, request)
	return &go_appv1.Hellos{Hellos: hellos}, err
}
