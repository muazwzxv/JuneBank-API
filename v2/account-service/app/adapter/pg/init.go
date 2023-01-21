package pg

import (
	"account-service/app/adapter"

	"github.com/jmoiron/sqlx"
)

type Adapter struct {
	db  *sqlx.DB
	dsn string
}

func New(dsn string) adapter.IDatabaseAdapter {
	return &Adapter{dsn: dsn}
}
