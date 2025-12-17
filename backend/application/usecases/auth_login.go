package usecases

import (
	"context"
	app_errors "sub-watch-backend/application/errors"
	"sub-watch-backend/infra/repositories"
	"sub-watch-backend/pkg/hash"
)

type AuthLoginUseCase struct {
	repo repositories.UserRepository
	hasher hash.PasswordHasher
}

func NewAuthLoginUseCase(userRepo repositories.UserRepository, hasher hash.PasswordHasher) *AuthLoginUseCase {
	return &AuthLoginUseCase{repo: userRepo, hasher: hasher}
}

type AuthLoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

func (u *AuthLoginUseCase) Execute(ctx context.Context, input AuthLoginInput) (string, *app_errors.Error) {
	user, appErr := u.repo.FindByEmail(ctx, input.Email)
	if appErr != nil {
		if appErr.Code == 404 {
			return "", &app_errors.Error{
				Code:    400,
				Message: "Invalid email or password",
			}
		}
		return "", appErr
	}

	if err := u.hasher.Compare(input.Password, user.PasswordHash); err != nil {
		return "", &app_errors.Error{
			Code:    400,
			Message: "Invalid email or password",
		}
	}

	return user.ID, nil
}
