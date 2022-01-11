package services

import (
	"context"
	"fmt"

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

func (i *itemService) GetItem(ctx context.Context, id *grpc_item.ItemId) (*grpc_item.Item, error) {
	item, err := i.queries.GetItemById(ctx, id.Id)
	if err != nil {
		return &grpc_item.Item{}, err
	}
	return &grpc_item.Item{Id: item.ID, Name: item.Name, Value: item.Value}, nil
}

func (i *itemService) DeleteItem(ctx context.Context, id *grpc_item.ItemId) (*grpc_item.ItemId, error) {
	err := i.queries.DeleteItem(ctx, id.Id)
	if err != nil {
		return &grpc_item.ItemId{}, err
	}
	return id, nil
}

func (i *itemService) CreateItem(ctx context.Context, item *grpc_item.Item) (*grpc_item.ItemId, error) {
	id, err := i.queries.CreateItem(ctx, dbs.CreateItemParams{Name: item.Name, Value: item.Value})
	if err != nil {
		return &grpc_item.ItemId{}, err
	}
	return &grpc_item.ItemId{Id: id}, nil
}

func (i *itemService) UpdateItem(ctx context.Context, item *grpc_item.Item) (*grpc_item.ItemId, error) {
	err := i.queries.UpdateItem(ctx, dbs.UpdateItemParams{ID: item.Id, Name: item.Name, Value: item.Value})
	if err != nil {
		return &grpc_item.ItemId{}, err
	}
	return &grpc_item.ItemId{Id: item.Id}, nil
}

func (i *itemService) GetItems(ctx context.Context, v *grpc_item.Void) (*grpc_item.Items, error) {
	items, err := i.queries.GetItems(ctx)
	if err != nil {
		return &grpc_item.Items{}, err
	}

	itemsResponse := make([]*grpc_item.Item, len(items))
	for i, item := range items {
		itemsResponse[i] = &grpc_item.Item{Id: item.ID, Name: item.Name, Value: item.Value}
	}
	fmt.Println(itemsResponse)
	return &grpc_item.Items{Aliases: itemsResponse}, nil
}
