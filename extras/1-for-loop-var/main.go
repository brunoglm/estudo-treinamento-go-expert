package main

import "fmt"

func main() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for range values {
		<-done
	}

	fmt.Println("<-- Separador -->")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("<-- Separador -->")

	for i := range 10 {
		fmt.Println(i)
	}
}
