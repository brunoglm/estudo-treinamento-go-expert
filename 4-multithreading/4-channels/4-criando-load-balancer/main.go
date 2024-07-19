package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)

	for i := 1; i <= runtime.NumCPU(); i++ {
		go worker(i, data)
	}

	for i := 0; i < 10; i++ {
		data <- i
	}
}
