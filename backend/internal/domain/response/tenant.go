package response

import (
	"backend/internal/domain"
	"time"
)

type TenantResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Subdomain string    `json:"subdomain"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTenantResponse(tenant domain.Tenant) TenantResponse {
	return TenantResponse{
		ID:        tenant.ID,
		Name:      tenant.Name,
		Subdomain: tenant.Subdomain,
		CreatedAt: tenant.CreatedAt,
		UpdatedAt: tenant.UpdatedAt,
	}
}
