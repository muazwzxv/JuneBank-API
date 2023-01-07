package server

import "github.com/jmoiron/sqlx"

type IServer interface {
	Start() error
	Setup() error
	Routes(db *sqlx.DB) error
	Shutdown() error
}
