package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && c > a || c == 5 {
		println(a)
	} else {
		println(b)
	}

	switch {
	case a > b:
		println("a > b")
	case b > a:
		println("b > a")
	case b == a:
		println("b == a")
	default:
		panic("invalid")
	}

	switch a {
	case 1:
		println("1")
	case 2:
		println("2")
	}
}
