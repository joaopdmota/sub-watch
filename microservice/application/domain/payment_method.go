package domain

import "time"

type PaymentMethod struct {
	ID        string
	UserID    string
	Name      string // e.g., "Nubank Black"
	Type      string // "CREDIT_CARD", "DEBIT", "PIX", etc.
	CreatedAt time.Time
}
