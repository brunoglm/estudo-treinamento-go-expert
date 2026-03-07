package auction

import (
	"auction-go/internal/internalerror"
	"context"
	"time"
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
	CreateAuction(
		ctx context.Context,
		auctionEntity *Auction) *internalerror.InternalError

	FindAuctions(
		ctx context.Context,
		status AuctionStatus,
		category, productName string) ([]Auction, *internalerror.InternalError)

	FindAuctionById(
		ctx context.Context, id string) (*Auction, *internalerror.InternalError)
}
