package main

import (
	"context"
	"fmt"
	"log"
	"mongodb-lab/config"
	"mongodb-lab/database/mongodb"
	"mongodb-lab/entity"
	"mongodb-lab/repository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	client, err := mongodb.Connect(config)
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

	userDatabase := repository.NewUserRepository(userCollection)

	log.Println("Iniciando operações de usuário...")

	appCtx := context.Background()

	newUser := &entity.User{
		Name:  "João Silva",
		Email: "email@gmail.com",
	}

	userID, err := userDatabase.Create(appCtx, newUser)
	if err != nil {
		log.Fatalf("Erro ao criar usuário: %v", err)
	}
	log.Printf("Usuário criado com ID: %s", userID.Hex())

	retrievedUser, err := userDatabase.GetByID(appCtx, userID)
	if err != nil {
		log.Fatalf("Erro ao buscar usuário: %v", err)
	}
	log.Printf("Usuário recuperado: %+v", retrievedUser)

	updatedCount, err := userDatabase.UpdateName(appCtx, userID, "João Pereira")
	if err != nil {
		log.Fatalf("Erro ao atualizar nome do usuário: %v", err)
	}
	log.Printf("Número de documentos atualizados: %d", updatedCount)

	updatedUser, err := userDatabase.GetByID(appCtx, userID)
	if err != nil {
		log.Fatalf("Erro ao buscar usuário atualizado: %v", err)
	}
	log.Printf("Usuário atualizado: %+v", updatedUser)

	deleteCount, err := userDatabase.Delete(appCtx, userID)
	if err != nil {
		log.Fatalf("Erro ao deletar usuário: %v", err)
	}
	log.Printf("Número de documentos deletados: %d", deleteCount)

	_, err = userDatabase.GetByID(appCtx, userID)
	if err != nil {
		log.Printf("Confirmação de deleção: %v", err)
	} else {
		log.Fatalf("Erro: Usuário ainda existe após deleção")
	}

	log.Println("Operações de usuário concluídas com sucesso.")

	// _ = executeOpsWithTransaction(client, config)
}

func executeOpsWithTransaction(client *mongo.Client, config *config.AppConfig) error {
	userCollection := client.Database(config.DatabaseName).Collection(config.UserCollectionName)

	userDatabase := repository.NewUserRepository(userCollection)

	session, err := client.StartSession()
	if err != nil {
		return fmt.Errorf("Erro ao iniciar sessão: %v", err)
	}
	defer session.EndSession(context.Background())

	callback := func(sessCtx mongo.SessionContext) (any, error) {
		log.Println("[TX] Iniciando transação de usuário...")

		newUser := &entity.User{
			Name:  "João Silva",
			Email: "email@gmail.com",
		}

		userID, err := userDatabase.Create(sessCtx, newUser)
		if err != nil {
			log.Fatalf("[TX] Erro ao criar usuário: %v", err)
		}
		log.Printf("[TX] Usuário criado temporariamente com ID: %s", userID.Hex())

		// simulando rollback
		//log.Println("[TX] Simulando erro para rollback...")
		//return nil, fmt.Errorf("Erro simulado para rollback") // Rollback

		// simulando commit
		log.Println("[TX] Finalizando transação com commit...")

		return "Sucesso", nil // Commit
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := session.WithTransaction(ctx, callback)
	if err != nil {
		log.Printf("Transacção revertida: %v", err)
		return err
	} else {
		log.Printf("Transacção concluída com sucesso: %v", result)
	}

	return nil
}
