package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	AccessTokenExpiry  = 10 * time.Minute
	RefreshTokenExpiry = 24 * time.Hour
)

// CreateJwtToken creates a JWT token with the given username
func CreateJwtToken(username string, dur time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(dur).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Printf("(Error): error creating JWT token for user (%s): %v", username, err)
		return "3", err
	}

	return tokenString, nil
}
