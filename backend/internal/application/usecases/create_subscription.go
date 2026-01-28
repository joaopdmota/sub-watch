package usecases

import (
	"context"
	"sub-watch-backend/internal/application"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
	"sub-watch-backend/internal/pkg/date"
	id "sub-watch-backend/internal/pkg/uuid"
	"time"
)

type CreateSubscriptionInput struct {
	UserID          string    `json:"user_id" validate:"required,uuid4"`
	CategoryID      string    `json:"category_id" validate:"required,uuid4"`
	PaymentMethodID string    `json:"payment_method_id" validate:"omitempty,uuid4"`
	ServiceName     string    `json:"service_name" validate:"required,min=2,max=100"`
	Price           float64   `json:"price" validate:"required,gt=0"`
	Currency        string    `json:"currency" validate:"required,len=3"`
	Cycle           string    `json:"cycle" validate:"required,oneof=MONTHLY YEARLY"`
	NextBillingDate time.Time `json:"next_billing_date" validate:"required"`
	Notes           string    `json:"notes" validate:"max=500"`
}

type CreateSubscriptionUseCase struct {
	subRepo  application.SubscriptionRepository
	uuid     id.UuidProvider
	date     date.DateProvider
	validator application.Validator
	log       application.Logger
}

func NewCreateSubscriptionUseCase(
	subRepo application.SubscriptionRepository,
	uuid id.UuidProvider,
	date date.DateProvider,
	validator application.Validator,
	log application.Logger,
) *CreateSubscriptionUseCase {
	return &CreateSubscriptionUseCase{
		subRepo:   subRepo,
		uuid:      uuid,
		date:      date,
		validator: validator,
		log:       log,
	}
}

func (u *CreateSubscriptionUseCase) Execute(ctx context.Context, input CreateSubscriptionInput) *app_errors.Error {
	if err := u.validator.ValidateStruct(input); err != nil {
		return &app_errors.Error{
			Code:    400,
			Type:    app_errors.ERROR_BAD_REQUEST,
			Message: "Validation failed",
			Details: u.validator.FormatValidationErrors(err),
		}
	}

	subscription := &domain.Subscription{
		ID:              u.uuid.NewID(),
		UserID:          input.UserID,
		CategoryID:      input.CategoryID,
		PaymentMethodID: input.PaymentMethodID,
		ServiceName:     input.ServiceName,
		Price:           input.Price,
		Currency:        input.Currency,
		Cycle:           input.Cycle,
		NextBillingDate: input.NextBillingDate,
		Status:          domain.SubscriptionStatusActive,
		Notes:           input.Notes,
		CreatedAt:       u.date.Now(),
		UpdatedAt:       u.date.Now(),
	}

	if err := u.subRepo.Insert(ctx, subscription); err != nil {
		u.log.Error("CreateSubscriptionUseCase.Execute: failed to insert subscription", "error", err)
		return err
	}

	return nil
}
