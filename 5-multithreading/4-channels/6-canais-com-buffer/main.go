package main

func main() {
	ch := make(chan string, 2)
	ch <- "ola1"
	ch <- "ola2"

	println(<-ch)
	println(<-ch)

}
