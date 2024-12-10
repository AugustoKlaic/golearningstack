package configuration

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"sync"
)

/*
 - Todo hide connection url details in a separate file
 - Operator "once.Do()" returns a singleton
*/

var rabbitConfigLogger = log.New(os.Stdout, "CONFIGURATION: ", log.Ldate|log.Ltime|log.Lshortfile)

var (
	conn         *amqp091.Connection
	once         sync.Once
	ExchangeName = "message-api-exchange"
	exchangeType = "direct"
	QueueName    = "message-api-queue"
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
			rabbitConfigLogger.Fatalf("Error connecting to RabbitMQ: %v", err)
		}
	})

	return conn
}

func ConfigureRabbitMQ(conn *amqp091.Connection) {
	channel, err := conn.Channel()
	if err != nil {
		rabbitConfigLogger.Fatalf("Error creating channel: %v", err)
	}

	setupExchange(channel)
	setupQueue(channel)
	bindQueueExchange(channel)
}

func CloseConnection() {
	if conn != nil {
		rabbitConfigLogger.Println("Closing connection to RabbitMQ")
		channel, _ := conn.Channel()
		_ = channel.Close()
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
		rabbitConfigLogger.Fatalf("Error declaring exchange: %v", err)
	}
}

func setupQueue(channel *amqp091.Channel) {
	_, err := channel.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		rabbitConfigLogger.Fatalf("Error declaring queue: %v", err)
	}
}

func bindQueueExchange(channel *amqp091.Channel) {
	err := channel.QueueBind(
		QueueName,
		RoutingKey,
		ExchangeName,
		false,
		nil,
	)
	if err != nil {
		rabbitConfigLogger.Fatalf("Error linking queue to exchange: %v", err)
	}

	rabbitConfigLogger.Println("Exchange and queue successfully bind!")
}
