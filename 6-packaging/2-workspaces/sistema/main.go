package main

import (
	"fmt"

	"github.com/brunoglm/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2, 3)
	fmt.Println(m)
	fmt.Println(uuid.New().String())
}
