package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

func main() {
	payload := []byte("Hello, World!")

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	jweTokenBytes, err := jwe.Encrypt(payload,
		jwe.WithKey(jwa.RSA_OAEP_256, publicKey),
		jwe.WithContentEncryption(jwa.A256GCM),
	)
	if err != nil {
		panic(err)
	}

	jweTokenString := string(jweTokenBytes)
	fmt.Println("JWE Token:\n", jweTokenString)

	decryptedPayload, err := jwe.Decrypt([]byte(jweTokenString),
		jwe.WithKey(jwa.RSA_OAEP_256, privateKey),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted Payload:\n", string(decryptedPayload))
}
