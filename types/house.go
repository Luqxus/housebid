package types

type House struct {
	HouseID string `db:"house_id"`
	Address string `db:"address"`
	ListingPrice float64 `db:"listing_price"`
	Description string `db:"description"`
	Status string `db:"status"` // | AVAILABLE | PENDING | SOLD | EXPIRED
}