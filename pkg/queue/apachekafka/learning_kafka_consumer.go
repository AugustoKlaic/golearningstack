package apachekafka

import (
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
	"sync"
)

// https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/consumer_example/consumer_example.go

var (
	groupId             = "learning_kafka_golang"
	consumer            *kafka.Consumer
	consumerInitOnce    sync.Once
	kafkaConsumerLogger = log.New(os.Stdout, "KAFKA_CONSUMER: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func InitializeConsumer() {
	consumerInitOnce.Do(func() {
		var err error
		consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": configuration.GetKafkaBroker(),
			"group.id":          groupId,
			"auto.offset.reset": "earliest",
		})
		if err != nil {
			kafkaConsumerLogger.Fatalf("failed to create consumer: %v", err)
		}
	})
}

func ConsumeMessages(topic string, handler func(message *kafka.Message) error) {
	err := consumer.Subscribe(topic, nil)
	if err != nil {
		kafkaConsumerLogger.Fatalf("failed to subscribe to consumer: %v", err)
	}

	go func() {
		for {
			msg, err := consumer.ReadMessage(-1)
			if err != nil {
				kafkaConsumerLogger.Printf("error while consuming: %v", err)
				continue
			}
			if err := handler(msg); err != nil {
				kafkaConsumerLogger.Printf("Error processing message: %v", err)
			}
		}
	}()
}

func CloseConsumer() {
	if consumer != nil {
		_ = consumer.Close()
	}
}
