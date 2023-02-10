package events

import (
	"account-service/app/adapter"
	"log"

	"github.com/wagslane/go-rabbitmq"
)

type Subscriber struct {
	*rabbitmq.Consumer
	Log *log.Logger
}

func NewConsumer(rbmq adapter.IRabbitAdapter, log *log.Logger) (*Subscriber, error) {

	conn, err := rbmq.GetConn()
	if err != nil {
		return &Subscriber{}, err
	}

	// TODO: Setup a proper queue consumer to listen to events
	// Example from Wagslane go-rabbitmq
	consumer, err := rabbitmq.NewConsumer(
		conn,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			log.Printf("consumed: %v", string(d.Body))
			// rabbitmq.Ack, rabbitmq.NackDiscard, rabbitmq.NackRequeue
			return rabbitmq.Ack
		},
		"my_queue",
		rabbitmq.WithConsumerOptionsRoutingKey("my_routing_key"),
		rabbitmq.WithConsumerOptionsExchangeName("events"),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	return &Subscriber{
		Consumer: consumer,
		Log:      log,
	}, nil
}
