package events

import (
	"account-service/app/adapter"
	"log"

	"github.com/wagslane/go-rabbitmq"
)

type Publisher struct {
	*rabbitmq.Publisher
	Log *log.Logger
}

func NewPublisher(rbmq adapter.IRabbitAdapter, log *log.Logger) (*Publisher, error) {
	conn, err := rbmq.GetConn()
	if err != nil {
		return &Publisher{}, err
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName("account-service"),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)

	if err != nil {
		return &Publisher{}, err
	}

	log.Printf("rabbitmq: publisher created")

	return &Publisher{
		Publisher: publisher,
		Log:       log,
	}, nil
}
