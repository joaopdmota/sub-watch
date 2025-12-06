package database

import (
	"context"
	"fmt"
	"time"
)

func NewConnection() (Database, error) {
	adapter := NewPostgresAdapter()

	if err := adapter.Connect(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := adapter.Ping(ctx); err != nil {
		adapter.Close()
		return nil, fmt.Errorf("failed to verify database connection: %w", err)
	}

	return adapter, nil
}
