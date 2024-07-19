package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type message struct {
	id  int64
	msg string
}

func main() {
	ch1 := make(chan message)
	ch2 := make(chan message)
	var i int64

	//RabbitMQ por exemplo
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			ch1 <- message{i, "Message from RabbitMQ"}
			time.Sleep(time.Second * 4)
		}
	}()

	//Kafka por exemplo
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			ch2 <- message{i, "Message from Kafka"}
			time.Sleep(time.Second * 4)
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println("msg ch1: ", msg.msg, "id: ", msg.id)
		case msg := <-ch2:
			fmt.Println("msg ch2: ", msg.msg, "id: ", msg.id)
		}
	}
}
