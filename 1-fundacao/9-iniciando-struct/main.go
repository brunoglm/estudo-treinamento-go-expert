package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	cliente1 := Cliente{
		Nome:  "nome1",
		Idade: 20,
		Ativo: true,
	}
	cliente1.Ativo = false

	fmt.Printf("Nome: %s Idade: %d Ativo: %t\n", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
	fmt.Printf("Ponteiro de cliente1: %p", &cliente1)
}
