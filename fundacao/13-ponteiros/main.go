package main

import "fmt"

func main() {
	// Memória -> Endereço -> Valor
	a := 10
	b := &a

	fmt.Println(a, &a)
	fmt.Println(*b, b)
}
