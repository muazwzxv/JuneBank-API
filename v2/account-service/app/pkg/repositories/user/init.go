package user

import (
	"account-service/app/pkg/core/ports"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

var _ ports.IUserRepository = (*userRepo)(nil)

func New(db *sqlx.DB) ports.IUserRepository {
	return &userRepo{
		db: db,
	}
}
