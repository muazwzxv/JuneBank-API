package pg

import "github.com/jmoiron/sqlx"

type Adapter struct {
	db  *sqlx.DB
	dsn string
}

func New(dsn string) *Adapter {
	return &Adapter{dsn: dsn}
}
