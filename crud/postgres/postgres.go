package postgres

import (
	"crud-connection/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("can not connect to database", err)
	}
	// defer db.Close()
	////////////////////////s

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatal("failed connection to database: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migration/",
		"postgres", driver)
	if err != nil {
		fmt.Println(err.Error(), &driver)
		log.Fatal("failed to create new migration instance:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to apply migrations:", err)
	}
	/////////////////////////////

	err = db.Ping()
	if err != nil {
		log.Fatal("error while pinging database", err)
	}
	return db, nil

}
func (p *Postgres) CloseDB() {
	p.db.Close()
}
