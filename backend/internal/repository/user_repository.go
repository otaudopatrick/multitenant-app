package repository

import (
	"backend/internal/database"
	"backend/internal/domain"
	"context"
	"database/sql"
	"time"
)

type UserRepository struct {
	db database.Service
}

func NewUserRepository(db database.Service) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateWithTx(ctx context.Context, tx *sql.Tx, user *domain.User) error {
	query := `
        INSERT INTO users (tenant_id, name, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := tx.QueryRowContext(
		ctx,
		query,
		user.Tenant.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}
