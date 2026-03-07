package bid

import (
	"auction-go/internal/errors"
	"context"
	"time"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

type BidEntityRepository interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *errors.Error
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]Bid, *errors.Error)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *errors.Error)
}
