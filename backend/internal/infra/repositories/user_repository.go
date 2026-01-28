package repositories

import (
	"context"
	"fmt"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
	"sub-watch-backend/internal/infra/database"
	"sub-watch-backend/internal/infra/database/adapters"
)


type UserRepository struct {
	db database.Database
}

func NewUserRepository(db database.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(ctx context.Context, user *domain.User) *app_errors.Error {
	const q = `
		INSERT INTO users (id, name, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, q,
		user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return app_errors.NewInternalError(
			"UserRepository.Insert: failed to execute insert",
			fmt.Errorf("db exec: %w", err),
		)
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, *app_errors.Error) {
	const q = "SELECT * FROM users WHERE email = $1"

	rows, err := r.db.QueryContext(ctx, q, email)
	if err != nil {
		return nil, app_errors.NewInternalError(
			"UserRepository.FindByEmail: failed to execute query",
			fmt.Errorf("db query: %w", err),
		)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, app_errors.NewNotFoundError(
			"UserRepository.FindByEmail: user not found",
			fmt.Errorf("email %s: %w", email, adapters.ErrNoRows),
		)
	}

	var user domain.User
	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, app_errors.NewInternalError(
			"UserRepository.FindByEmail: failed to scan user",
			fmt.Errorf("rows scan: %w", err),
		)
	}
	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) *app_errors.Error {
	if err := r.db.Delete(ctx, "users", id); err != nil {

		if err == adapters.ErrNoRows {
			return app_errors.NewNotFoundError(
				"UserRepository.Delete: user not found",
				fmt.Errorf("id %s: %w", id, err),
			)
		}
		return app_errors.NewInternalError(
			"UserRepository.Delete: failed to delete user",
			fmt.Errorf("db delete: %w", err),
		)
	}
	return nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]domain.User, *app_errors.Error) {
	rows, err := r.db.FindAll(ctx, "users")

	if err != nil {
		return nil, app_errors.NewInternalError(
			"UserRepository.FindAll: failed to execute query",
			fmt.Errorf("db find all: %w", err),
		)
	}
	defer rows.Close()

	var users []domain.User

	if !rows.Next() {
		return nil, nil
	}

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, app_errors.NewInternalError(
				"UserRepository.FindAll: failed to scan user",
				fmt.Errorf("rows scan: %w", err),
			)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, app_errors.NewInternalError(
			"UserRepository.FindAll: rows error",
			fmt.Errorf("rows err: %w", err),
		)
	}

	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, *app_errors.Error) {
	rows, err := r.db.FindByID(ctx, "users", id)
	if err != nil {
		return nil, app_errors.NewInternalError(
			"UserRepository.FindByID: failed to execute query",
			fmt.Errorf("db find by id: %w", err),
		)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, app_errors.NewNotFoundError(
			"UserRepository.FindByID: user not found",
			fmt.Errorf("id %s: %w", id, adapters.ErrNoRows),
		)
	}

	var user domain.User
	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, app_errors.NewInternalError(
			"UserRepository.FindByID: failed to scan user",
			fmt.Errorf("rows scan: %w", err),
		)
	}

	return &user, nil
}