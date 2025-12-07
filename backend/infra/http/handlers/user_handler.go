package handlers

import (
	"net/http"
	"sub-watch/application/usecases"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	listUsersUseCase *usecases.ListUsersUseCase
	getUserUseCase   *usecases.GetUserUseCase
}

func NewUserHandler(listUsersUseCase *usecases.ListUsersUseCase, getUserUseCase *usecases.GetUserUseCase) *UserHandler {
	return &UserHandler{
		listUsersUseCase: listUsersUseCase,
		getUserUseCase:   getUserUseCase,
	}
}

// ListUsers godoc
// @Summary List all users
// @Description Get a list of all registered users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} domain.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (h *UserHandler) ListUsers(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.listUsersUseCase.Execute(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get a single user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} usecases.UserOutput
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")

	user, err := h.getUserUseCase.Execute(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if user == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
	}

	return c.JSON(http.StatusOK, user)
}
