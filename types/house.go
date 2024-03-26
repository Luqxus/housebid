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
	AddressID      string `db:"address_id"`
	HouseID        string `db:"house_id"`
	BuildingNumber string `db:"building_no"`
	Surbub         string `db:"surbub"`
	City           string `db:"city"`
	Street         string `db:"street"`
	Province       string `db:"province"`
	PostalCode     string `db:"postal_code"`
}

type HouseImage struct {
	ImageID   string `db:"images_id"`
	ImagesURL string `db:"images_url"`
	HouseID   string `db:"house_id"`
}

type HouseResponse struct {
	HouseID      string   `json:"house_id"`
	UID          string   `json:"uid"`
	Address      string   `json:"address"`
	ListingPrice float64  `json:"listing_price"`
	Description  string   `json:"description"`
	Images       []string `json:"images"`
	Status       string   `json:"status"` // | AVAILABLE | PENDING | SOLD | EXPIRED

}

type CreateHouseData struct {
	ListingPrice   float64 `json:"listing_price"`
	Description    string  `json:"description"`
	Status         string  `json:"status"` // | AVAILABLE | PENDING | SOLD | EXPIRED
	BuildingNumber string  `db:"building_no"`
	Surbub         string  `db:"surbub"`
	City           string  `db:"city"`
	Street         string  `db:"street"`
	Province       string  `db:"province"`
	PostalCode     string  `db:"postal_code"`
}
