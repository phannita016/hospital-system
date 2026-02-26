package driver

import (
	"context"
	"fmt"
	"hospital/config"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB(dsn config.Database) (*pgxpool.Pool, func(), error) {
	f := func() {}

	if !dsn.IsValid() {
		return nil, f, fmt.Errorf("invalid database config")
	}

	conf, err := pgxpool.ParseConfig(dsn.DSN())
	if err != nil {
		return nil, f, err
	}

	var pool *pgxpool.Pool
	maxAttempts := 3
	for i := 1; i <= maxAttempts; i++ {
		pool, err = pgxpool.NewWithConfig(context.Background(), conf)
		if err == nil && pool != nil {
			err = pool.Ping(context.Background())
			if err == nil {
				return pool, func() { pool.Close() }, nil
			}
		}

		log.Printf("Waiting for DB to be ready (%d/%d)...", i, maxAttempts)
		time.Sleep(3 * time.Second)
	}

	return nil, f, fmt.Errorf("database not ready after %d attempts: %w", maxAttempts, err)
}

// func NewDB(dsn config.Database) (*pgxpool.Pool, func(), error) {
// 	f := func() {}

// 	if !dsn.IsValid() {
// 		return nil, f, fmt.Errorf("invalid database config")
// 	}

// 	conf, err := pgxpool.ParseConfig(dsn.DSN())
// 	if err != nil {
// 		return nil, f, err
// 	}

// 	ctx := context.Background()
// 	pool, err := pgxpool.NewWithConfig(ctx, conf)
// 	if err != nil {
// 		return nil, f, err
// 	}

// 	if err := pool.Ping(context.Background()); err != nil {
// 		return nil, f, err
// 	}

// 	return pool, func() { pool.Close() }, nil
// }
