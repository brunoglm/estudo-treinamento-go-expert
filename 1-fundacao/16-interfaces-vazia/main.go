package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "text"
	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("Tipo da var: %T\n", t)
}
