package main

import "fmt"

func main() {
	var array [3]int // usando array, as posições são fixas, começando do 0
	array[0] = 10
	array[1] = 20
	array[2] = 30
	testeAlterandoValorPorPonteiro(&array)
	fmt.Println(len(array))          // tamanho 3
	fmt.Println(cap(array))          // capacidade 3
	fmt.Println(len(array) - 1)      // index da ultima posição sendo 2
	fmt.Println(array[len(array)-1]) // valor da ultima posição sendo 50

	for i, v := range array {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}
}

func testeAlterandoValorPorPonteiro(array *[3]int) {
	array[2] = 50
}
