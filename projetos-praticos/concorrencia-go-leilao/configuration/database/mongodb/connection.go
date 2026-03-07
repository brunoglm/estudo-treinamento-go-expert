package mongodb

import (
	"auction-go/configuration/logger"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL = "MONGODB_URL"
	MONGODB_DB  = "MONGODB_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURL := os.Getenv(MONGODB_URL)
	mongoDB := os.Getenv(MONGODB_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		logger.Error("Failed to connect to MongoDB", err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Failed to ping MongoDB", err)
		return nil, err
	}

	logger.Info("Successfully connected to MongoDB")

	return client.Database(mongoDB), nil
}
