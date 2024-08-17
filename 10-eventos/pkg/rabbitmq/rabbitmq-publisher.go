package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQPublisher struct {
	RabbitMQClient
	exName string
}

func NewRabbitMQPublisher() *RabbitMQPublisher {
	rabbitMQPublisher := RabbitMQPublisher{
		exName: "ex-go-expert",
	}
	rabbitMQPublisher.URL = "amqp://guest:guest@localhost:5672/"
	rabbitMQPublisher.ResetConnection = make(chan bool)

	return &rabbitMQPublisher
}

func (client *RabbitMQPublisher) Publish(body string) error {
	err := client.ch.Publish(
		client.exName,
		"",
		true,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
