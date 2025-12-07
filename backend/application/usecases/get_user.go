package usecases

import (
	"context"
	"sub-watch/application/services"
)

type GetUserUseCase struct {
	service services.UserService
}

func NewGetUserUseCase(service services.UserService) *GetUserUseCase {
	return &GetUserUseCase{service: service}
}

func (u *GetUserUseCase) Execute(ctx context.Context, id string) (*UserOutput, error) {
	user, err := u.service.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil // User not found
	}

	return &UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
