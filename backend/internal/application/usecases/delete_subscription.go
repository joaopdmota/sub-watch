package usecases

import (
	"context"
	"sub-watch-backend/internal/application"
	app_errors "sub-watch-backend/internal/application/errors"
)

type DeleteSubscriptionUseCase struct {
	subRepo application.SubscriptionRepository
	log     application.Logger
}

func NewDeleteSubscriptionUseCase(subRepo application.SubscriptionRepository, log application.Logger) *DeleteSubscriptionUseCase {
	return &DeleteSubscriptionUseCase{subRepo: subRepo, log: log}
}

func (u *DeleteSubscriptionUseCase) Execute(ctx context.Context, id string) *app_errors.Error {
	if err := u.subRepo.Delete(ctx, id); err != nil {
		u.log.Error("DeleteSubscriptionUseCase.Execute: failed to delete subscription", "error", err)
		return err
	}
	return nil
}
