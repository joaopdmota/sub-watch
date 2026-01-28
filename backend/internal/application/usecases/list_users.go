package usecases

import (
	"context"
"sub-watch-backend/internal/application"
	app_errors "sub-watch-backend/internal/application/errors"
)

type ListUsersUseCase struct {
	repo application.UserRepository
}

func NewListUsersUseCase(repo application.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

func (u *ListUsersUseCase) Execute(ctx context.Context) ([]UserOutput, *app_errors.Error) {
	users, err := u.repo.FindAll(ctx)
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