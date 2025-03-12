package jwt

import (
	"log"

	"github.com/golang-jwt/jwt"
)

// VerifyToken verifies the JWT token
func VerifyToken(tokenString string) (error, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Printf("(Error): error parsing JWT token. Unacceptable token (not consistent with our secret key): %v", err)
		return err, false
	}

	if !token.Valid {
		log.Printf("(Error): expired token. Token is no longer valid.")
		return nil, false
	}

	return nil, true
}
