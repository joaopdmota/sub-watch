package repositories

import (
	"context"
	"fmt"
	"sub-watch-microservice/application/domain"
	"sub-watch-microservice/infra/database"
)

type UserRepository struct {
	db database.Database
}

func NewUserRepository(db database.Database) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) FindOne(ctx context.Context, id string) (*domain.User, error) {
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

func (r *UserRepository) Insert(ctx context.Context, user *domain.User) error {
	const q = `
		INSERT INTO users (id, name, email, password_hash, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, q,
		user.ID, user.Name, user.Email, user.PasswordHash, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	const q = "SELECT * FROM users WHERE email = $1"

	rows, err := r.db.QueryContext(ctx, q, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	var user domain.User
	if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return &user, rows.Err()
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	return r.db.Delete(ctx, "users", id)
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