package main

import "fmt"

func main() {
	total := func() int {
		return sum(10, 20, 20) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
