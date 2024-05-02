package main

import (
	"context"
	"exam-project/config"
	"exam-project/postgres"
	"exam-project/storage"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	postgresDB, err := postgres.NewPostgres(cfg, ctx)
	if err != nil {
		log.Panic("failed to connect database", err)
	}
	defer postgresDB.Close()

	http.HandleFunc("/login", storage.Login(ctx, postgresDB))
	http.HandleFunc("/register", storage.Register(ctx, postgresDB))
	http.HandleFunc("/create_task", storage.CreateTask(ctx, postgresDB))
	http.HandleFunc("/update_task", storage.UpdateTask(ctx, postgresDB))
	http.HandleFunc("/select_tasks", storage.SelectTask(ctx, postgresDB))
	http.HandleFunc("/select_priority", storage.SelectPriority(ctx, postgresDB))
	http.HandleFunc("/delete_task", storage.DeleteTask(ctx, postgresDB))

	fmt.Println("app is running...")

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("error while listening", err)
	}

}
