package rabbitmq

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	URL             string
	ResetConnection chan bool
	conn            *amqp.Connection
	ch              *amqp.Channel
	Queue           string
	exName          string
	key             string
}

func (client *RabbitMQClient) OpenConnect() error {
	if client.conn != nil && !client.conn.IsClosed() {
		fmt.Println("A conexão já está aberta")
		return nil
	}
	fmt.Println("Abrindo Conexão")
	conn, err := amqp.Dial(client.URL)
	if err != nil {
		return err
	}
	client.conn = conn
	fmt.Println("Conexão aberta")
	return nil
}

func (client *RabbitMQClient) CloseConnect() {
	if client.conn != nil && !client.conn.IsClosed() {
		fmt.Println("Fechando Conexão")
		err := client.conn.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("Conexão Fechada")
		return
	}
	fmt.Println("A Conexão já está fechada")
}

func (client *RabbitMQClient) HandleListenersReconnect() {
	connListener := client.conn.NotifyClose(make(chan *amqp.Error))
	chListener := client.ch.NotifyClose(make(chan *amqp.Error))
	select {
	case err := <-connListener:
		fmt.Println("close conn iniciado. Err: ", err.Error())
		client.ResetConnection <- true
	case err := <-chListener:
		fmt.Println("close ch iniciado. Err: ", err.Error())
		client.ResetConnection <- true
	}
}

func (client *RabbitMQClient) AssertInfra() error {
	_, err := client.ch.QueueDeclare(
		client.Queue,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}

	err = client.ch.ExchangeDeclare(
		client.exName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = client.ch.QueueBind(
		client.Queue,
		client.key,
		client.exName,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (client *RabbitMQClient) HandleReconnection() {
	// processo de retry e backoff
	fmt.Println("iniciando processo de retry e backoff")
	time.Sleep(15 * time.Second)
}

func (client *RabbitMQClient) OpenChannel() error {
	err := client.OpenConnect()
	if err != nil {
		return err
	}

	if client.ch != nil && !client.ch.IsClosed() {
		fmt.Println("O canal já está aberto")
		return nil
	}

	fmt.Println("Abrindo Canal")
	ch, err := client.conn.Channel()
	if err != nil {
		client.CloseConnect()
		return err
	}

	// Defina o prefetch para 10 mensagens
	err = ch.Qos(
		10,    // Prefetch count
		0,     // Prefetch size (0 desabilita o limite baseado no tamanho)
		false, // Global (false define apenas para este canal)
	)
	if err != nil {
		client.CloseConnect()
		return err
	}

	client.ch = ch
	fmt.Println("Canal Aberto")
	return nil
}
