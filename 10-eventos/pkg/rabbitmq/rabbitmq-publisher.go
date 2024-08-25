package rabbitmq

import (
	"errors"
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	RabbitMQClient
	confirms chan amqp.Confirmation
}

func NewRabbitMQPublisher() *RabbitMQPublisher {
	rabbitMQPublisher := RabbitMQPublisher{}
	rabbitMQPublisher.URL = "amqp://guest:guest@localhost:5672/"
	rabbitMQPublisher.ResetConnection = make(chan bool)
	rabbitMQPublisher.Queue = "go-expert"
	rabbitMQPublisher.exName = "ex-go-expert"
	rabbitMQPublisher.key = "key-test"

	return &rabbitMQPublisher
}

func (client *RabbitMQPublisher) PublishAndWaitConfirms(body string) error {
	err := client.ch.Publish(
		client.exName,
		client.key,
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

	// Aguarda a confirmação
	select {
	case confirm := <-client.confirms:
		if confirm.Ack {
			fmt.Println("Message confirmed")
			return nil
		} else {
			fmt.Println("Message failed")
			return errors.New("message failed")
		}
	case <-time.After(5 * time.Second):
		fmt.Println("No confirmation received, message might not be delivered")
		return errors.New("no confirmation received, message might not be delivered")
	}
}

func (client *RabbitMQPublisher) ActivateConfirms() error {
	// Ativa as confirmações do publisher
	if err := client.ch.Confirm(false); err != nil {
		fmt.Printf("Failed to enable publisher confirms: %v", err)
		return err
	}

	confirms := client.ch.NotifyPublish(make(chan amqp.Confirmation, 1))
	client.confirms = confirms

	return nil
}
