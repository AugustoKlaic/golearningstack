package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

var rabbitProducerLogger = log.New(os.Stdout, "RABBIT_PRODUCER: ", log.Ldate|log.Ltime|log.Lshortfile)

func PublishMessage(exchange, routingKey string, body []byte, rabbitConn *amqp091.Connection) {
	if channel, err := rabbitConn.Channel(); err != nil {
		rabbitProducerLogger.Fatalf("Failed to open a channel: %s", err)
	} else {
		if err := channel.Publish(
			exchange,
			routingKey,
			false,
			false,
			amqp091.Publishing{
				ContentType: "application/json",
				Body:        body,
			},
		); err != nil {
			rabbitProducerLogger.Printf("Error publishing message: %v", err)
		}
	}

	rabbitProducerLogger.Printf("Message published successfully: %s", body)
}
