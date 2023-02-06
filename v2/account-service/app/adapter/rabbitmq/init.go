package rabbitmq

import (
	"account-service/app/adapter"

	"github.com/wagslane/go-rabbitmq"
)

type RabbitAdapter struct {
	Conn *rabbitmq.Conn
}

func New() adapter.IEventQueueAdapter {
	return &RabbitAdapter{}
}
