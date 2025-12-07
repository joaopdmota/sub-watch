package repositories

import (
	"context"
	"sub-watch/application/domain"
	"sub-watch/infra/database"
)

type UserRepository struct {
	db database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	rows, err := r.db.FindAll(ctx, "users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	rows, err := r.db.FindByID(ctx, "users", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil // User not found
	}

	var user domain.User
	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, rows.Err()
}
