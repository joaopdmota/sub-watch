package database

import (
	"context"
	"fmt"
	"time"
)

func NewConnection(adapter Database) (Database, error) {
	if err := adapter.Connect(); err != nil {
		return nil, err
	}
	fmt.Println("Database connection established")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()


	return adapter, nil
}