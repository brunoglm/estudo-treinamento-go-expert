package auctionrepository

import (
	"auction-go/configuration/logger"
	"auction-go/internal/entity/auction"
	"auction-go/internal/errors"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction.Auction, *errors.Error) {
	filter := bson.M{"_id": id}

	var auctionEntityMongo AuctionEntityMongo
	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find auction by id = %s", id), err)

		return nil, errors.NewInternalServerError("Error trying to find auction by id")
	}

	return &auction.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
}

func (repo *AuctionRepository) FindAuctions(ctx context.Context, status auction.AuctionStatus, category string, productName string) ([]auction.Auction, *errors.Error) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["productName"] = primitive.Regex{Pattern: productName, Options: "i"}
	}

	cursor, err := repo.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error finding auctions", err)
		return nil, errors.NewInternalServerError("Error finding auctions")
	}
	defer cursor.Close(ctx)

	var auctionsMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionsMongo); err != nil {
		logger.Error("Error decoding auctions", err)
		return nil, errors.NewInternalServerError("Error decoding auctions")
	}

	auctionsEntity := make([]auction.Auction, 0, len(auctionsMongo))
	for _, auctionMongo := range auctionsMongo {
		auctionsEntity = append(auctionsEntity, auction.Auction{
			Id:          auctionMongo.Id,
			ProductName: auctionMongo.ProductName,
			Category:    auctionMongo.Category,
			Status:      auctionMongo.Status,
			Description: auctionMongo.Description,
			Condition:   auctionMongo.Condition,
			Timestamp:   time.Unix(auctionMongo.Timestamp, 0),
		})
	}

	return auctionsEntity, nil
}
