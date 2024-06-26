package main

type Conta struct {
	Saldo int
}

func NewConta(saldo int) *Conta {
	return &Conta{Saldo: saldo}
}

func (c Conta) Simular(valor int) int {
	c.Saldo += valor
	return c.Saldo
}

func (c *Conta) Creditar(valor int) int {
	c.Saldo += valor
	return c.Saldo
}

func main() {
	c := NewConta(100)
	c.Simular(200)
	println("saldo após simulação: ", c.Saldo)
	c.Creditar(200)
	println("saldo após credito: ", c.Saldo)
}
