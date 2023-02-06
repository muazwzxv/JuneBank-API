package adapter

import (
	"github.com/jmoiron/sqlx"
	"github.com/wagslane/go-rabbitmq"
)

type IDatabaseAdapter interface {
	GetDB() (*sqlx.DB, error)
}

type IEventQueueAdapter interface {
	GetConn() (*rabbitmq.Conn, error)
}
