package main

import "fmt"

func main() {
	forever := make(chan int)

	go func() {
		for i := 0; i < 100000; i++ {
			fmt.Println("I: ", i)
		}
		close(forever)
	}()

	<-forever
}
