package server

import "github.com/jmoiron/sqlx"

type IServer interface {
	Start() error
	Setup() error
	Routes() error
}

type IDatabaseAdapter interface {
	open(dsn string) error
	GetDB() *sqlx.DB
}
