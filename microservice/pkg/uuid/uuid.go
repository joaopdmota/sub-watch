package id

import "github.com/google/uuid"

type UuidProvider interface {
	NewID() string
}

type UUIDProvider struct{}

func NewUUIDProvider() UuidProvider {
	return &UUIDProvider{}
}

func (u *UUIDProvider) NewID() string {
	return uuid.New().String()
}
