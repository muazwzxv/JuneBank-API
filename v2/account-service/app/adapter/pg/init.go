package pg

import (
	"account-service/app/adapter"

	"github.com/jmoiron/sqlx"
)

type PgAdapter struct {
	db  *sqlx.DB
	dsn string
}

var _ adapter.IDatabaseAdapter = (*PgAdapter)(nil)

func New(dsn string) adapter.IDatabaseAdapter {
	return &PgAdapter{dsn: dsn}
}
