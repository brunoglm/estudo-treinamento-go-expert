package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 12}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0]) // len=0 cap=6

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4]) // len=4 cap=6

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // len=4 cap=4

	s = append(s, 14)
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // len=5 cap=12

	testeAlterandoValorPorPonteiro(s)

	for i, v := range s {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}
}

func testeAlterandoValorPorPonteiro(array []int) {
	array[len(array)-1] = 100
}
