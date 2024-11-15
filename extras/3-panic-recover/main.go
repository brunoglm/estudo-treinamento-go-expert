package main

import "fmt"

func panic1() {
	panic("panic1")
}

func panic2() {
	panic("panic2")
}

func panicGeneral() {
	panic("panicGeneral")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("Panic 1 recovered in main:", r)
				return
			}
			if r == "panic2" {
				fmt.Println("Panic 2 recovered in main:", r)
				return
			}
			fmt.Println("Panic geral recovered in main:", r)
		}
	}()
	panicGeneral()
}
