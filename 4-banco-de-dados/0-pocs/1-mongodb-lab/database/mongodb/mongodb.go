package mongodb

import (
	"context"
	"fmt"
	"mongodb-lab/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(config *config.AppConfig) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", config.DBUser, config.DBPass, config.DBHost, config.DBPort)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar ao MongoDB: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Erro ao pingar o MongoDB: %v", err)
	}

	fmt.Println("Database: Conectado ao MongoDB com sucesso!")

	return client, nil
}
