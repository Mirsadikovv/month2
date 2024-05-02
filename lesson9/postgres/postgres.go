package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"psql-connection/config"
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

	err = db.Ping()
	if err != nil {
		log.Fatal("error while pinging database", err)
	}
	return db, nil

}
func (p *Postgres) CloseDB() {
	p.db.Close()
}
