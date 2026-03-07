package bidrepository

import (
	"auction-go/configuration/logger"
	"auction-go/internal/entity/bid"
	"auction-go/internal/errors"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]bid.Bid, *errors.Error) {
	filter := bson.M{"auctionId": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		message := fmt.Sprintf("Find: Error trying to find bids by auctionId %s", auctionId)

		logger.Error(message, err)

		return nil, errors.NewInternalServerError(message)
	}

	var bidEntitiesMongo []BidEntityMongo
	err = cursor.All(ctx, &bidEntitiesMongo)
	if err != nil {
		message := fmt.Sprintf("Cursor: Error trying to find bids by auctionId %s", auctionId)

		logger.Error(message, err)

		return nil, errors.NewInternalServerError(message)
	}
	defer cursor.Close(ctx)

	bidEntities := make([]bid.Bid, 0, len(bidEntitiesMongo))
	for _, bidEntityMongo := range bidEntitiesMongo {
		bidEntities = append(bidEntities, bid.Bid{
			Id:        bidEntityMongo.Id,
			UserId:    bidEntityMongo.UserId,
			AuctionId: bidEntityMongo.AuctionId,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bid.Bid, *errors.Error) {
	filter := bson.M{"auction_id": auctionId}

	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	var bidEntityMongo BidEntityMongo
	err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo)
	if err != nil {
		logger.Error("Error trying to find the auction winner", err)
		return nil, errors.NewInternalServerError("Error trying to find the auction winner")
	}

	return &bid.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
