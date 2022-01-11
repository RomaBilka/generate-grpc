package main

import (
	"context"
	"fmt"
	"log"

	grpc_item "github.com/RomaBiliak/generate-grpc/proto"
	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := grpc_item.NewCRUDClient(conn)

	itemId, err := client.CreateItem(context.Background(), &grpc_item.Item{Name: "test", Value: "test_value"})
	fmt.Println(itemId)
	fmt.Println(err)

	item, err := client.GetItem(context.Background(), &grpc_item.ItemId{Id:1}, )
	fmt.Println(item)
	fmt.Println(err)

	items, err := client.GetItems(context.Background(),  &grpc_item.Void{})
	fmt.Println(items.GetAliases())
	fmt.Println(err)

	itemId, err = client.DeleteItem(context.Background(), &grpc_item.ItemId{Id:1}, )
	fmt.Println(itemId)
	fmt.Println(err)

	itemId, err = client.UpdateItem(context.Background(), &grpc_item.Item{Id: 2, Name: "new_test", Value: "test_value"})
	fmt.Println(itemId)
	fmt.Println(err)

	items, err = client.GetItems(context.Background(),  &grpc_item.Void{})
	fmt.Println(items.GetAliases())
	fmt.Println(err)
}
