package main

import (
	"auction-go/configuration/database/mongodb"
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
}
