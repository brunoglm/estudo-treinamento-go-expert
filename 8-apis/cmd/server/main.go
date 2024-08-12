package main

import (
	"apis/configs"
	"fmt"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	fmt.Println(configs)
}
