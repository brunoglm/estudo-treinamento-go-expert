package main

func soma(a, b *int) int {
	*a *= 2
	return *a + *b
}

func main() {
	a := 10
	b := 30
	result := soma(&a, &b)
	println("result: ", result)
	println("a: ", a)
}
