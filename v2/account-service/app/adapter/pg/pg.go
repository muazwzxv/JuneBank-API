package pg

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func (a *Adapter) open(dsn string) error {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *Adapter) GetDB() *sqlx.DB {
	if a.db == nil {
		a.open(a.dsn)
	}

	return a.db
}
