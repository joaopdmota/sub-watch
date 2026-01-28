package handlers

import (
	"net/http"
	"sub-watch-backend/internal/application/usecases"

	"github.com/labstack/echo/v4"
)

type PaymentMethodHandler struct {
	listPMsUseCase *usecases.ListPaymentMethodsUseCase
}

func NewPaymentMethodHandler(listPMsUseCase *usecases.ListPaymentMethodsUseCase) *PaymentMethodHandler {
	return &PaymentMethodHandler{listPMsUseCase: listPMsUseCase}
}

// ListPaymentMethods godoc
// @Summary      List payment methods
// @Description  Returns all supported payment methods
// @Tags         payment-methods
// @Produce      json
// @Success      200  {array}   domain.PaymentMethod
// @Failure      500  {object}  app_errors.Error
// @Router       /api/payment-methods [get]
func (h *PaymentMethodHandler) ListPaymentMethods(c echo.Context) error {
	pms, err := h.listPMsUseCase.Execute(c.Request().Context())
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, pms)
}
