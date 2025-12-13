package repositories

import (
	"context"
	"sub-watch-backend/application/domain"
	app_errors "sub-watch-backend/application/errors"
	"sub-watch-backend/infra/database"
	"sub-watch-backend/infra/database/adapters"
)

type CategoryRepository struct {
	db database.Database
}

func NewCategoryRepository(db database.Database) CategoryRepository {
	return CategoryRepository{db: db}
}

func (r *CategoryRepository) Insert(ctx context.Context, category *domain.Category) *app_errors.Error {
	const q = `
		INSERT INTO categories (id, name, icon, color, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.ExecContext(ctx, q,
		category.ID, category.Name, category.Icon, category.Color, category.CreatedAt,
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

func (r *CategoryRepository) Delete(ctx context.Context, id string) *app_errors.Error {
	if err := r.db.Delete(ctx, "categories", id); err != nil {

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

func (r *CategoryRepository) FindAll(ctx context.Context) ([]domain.Category, *app_errors.Error) {
	rows, err := r.db.FindAll(ctx, "categories")

	if err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}
	defer rows.Close()

	var categories []domain.Category

	if !rows.Next() {
		return nil, nil
	}

	for rows.Next() {
		var category domain.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Icon, &category.Color, &category.CreatedAt); err != nil {
			return nil, &app_errors.Error{
				Code:    500,
				Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
				Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
			}
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}

	return categories, nil
}

func (r *CategoryRepository) FindByID(ctx context.Context, id string) (*domain.Category, *app_errors.Error) {
	rows, err := r.db.FindByID(ctx, "categories", id)
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

	var category domain.Category
	if err := rows.Scan(&category.ID, &category.Name, &category.Icon, &category.Color, &category.CreatedAt); err != nil {
		return nil, &app_errors.Error{
			Code:    500,
			Type:    app_errors.ERROR_INTERNAL_SERVER_ERROR,
			Message: app_errors.ERROR_INTERNAL_SERVER_ERROR,
		}
	}

	return &category, nil
}