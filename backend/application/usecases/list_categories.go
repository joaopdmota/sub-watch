package usecases

import (
	"context"
	app_errors "sub-watch-backend/application/errors"
	"sub-watch-backend/infra/repositories"
)

type ListCategoriesUseCase struct {
	repo repositories.CategoryRepository
}

func NewListCategoriesUseCase(repo repositories.CategoryRepository) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{repo: repo}
}

func (u *ListCategoriesUseCase) Execute(ctx context.Context) ([]CategoryOutput, *app_errors.Error) {
	categories, err := u.repo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var output []CategoryOutput
	for _, category := range categories {
		output = append(output, CategoryOutput{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			Icon:      category.Icon,
			Color:     category.Color,
		})
	}
	return output, nil
}
