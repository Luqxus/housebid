package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/luquxSentinel/housebid/storage"
	"github.com/luquxSentinel/housebid/tokens"
	"github.com/luquxSentinel/housebid/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(ctx context.Context, data *types.CreateUserData) error
	LoginUser(ctx context.Context, data *types.LoginData) (*types.User, string, error)
}

type authService struct {
	storage storage.Storage
}

func NewAuthService(s storage.Storage) *authService {
	return &authService{
		storage: s,
	}
}

func (s *authService) RegisterUser(ctx context.Context, data *types.CreateUserData) error {
	//TODO: check if email is not in use
	count, err := s.storage.CountEmail(ctx, data.Email)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("email already in use")
	}

	// TODO: check if phone number is not in use
	count, err = s.storage.CountPhoneNumber(ctx, data.PhoneNumber)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("phone number already in use")
	}

	//TODO: create new user from data
	user := new(types.User)

	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.PhoneNumber = data.PhoneNumber
	user.Address = data.Address
	user.CreatedAt = time.Now().Local()
	user.Username = data.Username

	// hash password
	pwd, err := hashPassword(data.Password)
	if err != nil {
		return err
	}
	user.Password = pwd

	// TODO: persist user into storage
	return s.storage.CreateUser(ctx, user)
}

func (s *authService) LoginUser(ctx context.Context, data *types.LoginData) (*types.User, string, error) {
	// get user by email from storage
	user, err := s.storage.GetUserByEmail(ctx, data.Email)
	if err != nil {
		log.Println(err)
		return nil, "", errors.New("wrong email or password")
	}

	err = verifyPassword(user.Password, data.Password)
	if err != nil {
		log.Println(err)
		return nil, "", errors.New("wrong email or password")
	}

	// TODO: generate jwt token
	token, err := tokens.GenerateJWT(user.Username, user.UID)
	if err != nil {
		log.Println(err)
		return nil, "", errors.New("internal error occured while signing user in. please try again")
	}

	return user, token, nil

}

func hashPassword(pwd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)

	return string(b), err
}

func verifyPassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
