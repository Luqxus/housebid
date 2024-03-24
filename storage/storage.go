package storage

import (
	"context"
	"database/sql"

	"github.com/luquxSentinel/housebid/types"
)

type Storage interface {
	CreateUser(ctx context.Context, user *types.User) error
	CountEmail(ctx context.Context, email string) (int64, error)
	CountPhoneNumber(ctx context.Context, phonenumber string) (int64, error)
}

type postgresStorage struct {
	db *sql.DB
}

func New() *postgresStorage {
	db := dbconnection()
	return &postgresStorage{
		db: db,
	}
}

func (s *postgresStorage) CreateUser(ctx context.Context, user *types.User) error {
	return nil
}

func (s *postgresStorage) CountEmail(ctx context.Context, email string) (int64, error) {
	var count int64

	query := `SELECT COUNT(email) FROM Users WHERE email = $1`

	row, err := s.db.QueryContext(ctx, query, email)
	if err != nil {
		return -1, err
	}
	err = row.Scan(&count)
	return count, err
}
func (s *postgresStorage) CountPhoneNumber(ctx context.Context, phonenumber string) (int64, error) {
	var count int64

	query := `SELECT COUNT(phone_number) FROM Users WHERE phone_number = $1`

	row, err := s.db.QueryContext(ctx, query, phonenumber)
	if err != nil {
		return -1, err
	}
	err = row.Scan(&count)
	return count, err
}
