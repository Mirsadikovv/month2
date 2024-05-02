package postgres

import (
	"context"
	"fmt"
	"review/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(cfg config.Config, ctx context.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%d@%s:%s/%s?sslmode=disable", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser,
		cfg.PostgresPassword, cfg.PostgresDatabase),
	)

	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}
