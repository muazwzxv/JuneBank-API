package pg

import (
	"account-service/app/adapter"
	"log"

	"github.com/jmoiron/sqlx"
)

type PgAdapter struct {
	db  *sqlx.DB
	dsn string
	Log *log.Logger
}

var _ adapter.IDatabaseAdapter = (*PgAdapter)(nil)

func New(dsn string, log *log.Logger) adapter.IDatabaseAdapter {
	return &PgAdapter{
		dsn: dsn,
		Log: log,
	}
}
