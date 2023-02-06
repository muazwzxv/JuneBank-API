package rabbitmq

import (
	"account-service/app/adapter"
	"log"

	"github.com/wagslane/go-rabbitmq"
)

type RabbitAdapter struct {
	Conn *rabbitmq.Conn
	Log  *log.Logger
}

func New(log *log.Logger) adapter.IRabbitAdapter {
	return &RabbitAdapter{
		Log: log,
	}
}
