package configuration

import (
	"fmt"
	. "github.com/AugustoKlaic/golearningstack/pkg/queue/apachekafka"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
	"sync"
)

var (
	adminClient       *kafka.AdminClient
	configInitOnce    sync.Once
	kafkaConfigLogger = log.New(os.Stdout, "CONFIGURATION: ", log.Ldate|log.Ltime|log.Lshortfile)
	TopicName         = "message-api-topic"
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
	InitializeConsumer()
}

func CreateTopic() {
	if _, err := adminClient.CreateTopics(nil, []kafka.TopicSpecification{
		{
			Topic:             TopicName,
			NumPartitions:     Props.Kafka.NumPartitions,
			ReplicationFactor: Props.Kafka.ReplicationFactor,
		},
	}, nil); err != nil {
		kafkaConfigLogger.Fatalf("failed to create topic %s: %v", TopicName, err)
	}
}

func CloseKafkaResources() {
	if adminClient != nil {
		adminClient.Close()
	}
	CloseProducer()
	CloseConsumer()
}
