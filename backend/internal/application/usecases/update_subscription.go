package usecases

import (
	"context"
	"sub-watch-backend/internal/application"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
	"sub-watch-backend/internal/pkg/date"
	"time"
)

type UpdateSubscriptionInput struct {
	ID              string    `json:"id" validate:"required,uuid4"`
	CategoryID      string    `json:"category_id" validate:"required,uuid4"`
	PaymentMethodID string    `json:"payment_method_id" validate:"omitempty,uuid4"`
	ServiceName     string    `json:"service_name" validate:"required,min=2,max=100"`
	Price           float64   `json:"price" validate:"required,gt=0"`
	Currency        string    `json:"currency" validate:"required,len=3"`
	Cycle           string    `json:"cycle" validate:"required,oneof=MONTHLY YEARLY"`
	NextBillingDate time.Time `json:"next_billing_date" validate:"required"`
	Status          string    `json:"status" validate:"required,oneof=ACTIVE PAUSED CANCELED"`
	Notes           string    `json:"notes" validate:"max=500"`
}

type UpdateSubscriptionUseCase struct {
	subRepo   application.SubscriptionRepository
	date      date.DateProvider
	validator application.Validator
	log       application.Logger
}

func NewUpdateSubscriptionUseCase(
	subRepo application.SubscriptionRepository,
	date date.DateProvider,
	validator application.Validator,
	log application.Logger,
) *UpdateSubscriptionUseCase {
	return &UpdateSubscriptionUseCase{
		subRepo:   subRepo,
		date:      date,
		validator: validator,
		log:       log,
	}
}

func (u *UpdateSubscriptionUseCase) Execute(ctx context.Context, input UpdateSubscriptionInput) *app_errors.Error {
	if err := u.validator.ValidateStruct(input); err != nil {
		return &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "Validation failed",
			Details: u.validator.FormatValidationErrors(err),
		}
	}

	sub, err := u.subRepo.FindByID(ctx, input.ID)
	if err != nil {
		return err
	}

	sub.CategoryID = input.CategoryID
	sub.PaymentMethodID = input.PaymentMethodID
	sub.ServiceName = input.ServiceName
	sub.Price = input.Price
	sub.Currency = input.Currency
	sub.Cycle = input.Cycle
	sub.NextBillingDate = input.NextBillingDate
	sub.Status = domain.SubscriptionStatus(input.Status)
	sub.Notes = input.Notes
	sub.UpdatedAt = u.date.Now()

	if err := u.subRepo.Update(ctx, sub); err != nil {
		u.log.Error("UpdateSubscriptionUseCase.Execute: failed to update subscription", "error", err)
		return err
	}

	return nil
}
