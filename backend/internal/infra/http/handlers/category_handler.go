package handlers

import (
	"net/http"
	"sub-watch-backend/internal/application/usecases"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	getCategoryUseCase *usecases.GetCategoryUseCase
	listCategoriesUseCase *usecases.ListCategoriesUseCase
}

func NewCategoryHandler(getCategoryUseCase *usecases.GetCategoryUseCase, listCategoriesUseCase *usecases.ListCategoriesUseCase) *CategoryHandler {
	return &CategoryHandler{
		getCategoryUseCase: getCategoryUseCase,
		listCategoriesUseCase: listCategoriesUseCase,
	}
}

// GetCategory godoc
// @Summary Get a category by ID
// @Description Get a single category by their ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} domain.Category
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	category, err := h.getCategoryUseCase.Execute(ctx, id)
	if err != nil {
		return c.JSON(err.Code, err)
	}


	return c.JSON(http.StatusOK, category)
}

// ListCategories godoc
// @Summary List all categories
// @Description Get a list of all registered categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {array} domain.Category
// @Failure 500 {object} map[string]string
// @Router /categories [get]
func (h *CategoryHandler) ListCategories(c echo.Context) error {
	ctx := c.Request().Context()

	categories, err := h.listCategoriesUseCase.Execute(ctx)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, categories)
}