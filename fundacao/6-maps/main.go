package main

import "fmt"

func main() {
	salarios := map[string]int{
		"nome1": 10,
		"nome2": 20,
	}
	fmt.Println(salarios["nome1"])
	delete(salarios, "nome1")
	fmt.Println(salarios["nome1"])
	salarios["nome3"] = 30
	fmt.Println(salarios["nome3"])

	for k, v := range salarios {
		fmt.Printf("k: %s V: %d\n", k, v)
	}

	sal := make(map[string]int)
	sal1 := map[string]int{}
	fmt.Println(sal)
	fmt.Println(sal1)
}
