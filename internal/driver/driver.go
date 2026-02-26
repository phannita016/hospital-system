package driver

import (
	"fmt"
	"hospital/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn config.Database) (*gorm.DB, func(), error) {
	f := func() {}

	if !dsn.IsValid() {
		return nil, f, fmt.Errorf("invalid database config")
	}

	var db *gorm.DB
	var err error
	maxAttempts := 3
	for i := 1; i <= maxAttempts; i++ {
		db, err = gorm.Open(postgres.Open(dsn.DSN()), &gorm.Config{})
		if err == nil && db != nil {
			sqlDB, err := db.DB()
			if err == nil {
				err = sqlDB.Ping()
				if err == nil {
					return db, func() { sqlDB.Close() }, nil
				}
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
