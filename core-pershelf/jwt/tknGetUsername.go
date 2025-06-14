package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// GetUsernameFromToken gets the username from the token
func GetUsernameFromToken(tokenString string) (string, error) {
	if tokenString == "" {
		return "", errors.New("no token string")
	}

	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// Return if there's an error or if the token is invalid
	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	// Extract the username from the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if username, ok := claims["username"].(string); ok {
			return username, nil
		}
		return "", errors.New("username not found in token claims")
	}

	return "", errors.New("invalid token claims")
}
