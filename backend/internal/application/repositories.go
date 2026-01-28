package application

import (
	"context"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
)

type UserRepository interface {
	Insert(ctx context.Context, user *domain.User) *app_errors.Error
	FindByEmail(ctx context.Context, email string) (*domain.User, *app_errors.Error)
	FindByID(ctx context.Context, id string) (*domain.User, *app_errors.Error)
	FindAll(ctx context.Context) ([]domain.User, *app_errors.Error)
	Delete(ctx context.Context, id string) *app_errors.Error
}

type CategoryRepository interface {
	Insert(ctx context.Context, category *domain.Category) *app_errors.Error
	FindByID(ctx context.Context, id string) (*domain.Category, *app_errors.Error)
	FindAll(ctx context.Context) ([]domain.Category, *app_errors.Error)
	Delete(ctx context.Context, id string) *app_errors.Error
}

type SubscriptionRepository interface {
	Insert(ctx context.Context, sub *domain.Subscription) *app_errors.Error
	Update(ctx context.Context, sub *domain.Subscription) *app_errors.Error
	FindByID(ctx context.Context, id string) (*domain.Subscription, *app_errors.Error)
	FindByUserID(ctx context.Context, userID string) ([]domain.Subscription, *app_errors.Error)
	Delete(ctx context.Context, id string) *app_errors.Error
}

type PaymentMethodRepository interface {
	Insert(ctx context.Context, pm *domain.PaymentMethod) *app_errors.Error
	FindByID(ctx context.Context, id string) (*domain.PaymentMethod, *app_errors.Error)
	FindAll(ctx context.Context) ([]domain.PaymentMethod, *app_errors.Error)
	Delete(ctx context.Context, id string) *app_errors.Error
}
