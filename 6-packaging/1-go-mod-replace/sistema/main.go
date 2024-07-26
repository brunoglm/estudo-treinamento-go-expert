package main

import (
	"fmt"

	"github.com/brunoglm/math"
)

// ajustar o replace no go mod com o comando abaixo
// go mod edit -replace github.com/brunoglm/math=../math
func main() {
	m := math.NewMath(1, 2, 3)
	fmt.Println(m)
}
