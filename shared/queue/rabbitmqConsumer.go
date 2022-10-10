package queue

import (
	"context"
	"log"

	"github.com/streadway/amqp"
)

type rabbitmqConsumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func (r *rabbitmqConsumer) Consume(
	ctx context.Context,
	queue string,
	callback func(ctx context.Context, message []byte),
) error {
	_, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	messages, err := r.channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Println("Message received")
			callback(ctx, message.Body)
		}
	}()

	<-forever

	return nil
}

func newRabbitmqConsumer() (Consumer, error) {
	connection := open()
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to RabbitMQ at port :5672")
	log.Println("Waiting for messages")

	return &rabbitmqConsumer{
		connection: connection,
		channel:    channel,
	}, nil
}
