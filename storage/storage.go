package storage

import (
	"context"

	"github.com/luquxSentinel/housebid/types"
)

type Storage interface {
	CreateUser(ctx context.Context, user *types.User) error
	CountEmail(ctx context.Context, email string) (int64, error)
	CountPhoneNumber(ctx context.Context, phonenumber string) (int64, error)
}

type storage struct{}

func New() *storage {
	return &storage{}
}
