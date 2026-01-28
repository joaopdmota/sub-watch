package usecases

import (
	"context"
"sub-watch-backend/internal/application"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
)

type GetSubscriptionUseCase struct {
	subRepo application.SubscriptionRepository
}

func NewGetSubscriptionUseCase(subRepo application.SubscriptionRepository) *GetSubscriptionUseCase {
	return &GetSubscriptionUseCase{subRepo: subRepo}
}

func (u *GetSubscriptionUseCase) Execute(ctx context.Context, id string) (*domain.Subscription, *app_errors.Error) {
	return u.subRepo.FindByID(ctx, id)
}
