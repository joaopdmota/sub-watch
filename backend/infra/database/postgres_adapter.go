package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresAdapter struct {
	db *sql.DB
}

func (p *PostgresAdapter) Connect() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Set pool settings
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	p.db = db
	return nil
}

// Close closes the database connection
func (p *PostgresAdapter) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

// Ping checks the database connection
func (p *PostgresAdapter) Ping(ctx context.Context) error {
	if p.db == nil {
		return fmt.Errorf("database not connected")
	}
	return p.db.PingContext(ctx)
}

func (p *PostgresAdapter) GetClient() any {
	return p.db
}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}
