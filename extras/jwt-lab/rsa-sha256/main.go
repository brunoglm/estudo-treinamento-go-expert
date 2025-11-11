package main

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Comando Gerar Chaves RSA
// Gerar a chave privada:
// openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
//
// Gerar a chave pública a partir da chave privada:
// openssl rsa -pubout -in private_key.pem -out public_key.pem

type MyCustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func loadPrivateKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return privateKey, nil
}

func loadPublicKey(path string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	return publicKey, nil
}

func generateTokenRSA(userID uint, role string, privateKey *rsa.PrivateKey) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := MyCustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "meu-app-rsa",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func validateTokenRSA(tokenString string, publicKey *rsa.PublicKey) (*MyCustomClaims, error) {
	claims := &MyCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func main() {
	privateKey, _ := loadPrivateKey("./private_key.pem")
	token, err := generateTokenRSA(12345, "admin", privateKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}
	fmt.Println("Generated Token:\n", token)

	publicKey, _ := loadPublicKey("./public_key.pem")
	claims, err := validateTokenRSA(token, publicKey)
	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}
	fmt.Printf("Validated Claims: %+v\n", claims)
}
