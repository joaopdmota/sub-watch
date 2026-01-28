package domain

import "time"

type PriceHistory struct {
	ID             string
	SubscriptionID string
	OldPrice       float64
	NewPrice       float64
	ChangedAt      time.Time
	Reason         string
}
