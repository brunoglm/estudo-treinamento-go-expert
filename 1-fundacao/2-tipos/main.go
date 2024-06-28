package main

type ID int

var (
	a int
	b string
	c float64
	d bool
	e ID = 1
)

func main() {
	println(e)
}
