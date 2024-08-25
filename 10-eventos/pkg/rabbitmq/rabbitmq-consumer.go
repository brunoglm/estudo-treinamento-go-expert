package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	RabbitMQClient
	ConsumerTag string
}

func NewRabbitMQConsumer() *RabbitMQConsumer {
	rabbitMQConsumer := RabbitMQConsumer{
		ConsumerTag: "go-consumer",
	}
	rabbitMQConsumer.URL = "amqp://guest:guest@localhost:5672/"
	rabbitMQConsumer.ResetConnection = make(chan bool)
	rabbitMQConsumer.Queue = "go-expert"
	rabbitMQConsumer.exName = "ex-go-expert"
	rabbitMQConsumer.key = "key-test"

	return &rabbitMQConsumer
}

func (client *RabbitMQConsumer) Consume(out chan<- amqp.Delivery, chError chan<- error) {
	msgs, err := client.ch.Consume(
		client.Queue,
		client.ConsumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		chError <- err
	}
	for msg := range msgs {
		out <- msg
	}
	fmt.Println("Encerrando o consume")
}
