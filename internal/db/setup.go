package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Config holds database configuration
type Config struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    PoolMax  int32
}

// NewDBPool creates a new Database instance with connection pool
func NewDBPool(dbURL string) (*pgxpool.Pool, error) {
    // Configure the connection pool
    poolConfig, err := pgxpool.ParseConfig(dbURL)
    if err != nil {
        return nil, fmt.Errorf("error parsing pool config: %v", err)
    }

    // Set pool configuration
    poolConfig.MinConns = 2
    poolConfig.MaxConnLifetime = time.Hour
    poolConfig.MaxConnIdleTime = 30 * time.Minute
    poolConfig.HealthCheckPeriod = 1 * time.Minute

    // Create the connection pool
    pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
    if err != nil {
        return nil, fmt.Errorf("error creating connection pool: %v", err)
    }

	// pass the `queries` object to functions that require it
	return pool, nil
}