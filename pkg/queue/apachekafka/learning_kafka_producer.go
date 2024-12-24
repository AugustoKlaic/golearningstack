package apachekafka

import (
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
	"sync"
)

// https://github.com/confluentinc/confluent-kafka-go/blob/master/examples/producer_example/producer_example.go

var (
	producer            *kafka.Producer
	producerInitOnce    sync.Once
	kafkaProducerLogger = log.New(os.Stdout, "KAFKA_PRODUCER: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func InitializeProducer() {
	producerInitOnce.Do(func() {
		var err error
		producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": configuration.GetKafkaBroker()})
		if err != nil {
			kafkaProducerLogger.Fatalf("failed to create producer: %v", err)
		}
		go kafkaAckListener()
	})
}

func PublishMessage(topic, message string) {
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	if err != nil {
		kafkaProducerLogger.Printf("failed to produce message: %v", err)
	}
}

func kafkaAckListener() {
	e := <-producer.Events()
	switch ev := e.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			kafkaProducerLogger.Printf("delivery failed: %v", ev.TopicPartition.Error)
		}
		kafkaProducerLogger.Printf("Message delivered to %v\n", ev.TopicPartition)
	case kafka.Error:
		kafkaProducerLogger.Printf("Error: %v\n", ev)
	default:
		kafkaProducerLogger.Printf("Ignored event: %s\n", ev)
	}
}

func CloseProducer() {
	if producer != nil {
		producer.Close()
	}
}
