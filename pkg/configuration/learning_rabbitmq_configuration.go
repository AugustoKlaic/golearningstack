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
	conn *amqp091.Connection
	once sync.Once
)

func GetRabbitMQURL() string {
	return "amqp://guest:guest@localhost:5672/"
}

func GetConnection(url string) *amqp091.Connection {
	once.Do(func() {
		var err error
		conn, err = amqp091.Dial(url)
		if err != nil {
			log.Fatalf("Erro ao conectar ao RabbitMQ: %v", err)
		}
	})
	return conn
}

func CloseConnection() {
	if conn != nil {
		_ = conn.Close()
	}
}
