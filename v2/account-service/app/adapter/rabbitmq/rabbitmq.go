package rabbitmq

import (
	"github.com/wagslane/go-rabbitmq"
)

func (a *RabbitAdapter) GetConn() (*rabbitmq.Conn, error) {
	if a.Conn == nil {
		if err := a.open(); err != nil {
			return nil, err
		}
	}

	return a.Conn, nil
}

func (a *RabbitAdapter) open() error {

	// Default configs
	conn, err := rabbitmq.NewConn(
		"amqp://guest:guest@localhost",
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		return err
	}
	// TODO: Implement graceful shutdown for RabbitMQ

	a.Conn = conn
	return nil
}
