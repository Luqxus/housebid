package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/luquxSentinel/housebid/types"
)

type Storage interface {
	CreateUser(ctx context.Context, user *types.User) error
	CountEmail(ctx context.Context, email string) (int64, error)
	CountPhoneNumber(ctx context.Context, phonenumber string) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
	CountAddress(ctx context.Context, city, province, building_number string) (int64, error)
	CreateNewHouse(ctx context.Context, house *types.House, address *types.HouseAddress) error
	GetHouseByFilter(ctx context.Context, listingPrice float64, city string, surbub string) ([]*types.HouseResponse, error)
}

type postgresStorage struct {
	db *sqlx.DB
}

func New() (*postgresStorage, error) {
	db, err := dbconnection()
	if err != nil {
		return nil, err
	}
	return &postgresStorage{
		db: db,
	}, nil
}

func (s *postgresStorage) CreateUser(ctx context.Context, user *types.User) error {
	query := `INSERT INTO Users (
		uid, first_name, last_name, username, email, phone_number, password, created_at, address
	)

	VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, err := s.db.ExecContext(
		ctx,
		query,
		user.UID,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		user.PhoneNumber,
		user.Password, user.CreatedAt,
		user.Address,
	)

	return err
}

func (s *postgresStorage) CountEmail(ctx context.Context, email string) (int64, error) {
	var count int64

	query := `SELECT COUNT(email) FROM Users WHERE email = $1`

	row := s.db.QueryRowContext(ctx, query, email)

	err := row.Scan(&count)
	return count, err
}

func (s *postgresStorage) CountPhoneNumber(ctx context.Context, phonenumber string) (int64, error) {
	var count int64

	query := `SELECT COUNT(phone_number) FROM Users WHERE phone_number = $1`

	row := s.db.QueryRowContext(ctx, query, phonenumber)

	err := row.Scan(&count)
	return count, err
}

func (s *postgresStorage) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	query := `SELECT * FROM Users WHERE email = $1`

	row := s.db.QueryRowxContext(ctx, query, email)

	user := new(types.User)

	err := row.StructScan(user)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, err
}

func (s *postgresStorage) CountAddress(ctx context.Context, city, province, building_number string) (int64, error) {
	query := `
	SELECT COUNT(building_number)
	WHERE city = $1
	AND province = $2
	AND building_number = $3`

	row := s.db.QueryRowContext(ctx, query, city, province, building_number)

	var count int64
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, err
}

func (s *postgresStorage) CreateNewHouse(ctx context.Context, house *types.House, address *types.HouseAddress) error {

	// transactions
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `INSERT INTO Houses (house_id, uid, listing_price, description, status)
	VALUES ($1, $2, $3, $4, $5)`

	_, err = tx.ExecContext(ctx, query, house.HouseID, house.UID, house.ListingPrice, house.Description, house.Status)
	if err != nil {
		return err
	}

	query = `INSERT INTO HouseAddresses (address_id, house_id,building_no, city, street, surbub, province, postal_code)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = tx.ExecContext(
		ctx,
		query,
		address.AddressID,
		address.HouseID,
		address.BuildingNumber,
		address.City,
		address.Street,
		address.Surbub,
		address.Province,
		address.PostalCode,
	)

	if err != nil {
		return err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *postgresStorage) GetHouseByFilter(ctx context.Context, listingPrice float64, city string, surbub string) ([]*types.HouseResponse, error) {
	query := "SELECT * FROM Houses  INNER JOIN HouseAddresses ON Houses.house_id = HouseAddresses.house_id "

	if city != "" {
		query = fmt.Sprintf("%s WHERE city = %s", query, city)
	}

	if surbub != "" {
		var str string
		if city != "" {
			str = fmt.Sprintf("AND surbub = %s", surbub)
		} else {
			str = fmt.Sprintf("WHERE surbub = %s", surbub)
		}

		query = fmt.Sprintf("%s %s", query, str)

	}

	if listingPrice != 0 {
		if city != "" || surbub != "" {
			query = fmt.Sprintf("%s AND listing_price = %f", query, listingPrice)
		} else {
			query = fmt.Sprintf("%s WHERE listing_price = %f", query, listingPrice)
		}
	}

	rows, err := s.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}

	houses := make([]*types.HouseResponse, 0)

	for rows.Next() {
		house := types.HouseResponse{}
		err = rows.StructScan(house)
		if err != nil {
			return nil, err
		}

		houses = append(houses, &house)

	}

	return houses, nil
}
