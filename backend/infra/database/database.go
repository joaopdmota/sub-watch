package database

import "context"

type Database interface {
	Connect() error
	Close() error
	Ping(ctx context.Context) error

	GetClient() any

	QueryContext(ctx context.Context, query string, args ...any) (Rows, error)
	ExecContext(ctx context.Context, query string, args ...any) (Result, error)

	FindAll(ctx context.Context, collection string) (Rows, error)
	FindByID(ctx context.Context, collection, id string) (Rows, error)
	Delete(ctx context.Context, collection, id string) error
}

type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close() error
	Err() error
}

type Result interface {
	RowsAffected() (int64, error)
	LastInsertID() (int64, error)
}
