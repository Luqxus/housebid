package types

import "time"

type User struct {
	UID         string    `db:"uid"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	PhoneNumber string    `db:"phone_number"`
	Password    string    `db:"password"`
	Address     string    `db:"address"`
	CreatedAt   time.Time `db:"created_at"`
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
