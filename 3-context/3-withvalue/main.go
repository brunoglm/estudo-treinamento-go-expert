package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "token", "senha123")
	fmt.Println(ctx.Value("token"))
}
