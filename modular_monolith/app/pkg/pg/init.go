package pg

import (
	"account-service/app/pkg"

	"github.com/jmoiron/sqlx"
)

type Adapter struct {
	db  *sqlx.DB
	dsn string
}

func New(dsn string) pkg.IDatabaseAdapter {
	return &Adapter{dsn: dsn}
}
