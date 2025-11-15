package main

import (
	"context"
	"fmt"
	"log"
	"mongodb-lab/config"
	"mongodb-lab/database"
	"time"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	client, err := database.Connect(config)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	defer func() {
		fmt.Println("Database: Desconectando do MongoDB...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("Erro ao desconectar do banco de dados: %v", err)
		}
	}()
}
