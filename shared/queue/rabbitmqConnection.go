package queue

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func open() *amqp.Connection {
	connection, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	for err != nil {
		log.Println("Connection to RabbitMQ is not yet ready. Trying again")
		time.Sleep(time.Second * 5)
		connection, err = amqp.Dial("amqp://guest:guest@rabbitmq:5672")
	}

	return connection
}
