package repository

import "github.com/jmoiron/sqlx"

type IRepository interface {
	CreateUserData()
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) IRepository {
	return &repository{db}
}
