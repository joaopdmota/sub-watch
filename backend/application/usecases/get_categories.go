package usecases

import (
	"context"
	app_errors "sub-watch-backend/application/errors"
	"sub-watch-backend/infra/repositories"
	"time"
)

type GetCategoryUseCase struct {
	repo  repositories.CategoryRepository
}

type CategoryOutput struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Color     string `json:"color"`
	CreatedAt time.Time `json:"created_at"`
}

func NewGetCategoryUseCase(repo repositories.CategoryRepository) *GetCategoryUseCase {
	return &GetCategoryUseCase{repo: repo}
}

func (u *GetCategoryUseCase) Execute(ctx context.Context, id string) (*CategoryOutput, *app_errors.Error) {
	category, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &CategoryOutput{
		ID:        category.ID,
		Name:      category.Name,
		Icon:      category.Icon,
		Color:     category.Color,
		CreatedAt: category.CreatedAt,
	}, nil
}