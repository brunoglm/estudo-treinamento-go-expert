package main

func main() {
	for i := 0; i < 10; i++ {
		println("i: ", i)
	}

	numeros := []int{10, 20, 30}
	for i, numero := range numeros {
		println("Indice: ", i, "numero: ", numero)
	}

	a := 0
	for a < 20 {
		println("a: ", a)
		a++
	}

	x := 0
	for {
		println("loop infinito: ", x)
		if x == 100 {
			break
		}
		x++
	}
}
