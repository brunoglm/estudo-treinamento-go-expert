package rabbitmq

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConsumer struct {
	URL             string
	Queue           string
	ConsumerTag     string
	ReconnectDelay  time.Duration
	MaxRetryBackoff time.Duration
	ResetConnection chan bool
	conn            *amqp.Connection
	ch              *amqp.Channel
}

func NewRabbitMQConsumer() *RabbitMQConsumer {
	return &RabbitMQConsumer{
		URL:             "amqp://guest:guest@localhost:5672/",
		Queue:           "go-expert",
		ConsumerTag:     "go-consumer",
		ResetConnection: make(chan bool),
	}
}

func (consumer *RabbitMQConsumer) OpenConnect() error {
	if consumer.conn != nil && !consumer.conn.IsClosed() {
		fmt.Println("A conexão já está aberta")
		return nil
	}
	fmt.Println("Abrindo Conexão")
	conn, err := amqp.Dial(consumer.URL)
	if err != nil {
		return err
	}
	consumer.conn = conn
	fmt.Println("Conexão aberta")
	return nil
}

func (consumer *RabbitMQConsumer) AssertInfraConsumer() error {
	_, err := consumer.ch.QueueDeclare(
		consumer.Queue,
		true,
		false,
		false,
		false,
		nil)
	return err
}

func (consumer *RabbitMQConsumer) CloseConnect() {
	if consumer.conn != nil && !consumer.conn.IsClosed() {
		fmt.Println("Fechando Conexão")
		err := consumer.conn.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("Conexão Fechada")
		return
	}
	fmt.Println("A Conexão já está fechada")
}

func (consumer *RabbitMQConsumer) HandleListenersReconnect() {
	connListener := consumer.conn.NotifyClose(make(chan *amqp.Error))
	chListener := consumer.ch.NotifyClose(make(chan *amqp.Error))
	select {
	case err := <-connListener:
		fmt.Println("close conn iniciado. Err: ", err.Error())
		consumer.ResetConnection <- true
	case err := <-chListener:
		fmt.Println("close ch iniciado. Err: ", err.Error())
		consumer.ResetConnection <- true
	}
}

func (consumer *RabbitMQConsumer) HandleReconnection() {
	// processo de retry e backoff
	fmt.Println("iniciando processo de retry e backoff")
	time.Sleep(15 * time.Second)
}

func (consumer *RabbitMQConsumer) OpenChannel() error {
	err := consumer.OpenConnect()
	if err != nil {
		return err
	}

	if consumer.ch != nil && !consumer.ch.IsClosed() {
		fmt.Println("O canal já está aberto")
		return nil
	}

	fmt.Println("Abrindo Canal")
	ch, err := consumer.conn.Channel()
	if err != nil {
		consumer.CloseConnect()
		return err
	}
	consumer.ch = ch
	fmt.Println("Canal Aberto")
	return nil
}

func (consumer *RabbitMQConsumer) Consume(out chan<- amqp.Delivery, chError chan<- error) {
	msgs, err := consumer.ch.Consume(
		consumer.Queue,
		consumer.ConsumerTag,
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
