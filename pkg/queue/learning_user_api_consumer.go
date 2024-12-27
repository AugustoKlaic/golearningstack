package queue

import (
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"github.com/AugustoKlaic/golearningstack/pkg/queue/apachekafka"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	"github.com/AugustoKlaic/golearningstack/pkg/utils"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"os"
)

var userApiConsumerLogger = log.New(os.Stdout, "USER_API_CONSUMER: ", log.Ldate|log.Ltime|log.Lshortfile)

type UserApiConsumer struct {
	service service.UserCredentialsServiceInterface
}

func NewUserApiConsumer(service service.UserCredentialsServiceInterface) *UserApiConsumer {
	return &UserApiConsumer{
		service: service,
	}
}

func (c *UserApiConsumer) Consume() {
	userApiConsumerLogger.Println("Starting user API consumer...")
	apachekafka.ConsumeMessages(configuration.TopicName, c.processMessage)
}

func (c *UserApiConsumer) processMessage(msg *kafka.Message) error {
	var receivedMessage *entity.UserCredentials
	utils.JsonDecoder(msg.Value, &receivedMessage)
	userApiConsumerLogger.Printf("User created with Id: %s", receivedMessage.Id.String())
	return nil
}
