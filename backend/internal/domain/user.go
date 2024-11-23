package domain

import "time"

type User struct {
	ID        int64
	Tenant    Tenant
	Name      string
	Email     string
	Password  string
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
