package handlers

import (
	"net/http"
	"sub-watch-backend/internal/application/usecases"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	listUsersUseCase  *usecases.ListUsersUseCase
	getUserUseCase    *usecases.GetUserUseCase
	createUserUseCase *usecases.CreateUserUseCase
}

func NewUserHandler(listUsersUseCase *usecases.ListUsersUseCase, getUserUseCase *usecases.GetUserUseCase, createUserUseCase *usecases.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		listUsersUseCase:  listUsersUseCase,
		getUserUseCase:    getUserUseCase,
		createUserUseCase: createUserUseCase,
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
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": users,
	})
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
		return c.JSON(err.Code, err)
	}


	return c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body usecases.UserInput true "User object"
// @Success 200 {object} usecases.UserOutput
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	var userInput usecases.UserInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.createUserUseCase.Execute(ctx, userInput); err != nil {
		return c.JSON(err.Code, err)
	}

	return c.NoContent(http.StatusCreated)
}