package main

import (
	"errors"
	"fmt"
)

func main() {
	var a interface{} = 10
	var b interface{} = "text"
	var c interface{} = true

	showType(a)
	showType(b)
	showType(c)

	println("afirmando que é um int: ", a.(int))

	res, ok := c.(int)
	println("res: ", res)
	println("ok: ", ok)
}

func showType(t interface{}) {
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
