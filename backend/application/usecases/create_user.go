package usecases

import (
	"context"
	"fmt"
	"sub-watch-backend/application/domain"
	app_errors "sub-watch-backend/application/errors"
	"sub-watch-backend/infra/repositories"
	"sub-watch-backend/pkg/date"
	"sub-watch-backend/pkg/hash"
	id "sub-watch-backend/pkg/uuid"
)

type CreateUserUseCase struct {
	repo repositories.UserRepository
	uuid id.UuidProvider
	hasher hash.PasswordHasher
	date date.DateProvider
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewCreateUserUseCase(repo repositories.UserRepository, uuid id.UuidProvider, hasher hash.PasswordHasher, date date.DateProvider) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo, uuid: uuid, hasher: hasher, date: date}
}

func (u *UserInput) Validate() *app_errors.Error {
	if u.Name == "" {
		return &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "Name is required",
		}
	}
	if u.Email == "" {
		return &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "Email is required",
		}
	}
	if u.Password == "" {
		return &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "Password is required",
		}
	}

	return nil
}

func (u *CreateUserUseCase) Execute(ctx context.Context, input UserInput) *app_errors.Error {
	userInput := &UserInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := userInput.Validate(); err != nil {
		return err
	}

	hashedPassword, err := u.hasher.Hash(userInput.Password)
	if err != nil {
		return &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: "Failed to hash password",
		}
	}

	user, err := u.repo.FindByEmail(ctx, userInput.Email)
	if err != nil {
		return &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: "Failed to find user by email",
		}
	}

	if user != nil {
		return &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "User already exists",
		}
	}

	userEntity := domain.User{
		UpdatedAt: u.date.Now(),
		CreatedAt: u.date.Now(),
		ID:       u.uuid.NewID(),
		Name:     userInput.Name,
		Email:    userInput.Email,
		PasswordHash: hashedPassword,
	}

	if err := u.repo.Insert(ctx, &userEntity); err != nil {
		fmt.Println(err)
		return &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: "Failed to insert user",
		}
	}

	return nil
}