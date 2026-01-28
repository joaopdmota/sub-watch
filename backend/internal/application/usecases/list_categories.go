package usecases

import (
	"context"
"sub-watch-backend/internal/application"
	app_errors "sub-watch-backend/internal/application/errors"
)

type ListCategoriesUseCase struct {
	repo application.CategoryRepository
}

func NewListCategoriesUseCase(repo application.CategoryRepository) *ListCategoriesUseCase {
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
