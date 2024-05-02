package main

import (
	"crud-connection/config"
	"crud-connection/model"
	"crud-connection/postgres"
	"database/sql"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	postgresDB, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Panic("Failed to connect to postgres", err)
	}
	defer postgresDB.Close()

	newID, err := uuid.NewV4()
	if err != nil {
		log.Fatal("Failed to create new uuid", err)
	}

	user := model.Users{
		Id:         newID.String(),
		FirstName:  "Do' 9876",
		TelegramId: 12321,
		Status:     "blocked",
		// StartTime:  time.Now(),
	}
	// createTable(postgresDB)
	userId := createUser(postgresDB, user)

	userData := selectUser(postgresDB, userId)
	fmt.Println("selected data -> ", userData)

	// err = updateUser(postgresDB, userData)
	// if err != nil {
	// 	log.Fatal("updateUser failed: ", err)
	// }

	// deleteUser(postgresDB, userId)

}

// func createTable(db *sql.DB) {
// 	query := `CREATE TABLE IF NOT EXISTS users(id uuid,first_name varchar,telegram_id int,status varchar,starttime timestamp)`
// 	db.QueryRow(query)
// 	fmt.Println("table created")
// }

func createUser(db *sql.DB, user model.Users) string {
	query := `
	INSERT INTO users(id,first_name,telegram_id,status,starttime)
	VALUES ($1,$2,$3,$4, CURRENT_TIMESTAMP) RETURNING id
	`
	var id string
	db.QueryRow(query, user.Id, user.FirstName, user.TelegramId, user.Status).Scan(&id)
	return id

}

func selectUser(db *sql.DB, id string) model.Users {
	var user model.Users
	query := `SELECT * FROM users WHERE id = $1`

	row := db.QueryRow(query, id)
	row.Scan(&user.Id, &user.FirstName, &user.TelegramId, &user.Status, &user.StartTime)
	return user

}

// func updateUser(db *sql.DB, user model.Users) error {
// 	query := `
// 		UPDATE users
// 		SET first_name = $1,
// 			telegram_id = $2,
// 			status = $3
// 		WHERE id = $4
// 	`

// 	res, err := db.Exec(query,
// 		user.FirstName,
// 		user.TelegramId,
// 		user.Status,
// 		user.Id,
// 	)
// 	if err != nil {
// 		log.Fatal("Error updating users ", err)
// 		return err
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		log.Fatal("Error  finding rows affected", err)
// 		return err
// 	}

// 	fmt.Println("count >>>>", count)

// 	return nil
// }

// func deleteUser(db *sql.DB, id string) {
// 	query := `
// 		DELETE FROM users WHERE id = $1
// 	`

// 	db.QueryRow(query, id)

// }

// func deleteUsers(db *sql.DB) {
// 	query := `
// 		DELETE FROM users WHERE id is not null
// 	`

// 	db.QueryRow(query)

// }
