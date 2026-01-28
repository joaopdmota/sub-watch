package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidUserName = errors.New("invalid user name")
)

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

func NewUser(name, email string) (*User, error) {
	if name == "" {
		return nil, ErrInvalidUserName
	}
	return &User{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
	}, nil
}
