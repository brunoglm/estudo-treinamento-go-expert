package main

import (
	"fmt"
	"trabalhando-com-eventos/pkg/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	consumer := rabbitmq.NewRabbitMQConsumer()

outerbreak:
	for {
		fmt.Println("passando no for externo")
		err := consumer.OpenChannel()
		if err != nil {
			fmt.Println("error 1: ", err.Error())
			continue
		}

		go consumer.HandleListenersReconnect()

		err = consumer.AssertInfraConsumer()
		if err != nil {
			fmt.Println("error 2: ", err.Error())
			continue
		}

		chMsgs := make(chan amqp.Delivery)
		chError := make(chan error)
		go consumer.Consume(chMsgs, chError)

		for {
			fmt.Println("passando no for interno")
			select {
			case msg := <-chMsgs:
				fmt.Println("processando a mensagem: ", string(msg.Body))
				msg.Ack(false)
			case err := <-chError:
				fmt.Println("error 3: ", err.Error())
				consumer.HandleReconnection()
				continue outerbreak
			case <-consumer.ResetConnection:
				consumer.HandleReconnection()
				continue outerbreak
			}
		}
	}
}
