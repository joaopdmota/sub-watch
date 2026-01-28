package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"sub-watch-backend/internal/infra/database"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresAdapter struct {
	db *sql.DB
}

var ErrNoRows = sql.ErrNoRows

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{}
}

func (p *PostgresAdapter) Connect() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" {
		return fmt.Errorf("database environment variables are not fully set")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
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

type PostgresResult struct {
	res sql.Result
}

func (r *PostgresResult) RowsAffected() (int64, error) {
	return r.res.RowsAffected()
}

func (r *PostgresResult) LastInsertID() (int64, error) {
	return r.res.LastInsertId()
}

func (p *PostgresAdapter) QueryContext(
	ctx context.Context,
	query string,
	args ...any,
) (database.Rows, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database not connected")
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &PostgresRows{rows: rows}, nil
}

func (p *PostgresAdapter) ExecContext(
	ctx context.Context,
	query string,
	args ...any,
) (database.Result, error) {
	if p.db == nil {
		return nil, fmt.Errorf("database not connected")
	}

	res, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &PostgresResult{res: res}, nil
}

func (p *PostgresAdapter) FindAll(
	ctx context.Context,
	collection string,
) (database.Rows, error) {
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

func (p *PostgresAdapter) FindByID(
	ctx context.Context,
	collection, id string,
) (database.Rows, error) {
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

func (p *PostgresAdapter) Delete(
	ctx context.Context,
	collection, id string,
) error {
	if p.db == nil {
		return fmt.Errorf("database not connected")
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", collection)

	_, err := p.db.ExecContext(ctx, query, id)
	return err
}
