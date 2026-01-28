package usecases

import (
	"context"
"sub-watch-backend/internal/application"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
)

type ListSubscriptionsUseCase struct {
	subRepo application.SubscriptionRepository
}

func NewListSubscriptionsUseCase(subRepo application.SubscriptionRepository) *ListSubscriptionsUseCase {
	return &ListSubscriptionsUseCase{subRepo: subRepo}
}

func (u *ListSubscriptionsUseCase) Execute(ctx context.Context, userID string) ([]domain.Subscription, *app_errors.Error) {
	return u.subRepo.FindByUserID(ctx, userID)
}
