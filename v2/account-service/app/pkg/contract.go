package pkg

import "github.com/jmoiron/sqlx"

type IDatabaseAdapter interface {
	GetDB() (*sqlx.DB, error)
}
