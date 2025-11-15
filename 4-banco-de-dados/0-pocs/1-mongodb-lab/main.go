package main

import (
	"context"
	"fmt"
	"log"
	"mongodb-lab/config"
	"mongodb-lab/database"
	"mongodb-lab/entity"
	"mongodb-lab/repository/user"
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

	userCollection := client.Database(config.DatabaseName).Collection(config.UserCollectionName)

	userDatabase := user.NewStore(userCollection)

	log.Println("Iniciando operações de usuário...")

	appCtx := context.Background()

	newUser := &entity.User{
		Name:  "João Silva",
		Email: "email@gmail.com",
	}

	userID, err := userDatabase.CreateUser(appCtx, newUser)
	if err != nil {
		log.Fatalf("Erro ao criar usuário: %v", err)
	}
	log.Printf("Usuário criado com ID: %s", userID.Hex())

	retrievedUser, err := userDatabase.GetUserByID(appCtx, userID)
	if err != nil {
		log.Fatalf("Erro ao buscar usuário: %v", err)
	}
	log.Printf("Usuário recuperado: %+v", retrievedUser)

	updatedCount, err := userDatabase.UpdateUserName(appCtx, userID, "João Pereira")
	if err != nil {
		log.Fatalf("Erro ao atualizar nome do usuário: %v", err)
	}
	log.Printf("Número de documentos atualizados: %d", updatedCount)

	updatedUser, err := userDatabase.GetUserByID(appCtx, userID)
	if err != nil {
		log.Fatalf("Erro ao buscar usuário atualizado: %v", err)
	}
	log.Printf("Usuário atualizado: %+v", updatedUser)

	deleteCount, err := userDatabase.DeleteUser(appCtx, userID)
	if err != nil {
		log.Fatalf("Erro ao deletar usuário: %v", err)
	}
	log.Printf("Número de documentos deletados: %d", deleteCount)

	_, err = userDatabase.GetUserByID(appCtx, userID)
	if err != nil {
		log.Printf("Confirmação de deleção: %v", err)
	} else {
		log.Fatalf("Erro: Usuário ainda existe após deleção")
	}

	log.Println("Operações de usuário concluídas com sucesso.")
}
