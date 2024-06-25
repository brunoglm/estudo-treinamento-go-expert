package main

import "fmt"

type Pessoa interface { // interface em go tem somente assinatura de métodos, e não de atributos também
	Desativar()
}

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
	Desativar(&cliente1)
	fmt.Printf("Nome: %s Idade: %d Ativo: %t\n", cliente1.Nome, cliente1.Idade, cliente1.Ativo)
}

func Desativar(cliente Pessoa) {
	cliente.Desativar()
}
