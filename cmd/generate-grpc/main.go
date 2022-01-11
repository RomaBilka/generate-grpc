//Home task (retraining program)
package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/RomaBiliak/generate-grpc/database/dbs"
	services "github.com/RomaBiliak/generate-grpc/internal/service"
	"github.com/RomaBiliak/generate-grpc/pkg/database/postgres"
	grpc_item "github.com/RomaBiliak/generate-grpc/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	pgUser, ok := os.LookupEnv("PG_USER")
	if !ok {
		panic(errors.New("PG_USER is empty"))
	}
	pgPassword, ok := os.LookupEnv("PG_PASSWORD")
	if !ok {
		panic(errors.New("PG_PASSWORD is empty"))
	}
	pgHost, ok := os.LookupEnv("PG_HOST")
	if !ok {
		panic(errors.New("Db Host is empty"))
	}
	pgDatabase, ok := os.LookupEnv("PG_DATABASE")
	if !ok {
		panic(errors.New("PG_DATABASE is empty"))
	}

	dbConfig := postgres.Config{
		pgUser,
		pgPassword,
		pgDatabase,
		pgHost,
	}
	db := postgres.Run(dbConfig)
	defer db.Close()

	queries := dbs.New(db)

	server := grpc.NewServer()

	serviceItem := services.NewItemService(queries)
	grpc_item.RegisterCRUDServer(server, serviceItem)

	l, err:= net.Listen("tcp", ":8080")

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Listen 8080")
	if err:= server.Serve(l); err != nil{
		log.Fatal(err)
	}
}
