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
func NewDBPool(cfg Config) (*pgxpool.Pool, error) {
    // Construct connection string
    connStr := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
    )

    // Configure the connection pool
    poolConfig, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        return nil, fmt.Errorf("error parsing pool config: %v", err)
    }

    // Set pool configuration
    poolConfig.MaxConns = cfg.PoolMax
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