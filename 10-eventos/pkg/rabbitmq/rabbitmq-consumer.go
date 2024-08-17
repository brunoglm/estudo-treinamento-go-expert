package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	RabbitMQClient
	Queue       string
	ConsumerTag string
}

func NewRabbitMQConsumer() *RabbitMQConsumer {
	rabbitMQConsumer := RabbitMQConsumer{
		Queue:       "go-expert",
		ConsumerTag: "go-consumer",
	}
	rabbitMQConsumer.URL = "amqp://guest:guest@localhost:5672/"
	rabbitMQConsumer.ResetConnection = make(chan bool)

	return &rabbitMQConsumer
}

func (client *RabbitMQConsumer) AssertInfraConsumer() error {
	_, err := client.ch.QueueDeclare(
		client.Queue,
		true,
		false,
		false,
		false,
		nil)
	return err
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
