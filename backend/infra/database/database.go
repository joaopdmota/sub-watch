package database

import (
	"context"
	"database/sql"
)

type Database interface {
	Connect() error
	Close() error
	Ping(ctx context.Context) error
	GetClient() any
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	FindAll(ctx context.Context, collection string) (Rows, error)
	FindByID(ctx context.Context, collection, id string) (Rows, error)
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close() error
	Err() error
}
