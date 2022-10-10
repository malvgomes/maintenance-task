package queue

type Producer interface {
	Publish(queue string, message []byte) error
}

func GetProducer() (Producer, error) {
	return newRabbitmqProducer()
}
