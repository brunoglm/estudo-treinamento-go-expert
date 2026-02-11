package main

import "fmt"

func startGoroutine() {
	i := 42
	go func() {
		fmt.Println("Value of i:", i)
	}()
}

func main() {

}

// go build -gcflags="-m" main.go

// O comando acima compila o programa Go e exibe informações sobre
// as decisões de alocação de memória feitas pelo compilador.
// A flag -gcflags="-m" instrui o compilador a mostrar detalhes sobre as otimizações de escape,
// indicando se as variáveis são alocadas na stack ou no heap.
// Isso é útil para entender como o Go gerencia a memória e otimiza o desempenho do programa.
