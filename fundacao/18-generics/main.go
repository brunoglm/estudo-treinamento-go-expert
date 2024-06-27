package main

func soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

type Number interface {
	~int | ~float64
}

func soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func compara[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	a := map[string]int{
		"1": 10,
		"2": 20,
	}
	b := map[string]float64{
		"1": 10.11,
		"2": 10.11,
	}
	c := map[string]MyNumber{
		"1": 10,
		"2": 12,
	}
	println("soma com int: ", soma(a))
	println("soma com float64: ", soma(b))
	println("soma2 com int: ", soma2(a))
	println("soma2 com float64: ", soma2(b))

	println("soma2 com float64: ", soma2(c))

	println("compara: ", compara(1, 2))
	println("compara: ", compara(2, 2))
	println("compara: ", compara(2, 2.2))
}
