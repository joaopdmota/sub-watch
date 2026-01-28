package handlers

import (
	"net/http"
	"sub-watch-backend/internal/application/usecases"

	"github.com/labstack/echo/v4"
)

type SubscriptionHandler struct {
	createSubUseCase *usecases.CreateSubscriptionUseCase
	listSubsUseCase  *usecases.ListSubscriptionsUseCase
	getSubUseCase    *usecases.GetSubscriptionUseCase
	updateSubUseCase *usecases.UpdateSubscriptionUseCase
	deleteSubUseCase *usecases.DeleteSubscriptionUseCase
}

func NewSubscriptionHandler(
	createSubUseCase *usecases.CreateSubscriptionUseCase,
	listSubsUseCase *usecases.ListSubscriptionsUseCase,
	getSubUseCase *usecases.GetSubscriptionUseCase,
	updateSubUseCase *usecases.UpdateSubscriptionUseCase,
	deleteSubUseCase *usecases.DeleteSubscriptionUseCase,
) *SubscriptionHandler {
	return &SubscriptionHandler{
		createSubUseCase: createSubUseCase,
		listSubsUseCase:  listSubsUseCase,
		getSubUseCase:    getSubUseCase,
		updateSubUseCase: updateSubUseCase,
		deleteSubUseCase: deleteSubUseCase,
	}
}

// CreateSubscription godoc
// @Summary      Create a new subscription
// @Description  Creates a new subscription for the authenticated user
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription body usecases.CreateSubscriptionInput true "Subscription creation data"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  app_errors.Error
// @Failure      500  {object}  app_errors.Error
// @Router       /api/subscriptions [post]
func (h *SubscriptionHandler) CreateSubscription(c echo.Context) error {
	var input usecases.CreateSubscriptionInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// In a real scenario, we'd get the user ID from the JWT token
	// For now, let's assume it's in the input or extracted from context
	
	if err := h.createSubUseCase.Execute(c.Request().Context(), input); err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Subscription created successfully"})
}

// ListSubscriptions godoc
// @Summary      List subscriptions
// @Description  Returns all subscriptions for a given user (requires ?user_id=)
// @Tags         subscriptions
// @Produce      json
// @Param        user_id query string true "User ID"
// @Success      200  {array}   domain.Subscription
// @Failure      500  {object}  app_errors.Error
// @Router       /api/subscriptions [get]
func (h *SubscriptionHandler) ListSubscriptions(c echo.Context) error {
	userID := c.QueryParam("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "user_id is required"})
	}

	subs, err := h.listSubsUseCase.Execute(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, subs)
}

// GetSubscription godoc
// @Summary      Get a subscription
// @Description  Returns details of a specific subscription
// @Tags         subscriptions
// @Produce      json
// @Param        id   path      string  true  "Subscription ID"
// @Success      200  {object}  domain.Subscription
// @Failure      404  {object}  app_errors.Error
// @Router       /api/subscriptions/{id} [get]
func (h *SubscriptionHandler) GetSubscription(c echo.Context) error {
	id := c.Param("id")
	sub, err := h.getSubUseCase.Execute(c.Request().Context(), id)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, sub)
}

// UpdateSubscription godoc
// @Summary      Update a subscription
// @Description  Updates existing subscription details
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Subscription ID"
// @Param        subscription body usecases.UpdateSubscriptionInput true "Subscription update data"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  app_errors.Error
// @Router       /api/subscriptions/{id} [put]
func (h *SubscriptionHandler) UpdateSubscription(c echo.Context) error {
	id := c.Param("id")
	var input usecases.UpdateSubscriptionInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	input.ID = id

	if err := h.updateSubUseCase.Execute(c.Request().Context(), input); err != nil {
		return c.JSON(err.Code, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Subscription updated successfully"})
}

// DeleteSubscription godoc
// @Summary      Delete a subscription
// @Description  Removes a subscription
// @Tags         subscriptions
// @Produce      json
// @Param        id   path      string  true  "Subscription ID"
// @Success      204
// @Failure      500  {object}  app_errors.Error
// @Router       /api/subscriptions/{id} [delete]
func (h *SubscriptionHandler) DeleteSubscription(c echo.Context) error {
	id := c.Param("id")
	if err := h.deleteSubUseCase.Execute(c.Request().Context(), id); err != nil {
		return c.JSON(err.Code, err)
	}

	return c.NoContent(http.StatusNoContent)
}
