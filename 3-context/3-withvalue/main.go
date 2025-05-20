package main

import (
	"context"
	"fmt"
)

func main() {
	// context é imutável
	ctx := context.WithValue(context.Background(), "token", "senha123")
	ctx = context.WithValue(ctx, "userID", 42)
	ctx = context.WithValue(ctx, "role", "admin")

	fmt.Println(ctx.Value("token"))
	fmt.Println(ctx.Value("userID"))
	fmt.Println(ctx.Value("role"))
}
