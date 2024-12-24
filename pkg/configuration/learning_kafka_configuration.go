package configuration

import (
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/queue/apachekafka"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
	"sync"
)

// https://github.com/confluentinc/confluent-kafka-go/tree/master/examples

var (
	adminClient       *kafka.AdminClient
	configInitOnce    sync.Once
	kafkaConfigLogger = log.New(os.Stdout, "CONFIGURATION: ", log.Ldate|log.Ltime|log.Lshortfile)
	topicName         = "message-api-topic"
)

func GetKafkaBroker() string {
	return fmt.Sprintf("%s:%s", Props.Kafka.Host, Props.Kafka.Port)
}

func ConfigureKafka() {
	configInitOnce.Do(func() {
		var err error
		adminClient, err = kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": GetKafkaBroker()})
		if err != nil {
			kafkaConfigLogger.Fatalf("failed to create admin client: %v", err)
		}
	})

	CreateTopic()
	InitializeProducer()
}

func CreateTopic() {
	if _, err := adminClient.CreateTopics(nil, []kafka.TopicSpecification{
		{
			Topic:             topicName,
			NumPartitions:     Props.Kafka.NumPartitions,
			ReplicationFactor: Props.Kafka.ReplicationFactor,
		},
	}, nil); err != nil {
		kafkaConfigLogger.Fatalf("failed to create topic %s: %v", topicName, err)
	}
}

func CloseKafkaResources() {
	if adminClient != nil {
		adminClient.Close()
	}
	CloseProducer()
}
