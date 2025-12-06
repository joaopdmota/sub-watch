package domain

type Category struct {
	ID      string
	UserID  *string // Nullable for system defaults
	Name    string
	Icon    string
	Color   string
}
