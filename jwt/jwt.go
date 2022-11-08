package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"expire": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET_KEY"))

	if err != nil {
		log.Fatalf("TOKEN ERROR: %v", err)
	}

	return tokenString
}
