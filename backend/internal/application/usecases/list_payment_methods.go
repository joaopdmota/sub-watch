package usecases

import (
	"context"
"sub-watch-backend/internal/application"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
)

type ListPaymentMethodsUseCase struct {
	pmRepo application.PaymentMethodRepository
}

func NewListPaymentMethodsUseCase(pmRepo application.PaymentMethodRepository) *ListPaymentMethodsUseCase {
	return &ListPaymentMethodsUseCase{pmRepo: pmRepo}
}

func (u *ListPaymentMethodsUseCase) Execute(ctx context.Context) ([]domain.PaymentMethod, *app_errors.Error) {
	return u.pmRepo.FindAll(ctx)
}
