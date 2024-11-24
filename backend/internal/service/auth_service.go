package service

import (
	"backend/internal/domain"
	"backend/internal/domain/request"
	"backend/internal/repository"
	"backend/internal/utils"
	"context"
	"database/sql"
	"fmt"
)

type AuthService struct {
	db         *sql.DB
	tenantRepo *repository.TenantRepository
	userRepo   *repository.UserRepository
}

func NewAuthService(db *sql.DB, tenantRepo *repository.TenantRepository, userRepo *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		tenantRepo: tenantRepo,
		db:         db,
	}
}

func (s *AuthService) CreateUserWithTenant(ctx context.Context, req *request.CreateUserWithTenantRequest) (*domain.Tenant, *domain.User, error) {
	tx, err := s.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to start transaction: %v", err)
	}

	defer func(tx *sql.Tx) {
		err := tx.Rollback()
		if err != nil {
			return
		}
	}(tx)

	tenant := &domain.Tenant{
		Name:      req.Tenant.Name,
		Subdomain: req.Tenant.Subdomain,
	}

	err = s.tenantRepo.CreateWithTx(ctx, tx, tenant)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create tenant: %v", err)
	}

	hashedPassword, err := utils.HashPassword(req.User.Password)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to hash password: %v", err)
	}

	user := &domain.User{
		Tenant:   *tenant,
		Name:     req.User.Name,
		Email:    req.User.Email,
		Password: hashedPassword,
	}

	err = s.userRepo.CreateWithTx(ctx, tx, user)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create user: %v", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return tenant, user, nil
}
