package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go task("task1", wg)
	go task("task2", wg)

	wg.Wait()
}
