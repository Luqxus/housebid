package types

import "time"

type User struct {
	UID         string    `json:"-" db:"uid"`
	Username    string    `json:"username" db:"username"`
	Email       string    `json:"email" db:"email"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Password    string    `json:"-" db:"password"`
	Address     string    `json:"address" db:"address"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CreateUserData struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
	Address     string `json:"address"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
