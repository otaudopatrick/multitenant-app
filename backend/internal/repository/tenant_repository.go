package repository

import (
	"backend/internal/database"
	"backend/internal/domain"
	"context"
	"database/sql"
	"time"
)

type TenantRepository struct {
	db database.Service
}

func NewTenantRepository(db database.Service) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) CreateWithTx(ctx context.Context, tx *sql.Tx, tenant *domain.Tenant) error {
	query := `
        INSERT INTO tenants (name, subdomain, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	now := time.Now()
	tenant.CreatedAt = now
	tenant.UpdatedAt = now

	err := tx.QueryRowContext(
		ctx,
		query,
		tenant.Name,
		tenant.Subdomain,
		tenant.CreatedAt,
		tenant.UpdatedAt,
	).Scan(&tenant.ID)

	if err != nil {
		return err
	}

	return nil
}
