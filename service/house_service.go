package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/luquxSentinel/housebid/storage"
	"github.com/luquxSentinel/housebid/types"
)

type HouseService interface {
	// Create a house | cannot list a house with the same address twice

	// Create HouseImages

	// Create House Address | a house cannot 2 address | cannot have duplicate address

	// GetHouseByLocation

	// GetHouseByID
}

type houseService struct {
	storage storage.Storage
}

func NewHouseService(storage storage.Storage) *houseService {
	return &houseService{
		storage: storage,
	}
}

func (s *houseService) ListHouse(ctx context.Context, data *types.CreateHouseData, uid string) error {
	// TODO check if no house with same city & province & building number
	count, err := s.storage.CountAddress(ctx, data.City, data.Province, data.BuildingNumber)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("house already listed")
	}

	// TODO: create a house
	newhouse := new(types.House)
	newhouse.UID = uid
	newhouse.Description = data.Description
	newhouse.HouseID = uuid.NewString()
	newhouse.ListingPrice = data.ListingPrice
	newhouse.Status = "AVAILABLE"
	newhouse.CreatedAt = time.Now().Local()

	// TODO: create house address
	newaddress := new(types.HouseAddress)
	newaddress.AddressID = uuid.NewString()
	newaddress.BuildingNumber = data.BuildingNumber
	newaddress.City = data.City
	newaddress.HouseID = newhouse.HouseID
	newaddress.PostalCode = data.PostalCode
	newaddress.Surbub = data.Surbub
	newaddress.Street = data.Street
	newaddress.Province = data.Province

	//TODO: create house images

	// TODO: persist house & house address in storage

	// TODO: error handling
	return nil
}
