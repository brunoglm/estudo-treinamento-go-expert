package main

import "fmt"

func main() {
	var x any = 10
	var y interface{} = "text"
	showType(x)
	showType(y)
}

func showType(t any) {
	fmt.Printf("Tipo da var: %T\n", t)
}
