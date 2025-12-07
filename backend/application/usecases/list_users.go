package usecases

import (
	"context"
	"sub-watch/application/services"
	"time"
)

type UserOutput struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type ListUsersUseCase struct {
	service services.UserService
}

func NewListUsersUseCase(service services.UserService) *ListUsersUseCase {
	return &ListUsersUseCase{service: service}
}

func (u *ListUsersUseCase) Execute(ctx context.Context) ([]UserOutput, error) {
	users, err := u.service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var output []UserOutput
	for _, user := range users {
		output = append(output, UserOutput{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	return output, nil
}
