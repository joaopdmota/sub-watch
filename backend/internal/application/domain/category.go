package domain

import "time"

type Category struct {
	ID      string
	Name    string
	Icon    string
	Color   string
	CreatedAt time.Time
}
