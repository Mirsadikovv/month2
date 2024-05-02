package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"review/model"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
		}

		var user model.User

		err = json.Unmarshal(resBody, &user)
		if err != nil {
			fmt.Println("error while unmarshaling data", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
	}
}

func Login(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		var user model.User
		err = json.Unmarshal(resBody, &user)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		if len(user.Phone) != 13 && user.Phone[:3] != "998" {
			http.Error(w, "invalid phone number", http.StatusBadRequest)
			return
		}

		query := `SELECT "phone" FROM "user" WHERE "phone" = $1`

		res := db.QueryRow(ctx, query, user.Phone)

		err = res.Scan(&user.Phone)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

		fmt.Println("result: ", res)
		w.WriteHeader(http.StatusOK)

	}
}
