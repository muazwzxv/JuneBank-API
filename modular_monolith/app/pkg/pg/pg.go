package pg

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func (a *Adapter) GetDB() (*sqlx.DB, error) {
	if a.db == nil {
		err := a.open(a.dsn)
		if err != nil {
			return nil, err
		}
	}

	return a.db, nil
}

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
	log.Printf("Database is connected: %v", a.db)

	return nil
}
