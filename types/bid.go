package types

import "time"

type Bid struct {
	BidID string `db:"bid_id"`
	HouseID string `db:"house_id"`
	UID string `db:"uid"`
	BidAmount float64 `db:"bid_amount"`
	BidTime time.Time `db:"bid_time"`
}

type StandingBid struct {
	StandingBidID string `db:"standing_bid_id"`
	HouseID string `db:"house_id"`
	BidAmount float64 `db:"bid_amount"`
}