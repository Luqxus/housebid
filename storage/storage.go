package storage

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/luquxSentinel/housebid/types"
)

type Storage interface {
	CreateUser(ctx context.Context, user *types.User) error
	CountEmail(ctx context.Context, email string) (int64, error)
	CountPhoneNumber(ctx context.Context, phonenumber string) (int64, error)
	GetUserByEmail(ctx context.Context, email string) (*types.User, error)
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
