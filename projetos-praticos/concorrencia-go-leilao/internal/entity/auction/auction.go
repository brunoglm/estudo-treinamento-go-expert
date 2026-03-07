package auction

import (
	"context"
	"time"

	"auction-go/internal/errors"
)

type ProductCondition int

const (
	New ProductCondition = iota + 1
	Used
	Refurbished
)

type AuctionStatus int

const (
	Active AuctionStatus = iota
	Completed
)

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auctionEntity *Auction) *errors.Error
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *errors.Error)
	FindAuctionById(ctx context.Context, id string) (*Auction, *errors.Error)
}
