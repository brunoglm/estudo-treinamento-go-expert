package main

import "fmt"

func main() {
	var array [3]int // usando array, as posições são fixas, começando do 0
	array[0] = 10
	array[1] = 20
	array[2] = 30

	fmt.Println("Dados do array inicial:")
	fmt.Println(len(array)) // tamanho 3
	fmt.Println(cap(array)) // capacidade 3
	for i, v := range array {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

	testeAlterandoValorSemSerPorPonteiro(array)

	fmt.Println("Dados do array após testeAlterandoValorSemSerPorPonteiro:")
	fmt.Println(len(array)) // tamanho 3
	fmt.Println(cap(array)) // capacidade 3
	for i, v := range array {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

	testeAlterandoValorPorPonteiro(&array)

	fmt.Println("Dados do array após testeAlterandoValorPorPonteiro:")
	fmt.Println(len(array)) // tamanho 3
	fmt.Println(cap(array)) // capacidade 3
	for i, v := range array {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}
}

func testeAlterandoValorPorPonteiro(array *[3]int) {
	array[2] = 50
}

func testeAlterandoValorSemSerPorPonteiro(array [3]int) {
	array[1] = 50
}
