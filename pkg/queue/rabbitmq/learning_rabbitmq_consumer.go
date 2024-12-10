package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

var rabbitConsumerLogger = log.New(os.Stdout, "RABBIT_CONSUMER: ", log.Ldate|log.Ltime|log.Lshortfile)

func StartConsumer(queueName, consumerTag string, rabbitConn *amqp091.Connection, handler func(msg []byte) error) {
	if channel, err := rabbitConn.Channel(); err != nil {
		rabbitConsumerLogger.Fatalf("Failed to open a channel: %s", err)
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
			rabbitConsumerLogger.Fatalf("Error consuming menssages: %v", err)
		}

		go func() {
			for msg := range messages {
				if err := handler(msg.Body); err != nil {
					rabbitConsumerLogger.Printf("Error processing message: %v", err)
				}
			}
		}()
	}
}
