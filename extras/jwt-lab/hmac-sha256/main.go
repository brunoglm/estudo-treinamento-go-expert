package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var hmacSecret = []byte("my_secret_key")

type MyCustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func generateTokenHMAC(userID uint, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := MyCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "meu-app",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func validateTokenHMAC(tokenString string) (*MyCustomClaims, error) {
	claims := &MyCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func justOpenPayloadTokenHMAC(tokenString string) (*MyCustomClaims, error) {
	claims := &MyCustomClaims{}

	_, _, err := jwt.NewParser().ParseUnverified(tokenString, claims)

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	return claims, nil
}

func main() {
	token, err := generateTokenHMAC(12345, "admin")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	// fmt.Println("Generated Token:", token)

	claims, err := validateTokenHMAC(token)
	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}
	fmt.Printf("Validated Claims: %+v\n", claims)

	claims2, err := justOpenPayloadTokenHMAC(token)
	if err != nil {
		fmt.Println("Error just opening payload token:", err)
		return
	}
	fmt.Printf("Just Opened Payload Claims: %+v\n", claims2)
}
