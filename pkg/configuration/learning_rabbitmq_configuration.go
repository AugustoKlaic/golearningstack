package configuration

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

/*
 - Todo hide connection url details in a separate file
 - Operator "once.Do()" returns a singleton
*/

var (
	conn         *amqp091.Connection
	once         sync.Once
	ExchangeName = "message-api-exchange"
	exchangeType = "direct"
	queueName    = "message-api-queue"
	RoutingKey   = "message-api"
)

func GetRabbitMQURL() string {
	return "amqp://guest:guest@localhost:5672"
}

func GetConnection(url string) *amqp091.Connection {
	once.Do(func() {
		var err error
		conn, err = amqp091.Dial(url)
		if err != nil {
			log.Fatalf("Erro ao conectar ao RabbitMQ: %v", err)
		}
	})

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("erro ao criar canal: %v", err)
	}
	defer channel.Close()

	setupExchange(channel)
	setupQueue(channel)
	bindQueueExchange(channel)

	return conn
}

func CloseConnection() {
	if conn != nil {
		_ = conn.Close()
	}
}

func setupExchange(channel *amqp091.Channel) {
	err := channel.ExchangeDeclare(
		ExchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("erro ao declarar exchange: %v", err)
	}
}

func setupQueue(channel *amqp091.Channel) {
	_, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("erro ao declarar fila: %v", err)
	}
}

func bindQueueExchange(channel *amqp091.Channel) {
	err := channel.QueueBind(
		queueName,
		RoutingKey,
		ExchangeName,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("erro ao vincular fila à exchange: %v", err)
	}

	log.Println("Exchange e fila configuradas com sucesso!")
}
