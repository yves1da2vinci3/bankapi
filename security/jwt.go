package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a secret key for signing and verifying JWT tokens
var secretKey = []byte("your-secret-key")

// GenerateJWTToken generates a new JWT token for a given user ID
func GenerateJWTToken(userID uint) (string, error) {
	// Create a new JWT token with claims
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWTToken parses a JWT token and returns the user ID from the claims
func ParseJWTToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, err
	}

	userID := claims["userID"].(float64) // User ID is stored as a float64 in claims
	return uint(userID), nil
}
