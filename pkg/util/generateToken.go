package util

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint) (string, error) {
	// Parse JWT expiration time from environment variable
	expirationStr := os.Getenv("JWT_EXPIRATION")
	if expirationStr == "" {
		expirationStr = "1h" // Default to 1 hour if not specified
	}
	
	// Parse the duration
	expDuration, err := time.ParseDuration(expirationStr)
	if err != nil {
		return "", err
	}
	
	// Create token claims
	claims := jwt.MapClaims{
		"user_id": strconv.FormatUint(uint64(userID), 10),
		"exp":     float64(time.Now().Add(expDuration).Unix()),
	}
	
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	
	return tokenString, nil
}