package repository

import (
	"backend/internal/database"
)

type UserRepository struct {
	db database.Service
}

func NewUserRepository(db database.Service) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
