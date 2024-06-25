package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func (c *Cliente) Desativar() {
	c.Ativo = false
}

func main() {
	cliente1 := Cliente{
		Nome:  "nome1",
		Idade: 20,
		Ativo: true,
	}
	cliente1.Desativar()

	fmt.Printf("Nome: %s Idade: %d Ativo: %t\n", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
}
