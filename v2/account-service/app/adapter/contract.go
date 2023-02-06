package adapter

import (
	"github.com/jmoiron/sqlx"
	"github.com/wagslane/go-rabbitmq"
)

type IDatabaseAdapter interface {
	GetDB() (*sqlx.DB, error)
}

type IRabbitAdapter interface {
	GetConn() (*rabbitmq.Conn, error)
}
