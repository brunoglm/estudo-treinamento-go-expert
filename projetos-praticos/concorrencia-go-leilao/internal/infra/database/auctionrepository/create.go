package auctionrepository

import (
	"auction-go/configuration/logger"
	"auction-go/internal/entity/auction"
	"auction-go/internal/errors"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                   `bson:"_id"`
	ProductName string                   `bson:"product_name"`
	Category    string                   `bson:"category"`
	Description string                   `bson:"description"`
	Condition   auction.ProductCondition `bson:"condition"`
	Status      auction.AuctionStatus    `bson:"status"`
	Timestamp   int64                    `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auctionEntity *auction.Auction) *errors.Error {
	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		message := "Error trying to insert auction"
		logger.Error(message, err)
		return errors.NewInternalServerError(message)
	}

	return nil
}
