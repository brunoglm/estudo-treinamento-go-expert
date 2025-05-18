package main

import (
	"errors"
	"fmt"
)

func main() {
	var a any = 10
	var b any = "text"
	var c any = true

	showType(a)
	showType(b)
	showType(c)

	println("afirmando que é um int: ", a.(int))

	res, ok := c.(int)
	println("res: ", res)
	println("ok: ", ok)
}

func showType(t any) {
	switch v := t.(type) {
	case int:
		fmt.Printf("Tipo é string e seu valor é: %d\n", v)
	case string:
		fmt.Printf("Tipo é string e seu valor é: %s\n", v)
	case bool:
		fmt.Printf("Tipo é bool e seu valor é: %t\n", v)
	default:
		panic(errors.New("Tipo inválido"))
	}
}
