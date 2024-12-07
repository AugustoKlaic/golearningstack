package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
)

func PublishMessage(exchange, routingKey string, body []byte, rabbitConn *amqp091.Connection) error {
	if channel, err := rabbitConn.Channel(); err != nil {
		return err
	} else {
		defer channel.Close()

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
			log.Printf("Erro ao publicar mensagem: %v", err)
			return err
		}
	}

	log.Printf("Mensagem publicada: %s", body)
	return nil
}
