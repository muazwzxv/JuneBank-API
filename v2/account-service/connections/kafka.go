package connections

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

type KafkaInstance struct {
	Consumer *kafka.Consumer
	Producer *kafka.Producer
}

var kafkaEvents KafkaInstance

//var Producer *kafka.Producer

func GetKafkaProducer(log *log.Logger) *kafka.Producer {
	if kafkaEvents.Producer == nil {
		createKafkaProducer(log)
	}
	return kafkaEvents.Producer
}

func createKafkaProducer(log *log.Logger) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:19091",
	})

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	// Logger report for produced messages
	go func() {
		for event := range producer.Events() {
			switch e := event.(type) {
			case *kafka.Message:
				if e.TopicPartition.Error != nil {
					log.Printf("Delivery Success: %v", e.TopicPartition)
				} else {
					log.Printf("Delivery failed: %v", e.TopicPartition.Error)
				}
			}
		}
	}()

	kafkaEvents.Producer = producer
}

func CreateConsumer() {
	// TODO: Create a kafka connection as consumer
}
