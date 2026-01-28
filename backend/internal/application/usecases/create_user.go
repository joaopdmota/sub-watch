package usecases

import (
	"context"
	"sub-watch-backend/internal/application"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
	"sub-watch-backend/internal/pkg/date"
	"sub-watch-backend/internal/pkg/hash"
	id "sub-watch-backend/internal/pkg/uuid"
)

type CreateUserUseCase struct {
	repo      application.UserRepository
	uuid      id.UuidProvider
	hasher    hash.PasswordHasher
	date      date.DateProvider
	validator application.Validator
	log       application.Logger
}

type UserInput struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

func NewCreateUserUseCase(repo application.UserRepository, uuid id.UuidProvider, hasher hash.PasswordHasher, date date.DateProvider, validator application.Validator, log application.Logger) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo, uuid: uuid, hasher: hasher, date: date, validator: validator, log: log}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, input UserInput) *app_errors.Error {
	if err := u.validator.ValidateStruct(input); err != nil {
		formattedErrors := u.validator.FormatValidationErrors(err)

		return &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "Validation failed",
			Details: formattedErrors,
		}
	}

	hashedPassword, err := u.hasher.Hash(input.Password)
	if err != nil {
		return &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: "Failed to hash password",
		}
	}

	user, appErr := u.repo.FindByEmail(ctx, input.Email)
	if appErr != nil {
		if appErr.Code != 404 {
			return appErr
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
		Name:     input.Name,
		Email:    input.Email,
		PasswordHash: hashedPassword,
	}

	if err := u.repo.Insert(ctx, &userEntity); err != nil {
		u.log.Error("CreateUserUseCase.Execute: failed to insert user", "error", err)
		return &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: "Failed to insert user",
		}
	}

	return nil
}