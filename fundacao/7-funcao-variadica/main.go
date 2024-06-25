package main

import "fmt"

func main() {
	a := []int{10, 20, 30}
	fmt.Println("First")
	fmt.Println(sum(1, 2))
	fmt.Println("Second")
	fmt.Println(sum(a...))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
