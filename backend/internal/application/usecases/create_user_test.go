package usecases

import (
	"context"
	"sub-watch-backend/internal/application/domain"
	"testing"
	"time"
)

func TestCreateUserUseCase_Execute(t *testing.T) {
	now := time.Now()
	
	tests := []struct {
		name        string
		input       UserInput
		existingUsers map[string]*domain.User
		wantErr     bool
		errType     string
	}{
		{
			name: "Success - create new user",
			input: UserInput{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			existingUsers: make(map[string]*domain.User),
			wantErr:       false,
		},
		{
			name: "Failure - user already exists",
			input: UserInput{
				Name:     "John Doe",
				Email:    "john@example.com",
				Password: "password123",
			},
			existingUsers: map[string]*domain.User{
				"john@example.com": {Email: "john@example.com"},
			},
			wantErr: true,
			errType: "BAD_REQUEST",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &FakeUserRepository{Users: tt.existingUsers}
			uuid := &FakeUUIDProvider{ID: "123"}
			hasher := &FakeHasher{}
			date := &FakeDateProvider{Time: now}
			validator := &FakeValidator{}
			log := &FakeLogger{}

			usecase := NewCreateUserUseCase(repo, uuid, hasher, date, validator, log)

			err := usecase.Execute(context.Background(), tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil && err.Type != tt.errType {
				t.Errorf("Execute() error type = %v, want %v", err.Type, tt.errType)
			}

			if !tt.wantErr {
				user, _ := repo.FindByEmail(context.Background(), tt.input.Email)
				if user == nil {
					t.Errorf("Expected user to be created")
				}
				if user.Name != tt.input.Name {
					t.Errorf("Expected name %s, got %s", tt.input.Name, user.Name)
				}
			}
		})
	}
}
