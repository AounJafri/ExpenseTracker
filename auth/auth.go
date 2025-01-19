package auth

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(30 * time.Minute).Unix(), // Expiration time (1 day)
		"iat": time.Now().Unix(),                       // Issue time
		"iss": "expense-tracker-app",
	})

	tokenString, err := token.SignedString(secret)

	return tokenString, err
}

func VerifyToken(tokenString string) (bool, float64) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return false, 0
	}

	if !token.Valid {
		return false, 0
	}

	return true, token.Claims.(jwt.MapClaims)["id"].(float64)
}
