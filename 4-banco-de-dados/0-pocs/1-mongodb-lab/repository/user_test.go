package repository_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"mongodb-lab/config"
	mongointernallib "mongodb-lab/database/mongodb"
	"mongodb-lab/entity"
	"mongodb-lab/repository"

	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoClient *mongo.Client
var cfg *config.AppConfig

func TestMain(m *testing.M) {
	var err error
	cfg, err = config.LoadConfig("../.env")
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	ctx := context.Background()

	mongoContainer, err := mongodb.Run(ctx,
		"mongo:7.0",
		mongodb.WithUsername(cfg.DBUser),
		mongodb.WithPassword(cfg.DBPass),
	)
	if err != nil {
		log.Fatalf("Erro ao iniciar container do mongodb: %v", err)
	}

	defer func() {
		log.Println("Finalizando container do mongodb de testes...")
		if err := mongoContainer.Terminate(ctx); err != nil {
			log.Fatalf("Erro ao finalizar container do mongodb: %v", err)
		}
	}()

	// obtendo o host e port
	host, err := mongoContainer.Host(ctx)
	if err != nil {
		log.Fatalf("Erro ao obter host do mongodb: %v", err)
	}

	port, err := mongoContainer.MappedPort(ctx, "27017")
	if err != nil {
		log.Fatalf("Erro ao obter porta do mongodb: %v", err)
	}

	mongoClient, err = mongointernallib.Connect(&config.AppConfig{
		DBHost: host,
		DBPort: port.Port(),
		DBUser: cfg.DBUser,
		DBPass: cfg.DBPass,
	})
	if err != nil {
		log.Fatalf("Erro ao conectar no mongodb: %v", err)
	}

	defer func() {
		fmt.Println("Database: Desconectando do MongoDB...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Erro ao desconectar do banco de dados: %v", err)
		}
	}()

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestRepositoryCreateUser(t *testing.T) {
	userCollection := mongoClient.Database(cfg.DatabaseName).Collection(cfg.UserCollectionName)

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
	log.Printf("Test Usuário recuperado: %+v", retrievedUser)
}
