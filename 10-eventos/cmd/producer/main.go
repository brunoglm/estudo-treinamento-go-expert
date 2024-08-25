package main

import (
	"encoding/json"
	"fmt"
	"trabalhando-com-eventos/pkg/rabbitmq"
)

func main() {
	producer := rabbitmq.NewRabbitMQPublisher()

	err := producer.OpenChannel()
	if err != nil {
		fmt.Println("error 1: ", err.Error())
	}

	err = producer.AssertInfra()
	if err != nil {
		fmt.Println("error 2: ", err.Error())
	}

	producer.ActivateConfirms()

	type Body struct {
		Message string
	}
	for i := 0; i < 10; i++ {
		b := Body{Message: fmt.Sprintf("Message: %d", i)}
		bJson, _ := json.Marshal(b)
		err = producer.PublishAndWaitConfirms(string(bJson))
		if err != nil {
			fmt.Println("error 3: ", err.Error())
		}
	}
}
