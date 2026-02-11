package main

import "fmt"

func add(a, b int) int {
	sum := a + b // sum é alocada na stack, pois é uma variável local e não é retornada ou referenciada fora da função add
	return sum   // liberada quando a função termina
}

func main() {
	result := add(3, 4)
	fmt.Println("Resultado:", result)
}

// go build -gcflags="-m" main.go

// O comando acima compila o programa Go e exibe informações sobre
// as decisões de alocação de memória feitas pelo compilador.
// A flag -gcflags="-m" instrui o compilador a mostrar detalhes sobre as otimizações de escape,
// indicando se as variáveis são alocadas na stack ou no heap.
// Isso é útil para entender como o Go gerencia a memória e otimiza o desempenho do programa.
