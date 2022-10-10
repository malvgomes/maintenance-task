package queue

import (
	"log"

	"github.com/streadway/amqp"
)

type rabbitmqProducer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func (r *rabbitmqProducer) Publish(queue string, message []byte) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	amqpMessage := amqp.Publishing{
		ContentType: "text/plain",
		Body:        message,
	}

	if err := r.channel.Publish("", queue, false, false, amqpMessage); err != nil {
		return err
	}

	log.Println("Message enqueued successfully to queue:", queue)
	return nil
}

func newRabbitmqProducer() (Producer, error) {
	connection := open()
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to RabbitMQ at port :5672")
	return &rabbitmqProducer{
		connection: connection,
		channel:    channel,
	}, nil
}
