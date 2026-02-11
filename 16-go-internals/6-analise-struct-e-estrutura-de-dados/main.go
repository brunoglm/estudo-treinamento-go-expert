package main

// type User struct {
// 	name string
// }

// func NewUser(name string) *User {
// 	return &User{name: name} // Alocação no heap, pois o endereço é retornado e pode ser referenciado fora da função
// }

func storeInMap() map[string]*int {
	m := make(map[string]*int)
	i := 42
	m["key"] = &i // Alocação no heap, pois o endereço de i é armazenado no mapa que é retornado
	return m
}

func main() {
	mapa := storeInMap()
	println("User:", mapa)
}

// go build -gcflags="-m" main.go

// O comando acima compila o programa Go e exibe informações sobre
// as decisões de alocação de memória feitas pelo compilador.
// A flag -gcflags="-m" instrui o compilador a mostrar detalhes sobre as otimizações de escape,
// indicando se as variáveis são alocadas na stack ou no heap.
// Isso é útil para entender como o Go gerencia a memória e otimiza o desempenho do programa.
