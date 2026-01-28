package repositories

import (
	"context"
	"fmt"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
	"sub-watch-backend/internal/infra/database"
	"sub-watch-backend/internal/infra/database/adapters"
)

type SubscriptionRepository struct {
	db database.Database
}

func NewSubscriptionRepository(db database.Database) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Insert(ctx context.Context, sub *domain.Subscription) *app_errors.Error {
	const q = `
		INSERT INTO subscriptions (
			id, user_id, category_id, payment_method_id, service_name, 
			price, currency, cycle, next_billing_date, status, notes, 
			created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`
	_, err := r.db.ExecContext(ctx, q,
		sub.ID, sub.UserID, sub.CategoryID, sub.PaymentMethodID, sub.ServiceName,
		sub.Price, sub.Currency, sub.Cycle, sub.NextBillingDate, sub.Status, sub.Notes,
		sub.CreatedAt, sub.UpdatedAt,
	)
	if err != nil {
		return app_errors.NewInternalError(
			"SubscriptionRepository.Insert: failed to execute insert",
			fmt.Errorf("db exec: %w", err),
		)
	}
	return nil
}

func (r *SubscriptionRepository) Update(ctx context.Context, sub *domain.Subscription) *app_errors.Error {
	const q = `
		UPDATE subscriptions SET 
			category_id = $1, payment_method_id = $2, service_name = $3, 
			price = $4, currency = $5, cycle = $6, next_billing_date = $7, 
			status = $8, notes = $9, updated_at = $10
		WHERE id = $11
	`
	_, err := r.db.ExecContext(ctx, q,
		sub.CategoryID, sub.PaymentMethodID, sub.ServiceName,
		sub.Price, sub.Currency, sub.Cycle, sub.NextBillingDate,
		sub.Status, sub.Notes, sub.UpdatedAt, sub.ID,
	)
	if err != nil {
		return app_errors.NewInternalError(
			"SubscriptionRepository.Update: failed to execute update",
			fmt.Errorf("db exec: %w", err),
		)
	}
	return nil
}

func (r *SubscriptionRepository) FindByID(ctx context.Context, id string) (*domain.Subscription, *app_errors.Error) {
	const q = "SELECT * FROM subscriptions WHERE id = $1"
	rows, err := r.db.QueryContext(ctx, q, id)
	if err != nil {
		return nil, app_errors.NewInternalError(
			"SubscriptionRepository.FindByID: failed to execute query",
			fmt.Errorf("db query: %w", err),
		)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, app_errors.NewNotFoundError(
			"SubscriptionRepository.FindByID: subscription not found",
			fmt.Errorf("id %s: %w", id, adapters.ErrNoRows),
		)
	}

	var sub domain.Subscription
	err = rows.Scan(
		&sub.ID, &sub.UserID, &sub.CategoryID, &sub.PaymentMethodID, &sub.ServiceName,
		&sub.Price, &sub.Currency, &sub.Cycle, &sub.NextBillingDate, &sub.Status, &sub.Notes,
		&sub.CreatedAt, &sub.UpdatedAt,
	)
	if err != nil {
		return nil, app_errors.NewInternalError(
			"SubscriptionRepository.FindByID: failed to scan subscription",
			fmt.Errorf("rows scan: %w", err),
		)
	}

	return &sub, nil
}

func (r *SubscriptionRepository) FindByUserID(ctx context.Context, userID string) ([]domain.Subscription, *app_errors.Error) {
	const q = "SELECT * FROM subscriptions WHERE user_id = $1 ORDER BY service_name"
	rows, err := r.db.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, app_errors.NewInternalError(
			"SubscriptionRepository.FindByUserID: failed to execute query",
			fmt.Errorf("db query: %w", err),
		)
	}
	defer rows.Close()

	var subs []domain.Subscription
	for rows.Next() {
		var sub domain.Subscription
		err = rows.Scan(
			&sub.ID, &sub.UserID, &sub.CategoryID, &sub.PaymentMethodID, &sub.ServiceName,
			&sub.Price, &sub.Currency, &sub.Cycle, &sub.NextBillingDate, &sub.Status, &sub.Notes,
			&sub.CreatedAt, &sub.UpdatedAt,
		)
		if err != nil {
			return nil, app_errors.NewInternalError(
				"SubscriptionRepository.FindByUserID: failed to scan subscription",
				fmt.Errorf("rows scan: %w", err),
			)
		}
		subs = append(subs, sub)
	}

	return subs, nil
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id string) *app_errors.Error {
	if err := r.db.Delete(ctx, "subscriptions", id); err != nil {
		if err == adapters.ErrNoRows {
			return app_errors.NewNotFoundError(
				"SubscriptionRepository.Delete: subscription not found",
				fmt.Errorf("id %s: %w", id, err),
			)
		}
		return app_errors.NewInternalError(
			"SubscriptionRepository.Delete: failed to delete subscription",
			fmt.Errorf("db delete: %w", err),
		)
	}
	return nil
}
