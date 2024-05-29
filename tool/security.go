package tool

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const DEFAULT_COST = 13

type JwtCustomClaim struct {
	Id int64 `json:"id"`
	jwt.RegisteredClaims
}

func GenerateHash(password []byte) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword(password, DEFAULT_COST)
	if err != nil {
		return
	}

	return
}

func CompareHashAndPassword(hash []byte, password []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return
	}

	return
}

func GenerateJWT(id int64) (signed string, err error) {
	claims := JwtCustomClaim{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "GoBlog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SIGNING_KEY")))
}

func VerifyJWT(tokenString string) (claims jwt.MapClaims, err error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return
	}

	if !token.Valid {
		err = errors.New("invalid token")
		return
	}

	claims = token.Claims.(jwt.MapClaims)
	return
}
