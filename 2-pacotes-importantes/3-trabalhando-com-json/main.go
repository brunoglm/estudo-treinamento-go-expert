package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	c := Conta{Numero: 10, Saldo: 20}
	res, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	println("com a var na mão: ", string(res))

	err = json.NewEncoder(os.Stdout).Encode(c)
	if err != nil {
		panic(err)
	}

	var contax Conta
	jsonPuro := []byte(`{"numero":55,"saldo":22}`)
	err = json.Unmarshal(jsonPuro, &contax)
	if err != nil {
		panic(err)
	}
	fmt.Println("json.Unmarshal contaX: ", contax)
}
