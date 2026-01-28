package usecases

import (
	"context"
"sub-watch-backend/internal/application"
	app_errors "sub-watch-backend/internal/application/errors"
	"time"
)

type GetUserUseCase struct {
	repo  application.UserRepository
}

type UserOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewGetUserUseCase(repo application.UserRepository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

func (u *GetUserUseCase) Execute(ctx context.Context, id string) (*UserOutput, *app_errors.Error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return &UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}