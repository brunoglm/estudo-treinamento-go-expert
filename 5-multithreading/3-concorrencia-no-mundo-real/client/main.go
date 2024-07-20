package main

import (
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go doRequest(wg)
	go doRequest(wg)
	wg.Wait()
}

func doRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
		time.Sleep(1 * time.Second)
	}
}
