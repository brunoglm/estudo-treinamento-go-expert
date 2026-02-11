package main

func recursive(n int) int {
	if n <= 0 {
		return 1
	}
	return n * recursive(n-1) // A variável n é alocada na stack para cada chamada recursiva, e cada chamada tem sua própria cópia de n
}

func main() {
	println(recursive(5))
}

// go build -gcflags="-m" main.go
// go install github.com/go-delve/delve/cmd/dlv@latest
// go build -gcflags="all=-N -l" main.go
// dlv debug
// break main.recursive
// continue
// step...step...step...
// stack

// O comando acima compila o programa Go e exibe informações sobre
// as decisões de alocação de memória feitas pelo compilador.
// A flag -gcflags="-m" instrui o compilador a mostrar detalhes sobre as otimizações de escape,
// indicando se as variáveis são alocadas na stack ou no heap.
// Isso é útil para entender como o Go gerencia a memória e otimiza o desempenho do programa.
