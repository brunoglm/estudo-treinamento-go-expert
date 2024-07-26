package math

type Math struct {
	a int
	b int
	c int
}

func NewMath(a, b, c int) *Math {
	return &Math{
		a: a,
		b: b,
		c: c,
	}
}
