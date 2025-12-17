package handlers

import (
	"net/http"
	"sub-watch-backend/application/usecases"

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
