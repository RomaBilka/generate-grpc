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

	response, err := client.CreateItem(context.Background(), &grpc_item.Item{Name: "test", Value: "test_value"})
	fmt.Println(response)
	fmt.Println(err)

	response, err = client.GetItem(context.Background(), &grpc_item.ItemId{Id:1})
	fmt.Println(response)
	fmt.Println(err)
}
