package handlers

import (
	"net/http"
	"sub-watch-backend/internal/application/usecases"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	loginUseCase *usecases.AuthLoginUseCase
}

func NewAuthHandler(loginUseCase *usecases.AuthLoginUseCase) *AuthHandler {
	return &AuthHandler{
		loginUseCase: loginUseCase,
	}
}

// Login godoc
// @Summary Login
// @Description Authenticate a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body usecases.AuthLoginInput true "Login credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var input usecases.AuthLoginInput
	if err := c.Bind(&input); err != nil {
		return err
	}

	userID, appErr := h.loginUseCase.Execute(c.Request().Context(), input)
	if appErr != nil {
		return c.JSON(appErr.Code, appErr)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userID,
	})
}
