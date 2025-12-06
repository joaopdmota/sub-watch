package database

import (
	"context"
)

type Database interface {
	Connect() error
	Close() error
	Ping(ctx context.Context) error
	GetClient() any
}
