package services

import (
	"context"

	"github.com/RomaBiliak/generate-grpc/database/dbs"
	grpc_item "github.com/RomaBiliak/generate-grpc/proto"
)

type itemRepository interface {
	CreateItem(ctx context.Context, arg dbs.CreateItemParams) (int32, error)
	DeleteItem(ctx context.Context, id int32)
	GetItemById(ctx context.Context, id int32) (dbs.Item, error)
	GetItems(ctx context.Context) ([]dbs.Item, error)
	UpdateItem(ctx context.Context, arg dbs.UpdateItemParams)
}
/*
type itemService struct {
	queries *itemRepository
}

func NewItemApiService(queries *itemRepository) *itemService {
	return &itemService{queries: queries}
}*/

type itemService struct {
	queries *dbs.Queries
	grpc_item.UnsafeCRUDServer
}

func NewItemService(queries *dbs.Queries) *itemService {
	return &itemService{queries: queries}
}


func (i *itemService) GetItem(context.Context, *grpc_item.ItemId) (*grpc_item.Item, error){
	return &grpc_item.Item{}, nil
}

func (i *itemService) DelteItem(context.Context, *grpc_item.ItemId) (*grpc_item.ItemId, error){
	return &grpc_item.ItemId{}, nil
}

func (i *itemService) CreateItem(context.Context, *grpc_item.Item) (*grpc_item.Item, error){
	return &grpc_item.Item{}, nil
}

func (i *itemService) UpdateItem(context.Context, *grpc_item.Item) (*grpc_item.Item, error){
	return &grpc_item.Item{}, nil
}

func (i *itemService) GetItems(context.Context, *grpc_item.Void) (*grpc_item.Items, error){
	return &grpc_item.Items{}, nil
}