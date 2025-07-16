package driver

import (
	"context"
	"fmt"
	"hospital/internal/configs"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(dsn configs.Database) (*pgxpool.Pool, func(), error) {
	f := func() {}

	if !dsn.IsValid() {
		return nil, f, fmt.Errorf("invalid database config")
	}

	conf, err := pgxpool.ParseConfig(dsn.DSN())
	if err != nil {
		return nil, f, err
	}

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, f, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, f, err
	}

	return pool, func() { pool.Close() }, nil
}
