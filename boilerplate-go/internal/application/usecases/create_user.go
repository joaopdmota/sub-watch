package usecases

import (
	"boilerplate-go/internal/application/domain"
	"context"
)

type CreateUserRequest struct {
	Name  string
	Email string
}

type CreateUserResponse struct {
	ID string
}

type CreateUserUseCase struct {
	// Add repository interface here later
}

func NewCreateUserUseCase() *CreateUserUseCase {
	return &CreateUserUseCase{}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, req CreateUserRequest) (CreateUserResponse, error) {
	user, err := domain.NewUser(req.Name, req.Email)
	if err != nil {
		return CreateUserResponse{}, err
	}

	// Persist user using repository...

	return CreateUserResponse{
		ID: user.ID.String(),
	}, nil
}
