package main

const a = "texto"

var b bool // por padrão Go infere o valor padrão, nesse caso false

var ( // podemos declarar N vars junto
	c, d int
	e    string
	f    float64
	g    string = "definindo e atribuindo valor"
)

func main() {
	b = true
	c := "valor" // usando shorthand
	println(a)
	println(b)
	println(c)
}
