package matematica

type Number interface {
	int | float64
}

func Soma[T Number](a, b T) T {
	return a + b
}
