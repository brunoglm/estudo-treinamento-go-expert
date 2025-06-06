package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10, 12}
	// tamanhos iniciais len=6 cap=6
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// len=0 cap=6
	// capacidade é 6, pois o slice ainda aponta para o array original
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// len=4 cap=6
	// capacidade é 6, pois o slice ainda aponta para o array original
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// len=4 cap=4
	// capacidade é 4, mas o slice aponta para o array original ainda
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 14)
	// len=5 cap=10
	// capacidade é 10, pois foi criado um novo array para o slice, com o dobro do tamanho
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	fmt.Println("O slice original é", s)

	testeAlterandoValorPorPonteiro(s)

	fmt.Println("O slice após testeAlterandoValorPorPonteiro é", s)

	fmt.Println()

	// criando um slice b através de um slice s,
	// podemos ver que o slice b aponta para o mesmo array que o slice s
	b := s[1:]
	fmt.Println("O slice s é", s)
	fmt.Println("O slice b é", b)
	b[0] = 999
	fmt.Println("O slice s é", s)
	fmt.Println("O slice b é", b)

	fmt.Println()

	// dessa vez estamos fazendo uma cópia do slice s, e não apontando para o mesmo array
	// o slice c é uma cópia do slice s, e não aponta para o mesmo array que o slice s
	c := make([]int, len(s[1:]))
	copy(c, s[1:])
	fmt.Println("O slice s é", s)
	fmt.Println("O slice c é", c)
	c[0] = 888
	fmt.Println("O slice s é", s)
	fmt.Println("O slice c é", c)
}

func testeAlterandoValorPorPonteiro(array []int) {
	array[len(array)-1] = 100
}
