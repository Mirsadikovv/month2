package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"review/config"
	"review/postgres"
	"review/storage"
)

func main() {
	cfg := config.Config{}
	ctx := context.Background()

	postgresDB, err := postgres.NewPostgres(cfg, ctx)

	if err != nil {
		log.Panic("failed to connect database", err)
	}
	defer postgresDB.Close()

	http.HandleFunc("/user", storage.CreateUser(ctx, postgresDB))

	fmt.Println("app is running...")

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("error while listening", err)
	}

}
