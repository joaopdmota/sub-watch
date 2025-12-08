package usecases

import (
	"context"
	"sub-watch-microservice/infra/repositories"
)


type ListUsersUseCase struct {
	repo repositories.UserRepository
}

func NewListUsersUseCase(repo repositories.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{repo: repo}
}

func (u *ListUsersUseCase) Execute(ctx context.Context) ([]UserOutput, error) {
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