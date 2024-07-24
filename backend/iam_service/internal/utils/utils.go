package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ComputeSecretHash(clientId, clientSecret, username string) string {
	key := []byte(clientSecret)
	message := []byte(username + clientId)
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(jwtSecret)
}
