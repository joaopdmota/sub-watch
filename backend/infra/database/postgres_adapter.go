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

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(30 * time.Minute)

	p.db = db
	return nil
}

func (p *PostgresAdapter) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

func (p *PostgresAdapter) Ping(ctx context.Context) error {
	if p.db == nil {
		return fmt.Errorf("database not connected")
	}
	return p.db.PingContext(ctx)
}

func (p *PostgresAdapter) GetClient() any {
	return p.db
}

func (p *PostgresAdapter) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database not connected")
	}
	return p.db.QueryContext(ctx, query, args...)
}

func (p *PostgresAdapter) FindAll(ctx context.Context, collection string) (Rows, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database not connected")
	}

	query := fmt.Sprintf("SELECT * FROM %s", collection)
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	return &PostgresRows{rows: rows}, nil
}

type PostgresRows struct {
	rows *sql.Rows
}

func (r *PostgresRows) Next() bool {
	return r.rows.Next()
}

func (r *PostgresRows) Scan(dest ...any) error {
	return r.rows.Scan(dest...)
}

func (r *PostgresRows) Close() error {
	return r.rows.Close()
}

func (r *PostgresRows) Err() error {
	return r.rows.Err()
}

func (p *PostgresAdapter) FindByID(ctx context.Context, collection, id string) (Rows, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database not connected")
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", collection)
	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	return &PostgresRows{rows: rows}, nil
}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}
