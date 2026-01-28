package repositories

import (
	"context"
	"fmt"
	"sub-watch-backend/internal/application/domain"
	app_errors "sub-watch-backend/internal/application/errors"
	"sub-watch-backend/internal/infra/database"
	"sub-watch-backend/internal/infra/database/adapters"
)

type PaymentMethodRepository struct {
	db database.Database
}

func NewPaymentMethodRepository(db database.Database) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

func (r *PaymentMethodRepository) Insert(ctx context.Context, pm *domain.PaymentMethod) *app_errors.Error {
	const q = `
		INSERT INTO payment_methods (id, name, type, created_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(ctx, q, pm.ID, pm.Name, pm.Type, pm.CreatedAt)
	if err != nil {
		return app_errors.NewInternalError(
			"PaymentMethodRepository.Insert: failed to execute insert",
			fmt.Errorf("db exec: %w", err),
		)
	}
	return nil
}

func (r *PaymentMethodRepository) FindByID(ctx context.Context, id string) (*domain.PaymentMethod, *app_errors.Error) {
	const q = "SELECT * FROM payment_methods WHERE id = $1"
	rows, err := r.db.QueryContext(ctx, q, id)
	if err != nil {
		return nil, app_errors.NewInternalError(
			"PaymentMethodRepository.FindByID: failed to execute query",
			fmt.Errorf("db query: %w", err),
		)
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, app_errors.NewNotFoundError(
			"PaymentMethodRepository.FindByID: payment method not found",
			fmt.Errorf("id %s: %w", id, adapters.ErrNoRows),
		)
	}

	var pm domain.PaymentMethod
	if err := rows.Scan(&pm.ID, &pm.Name, &pm.Type, &pm.CreatedAt); err != nil {
		return nil, app_errors.NewInternalError(
			"PaymentMethodRepository.FindByID: failed to scan payment method",
			fmt.Errorf("rows scan: %w", err),
		)
	}

	return &pm, nil
}

func (r *PaymentMethodRepository) FindAll(ctx context.Context) ([]domain.PaymentMethod, *app_errors.Error) {
	rows, err := r.db.FindAll(ctx, "payment_methods")
	if err != nil {
		return nil, app_errors.NewInternalError(
			"PaymentMethodRepository.FindAll: failed to list",
			fmt.Errorf("db find all: %w", err),
		)
	}
	defer rows.Close()

	var pms []domain.PaymentMethod
	for rows.Next() {
		var pm domain.PaymentMethod
		if err := rows.Scan(&pm.ID, &pm.Name, &pm.Type, &pm.CreatedAt); err != nil {
			return nil, app_errors.NewInternalError(
				"PaymentMethodRepository.FindAll: failed to scan payment method",
				fmt.Errorf("rows scan: %w", err),
			)
		}
		pms = append(pms, pm)
	}

	return pms, nil
}

func (r *PaymentMethodRepository) Delete(ctx context.Context, id string) *app_errors.Error {
	if err := r.db.Delete(ctx, "payment_methods", id); err != nil {
		if err == adapters.ErrNoRows {
			return app_errors.NewNotFoundError(
				"PaymentMethodRepository.Delete: payment method not found",
				fmt.Errorf("id %s: %w", id, err),
			)
		}
		return app_errors.NewInternalError(
			"PaymentMethodRepository.Delete: failed to delete payment method",
			fmt.Errorf("db delete: %w", err),
		)
	}
	return nil
}
