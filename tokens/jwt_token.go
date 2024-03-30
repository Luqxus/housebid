package tokens

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string
	UID      string
	jwt.RegisteredClaims
}

func GenerateJWT(username, uid string) (signedToken string, err error) {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		err = errors.New("internal server error occured while logging user in. please try again")
		return
	}

	claims := &Claims{
		UID:      uid,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(30 * time.Hour)),
		},
	}

	// create new & signed token
	signedToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
	return
}

func ValidateJWT(signedToken string) (uid string, err error) {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		err = errors.New("internal server error occured while logging user in. please try again")
		return
	}
	var claims Claims
	token, err := jwt.ParseWithClaims(signedToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if !token.Valid {
		err = errors.New("invalid authorization header")
		return
	}

	uid = claims.UID
	return
}
