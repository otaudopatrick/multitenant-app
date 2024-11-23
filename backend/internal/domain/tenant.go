package domain

import "time"

type Tenant struct {
	ID        int64
	Name      string
	Subdomain string
	CreatedAt time.Time
	UpdatedAt time.Time
}
