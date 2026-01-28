package domain

import "time"

type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "ACTIVE"
	SubscriptionStatusPaused   SubscriptionStatus = "PAUSED"
	SubscriptionStatusCanceled SubscriptionStatus = "CANCELED"
)

type Subscription struct {
	ID              string
	UserID          string
	CategoryID      string
	PaymentMethodID string
	ServiceName     string
	Price           float64
	Currency        string
	Cycle           string // "MONTHLY", "YEARLY"
	NextBillingDate time.Time
	Status          SubscriptionStatus
	Notes           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
