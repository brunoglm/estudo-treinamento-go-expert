package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
