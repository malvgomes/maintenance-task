package queue

import (
	"context"
)

type Consumer interface {
	Consume(ctx context.Context, queue string, callback func(context.Context, []byte)) error
}

func GetConsumer() (Consumer, error) {
	return newRabbitmqConsumer()
}
