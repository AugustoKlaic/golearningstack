package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func StartConsumer(queueName, consumerTag string, rabbitConn *amqp091.Connection, handler func(msg []byte) error) {
	if channel, err := rabbitConn.Channel(); err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	} else {
		messages, err := channel.Consume(
			queueName,
			consumerTag,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatalf("Erro ao consumir mensagens: %v", err)
		}

		go func() {
			for msg := range messages {
				if err := handler(msg.Body); err != nil {
					log.Printf("Erro ao processar mensagem: %v", err)
				}
			}
		}()
	}

}
