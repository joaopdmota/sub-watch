package repositories

import (
	"context"
	"sub-watch-backend/application/domain"
	app_errors "sub-watch-backend/application/errors"
	"sub-watch-backend/infra/database"
	"sub-watch-backend/infra/database/adapters"
)


type UserRepository struct {
	db database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return UserRepository{db: db}
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
		return &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, *app_errors.Error) {
	const q = "SELECT * FROM users WHERE email = $1"

	rows, err := r.db.QueryContext(ctx, q, email)
	if err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, &app_errors.Error{
			Code:    404,
			Type:    app_errors.ERROR_NOT_FOUND,
			Message: app_errors.ERROR_NOT_FOUND,
		}
	}

	var user domain.User
	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}

	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id string) *app_errors.Error {
	if err := r.db.Delete(ctx, "users", id); err != nil {

		if err == adapters.ErrNoRows {
			return &app_errors.Error{
				Code:    404,
				Type:    app_errors.ERROR_NOT_FOUND,
				Message: app_errors.ERROR_NOT_FOUND,
			}
		}
		return &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}
	return nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]domain.User, *app_errors.Error) {
	rows, err := r.db.FindAll(ctx, "users")

	if err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}
	defer rows.Close()

	var users []domain.User

	if !rows.Next() {
		return nil, nil
	}

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, &app_errors.Error{
				Code:    500,
				Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
				Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
			}
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}

	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, *app_errors.Error) {
	rows, err := r.db.FindByID(ctx, "users", id)
	if err != nil {
		return nil, &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: app_errors.ERROR_BAD_REQUEST,
		}
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, &app_errors.Error{
			Code:    404,
			Type:    app_errors.ERROR_NOT_FOUND,
			Message: app_errors.ERROR_NOT_FOUND,
		}
	}

	var user domain.User
	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}

	return &user, nil
}