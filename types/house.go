package types

import "time"

type House struct {
	HouseID      string    `db:"house_id"`
	UID          string    `db:"uid"`
	Address      string    `db:"address"`
	ListingPrice float64   `db:"listing_price"`
	Description  string    `db:"description"`
	Status       string    `db:"status"` // | AVAILABLE | PENDING | SOLD | EXPIRED
	CreatedAt    time.Time `db:"created_at"`
}

type HouseAddress struct {
	AddressID      string `json:"dddress_id" db:"address_id"`
	HouseID        string `json:"house_id" db:"house_id"`
	BuildingNumber string `json:"building_number" db:"building_no"`
	Surbub         string `json:"surbub" db:"surbub"`
	City           string `json:"city" db:"city"`
	Street         string `json:"street" db:"street"`
	Province       string `json:"province" db:"province"`
	PostalCode     string `json:"postal_code" db:"postal_code"`
}

type HouseImage struct {
	ImageID   string `db:"images_id"`
	ImagesURL string `db:"images_url"`
	HouseID   string `db:"house_id"`
}

type HouseResponse struct {
	HouseID      string       `json:"house_id"`
	UID          string       `json:"uid"`
	Address      HouseAddress `json:"address"`
	ListingPrice float64      `json:"listing_price"`
	Description  string       `json:"description"`
	Images       []string     `json:"images"`
	Status       string       `json:"status"` // | AVAILABLE | PENDING | SOLD | EXPIRED

}

type CreateHouseData struct {
	ListingPrice   float64 `json:"listing_price"`
	Description    string  `json:"description"`
	Status         string  `json:"status"` // | AVAILABLE | PENDING | SOLD | EXPIRED
	BuildingNumber string  `json:"building_no"`
	Surbub         string  `json:"surbub"`
	City           string  `json:"city"`
	Street         string  `json:"street"`
	Province       string  `json:"province"`
	PostalCode     string  `json:"postal_code"`
}

type HouseQueryFilter struct {
	ListingPrice float64 `json:"listing_price"`
	City         string  `json:"city"`
	Surbub       string  `json:"suburb"`
}
