package main

import "fmt"

type ID int

var a ID = 1

func main() {
	fmt.Printf("O tipo de E é %T", a)
}
