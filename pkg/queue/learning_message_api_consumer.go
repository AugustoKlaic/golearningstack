package queue

import (
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"github.com/AugustoKlaic/golearningstack/pkg/queue/rabbitmq"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	"github.com/AugustoKlaic/golearningstack/pkg/utils"
)

type MessageApiConsumer struct {
	service service.LearningServiceInterface
}

func NewMessageApiConsumer(service service.LearningServiceInterface) *MessageApiConsumer {
	return &MessageApiConsumer{
		service: service,
	}
}

func (c *MessageApiConsumer) Consume() {
	var rabbitConn = configuration.GetConnection(configuration.GetRabbitMQURL())
	rabbitmq.StartConsumer(
		configuration.QueueName,
		"message-api",
		rabbitConn,
		c.processMessage,
	)
}

func (c *MessageApiConsumer) processMessage(msg []byte) error {
	var receivedMessage *entity.MessageEntity
	utils.JsonDecoder(msg, &receivedMessage)

	return c.service.DeleteMessage(receivedMessage.Id)
}
