package main

import (
	"context"
	"fmt"
	"http-api/config"
	"http-api/model"
	"http-api/postgres"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()

	ctx := context.Background()

	postgresDB, err := postgres.NewPostgres(cfg, ctx)
	if err != nil {
		log.Panic("Failed to connect to postgres", err)
	}

	user := model.User{
		Name:  "new user",
		Email: "example@example.com",
		Phone: "+998000000000",
	}

	defer postgresDB.Close()

	// http.HandleFunc("/user", createUser(postgresDB, user, ctx))
	http.HandleFunc("/login", login(postgresDB, user, ctx))

	fmt.Println("application running...")

	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("error when listening", err)
	}
}

func createUser(db *pgxpool.Pool, user model.User, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		query := `
			INSERT INTO "user" ("id", "name",  "email", "phone")
			VALUES ($1, $2, $3, $4) RETURNING id`

		id, err := uuid.NewV4()
		if err != nil {
			fmt.Println("error when creating uuid", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		res, err := db.Exec(ctx, query, id.String(), user.Name, user.Email, user.Phone)
		if err != nil {
			fmt.Println("error when creating user", err)

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println("result :", res)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fmt.Sprintf("%s ->id of user created", id))
	}
}

func login(db *pgxpool.Pool, user model.User, ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		fmt.Println("phone", user.Phone)
		query := `SELECT "phone" FROM "user" WHERE "phone" = $1`

		res := db.QueryRow(ctx, query, user.Phone)

		var user model.User

		err := res.Scan(&user.Phone)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Error(w, "Not Found", http.StatusNotFound)
				return
			}

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		fmt.Println("result :", res)
		w.WriteHeader(http.StatusOK)
		// fmt.Fprintf(w, fmt.Sprintf("%s ->id of user created", id))
	}
}
