package bidrepository

import (
	"auction-go/configuration/logger"
	"auction-go/internal/entity/auction"
	"auction-go/internal/entity/bid"
	"auction-go/internal/errors"
	"auction-go/internal/infra/database/auctionrepository"
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auctionrepository.AuctionRepository
}

func NewBidRepository(database *mongo.Database, auctionRepository *auctionrepository.AuctionRepository) *BidRepository {
	return &BidRepository{
		Collection:        database.Collection("bids"),
		AuctionRepository: auctionRepository,
	}
}

func (bd *BidRepository) CreateBid(ctx context.Context, bidEntities []bid.Bid) *errors.Error {
	var wg sync.WaitGroup

	bidsGroupedByAuctionId := make(map[string][]any)

	for _, bid := range bidEntities {
		bidEntityMongo := BidEntityMongo{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp.Unix(),
		}

		bidsGroupedByAuctionId[bid.AuctionId] = append(bidsGroupedByAuctionId[bid.AuctionId], bidEntityMongo)
	}

	for auctionId, bidsGrouped := range bidsGroupedByAuctionId {
		wg.Go(func() {
			auctionEntity, err := bd.AuctionRepository.FindAuctionById(ctx, auctionId)
			if err != nil {
				logger.Error("Error trying to find auction by id", err)
				return
			}

			if auctionEntity.Status == auction.Completed {
				return
			}

			if _, err := bd.Collection.InsertMany(ctx, bidsGrouped); err != nil {
				logger.Error("Error trying to insert bids", err)
				return
			}
		})
	}

	wg.Wait()

	return nil
}
