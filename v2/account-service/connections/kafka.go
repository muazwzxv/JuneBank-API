package connections

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

type KafkaInstance struct {
	KafkaConsumer *KafkaConsumer
	KafkaProducer *KafkaProducer
}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

type KafkaProducer struct {
	Producer *kafka.Producer
}

var kafkaEvents KafkaInstance

func GetKafkaConsumer(log *log.Logger) *KafkaConsumer {
	if kafkaEvents.KafkaConsumer.Consumer == nil {
		createKafkaConsumer(log)
	}
	return kafkaEvents.KafkaConsumer
}

func createKafkaConsumer(log *log.Logger) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:19091",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("create kafka consumer: %v", err)
	}

	defer func(consumer *kafka.Consumer) {
		err := consumer.Close()
		if err != nil {
			panic(err)
		}
	}(consumer)
}

func GetKafkaProducer(log *log.Logger) *KafkaProducer {
	if kafkaEvents.KafkaProducer.Producer == nil {
		createKafkaProducer(log)
	}
	return kafkaEvents.KafkaProducer
}

func createKafkaProducer(log *log.Logger) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:19091",
	})

	if err != nil {
		log.Fatalf("create kafka producer Error: %v", err)
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

	kafkaEvents.KafkaProducer.Producer = producer
}

func CreateConsumer() {
	// TODO: Create a kafka connection as consumer
}
