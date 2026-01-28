package usecases

import (
	"context"
	"fmt"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
	"time"
)

type FakeUserRepository struct {
	Users map[string]*domain.User
}

func (f *FakeUserRepository) Insert(ctx context.Context, user *domain.User) *app_errors.Error {
	f.Users[user.Email] = user
	return nil
}

func (f *FakeUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, *app_errors.Error) {
	user, ok := f.Users[email]
	if !ok {
		return nil, app_errors.NewNotFoundError("user not found", nil)
	}
	return user, nil
}

func (f *FakeUserRepository) FindByID(ctx context.Context, id string) (*domain.User, *app_errors.Error) {
	for _, u := range f.Users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, app_errors.NewNotFoundError("user not found", nil)
}

func (f *FakeUserRepository) FindAll(ctx context.Context) ([]domain.User, *app_errors.Error) {
	var users []domain.User
	for _, u := range f.Users {
		users = append(users, *u)
	}
	return users, nil
}

func (f *FakeUserRepository) Delete(ctx context.Context, id string) *app_errors.Error {
	for email, u := range f.Users {
		if u.ID == id {
			delete(f.Users, email)
			return nil
		}
	}
	return app_errors.NewNotFoundError("user not found", nil)
}

type FakeUUIDProvider struct {
	ID string
}

func (f *FakeUUIDProvider) NewID() string {
	return f.ID
}

type FakeHasher struct{}

func (f *FakeHasher) Hash(password string) (string, error) {
	return "hashed_" + password, nil
}

func (f *FakeHasher) Compare(hash, password string) error {
	if hash != "hashed_"+password {
		return fmt.Errorf("invalid password")
	}
	return nil
}

type FakeDateProvider struct {
	Time time.Time
}

func (f *FakeDateProvider) Now() time.Time {
	return f.Time
}

type FakeValidator struct{}

func (f *FakeValidator) ValidateStruct(s interface{}) error {
	return nil
}

func (f *FakeValidator) FormatValidationErrors(err error) map[string]interface{} {
	return nil
}

type FakeLogger struct{}

func (f *FakeLogger) Info(msg string, kv ...any)  {}
func (f *FakeLogger) Warn(msg string, kv ...any)  {}
func (f *FakeLogger) Error(msg string, kv ...any) {}
func (f *FakeLogger) Debug(msg string, kv ...any) {}
