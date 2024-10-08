package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"regexp"

	"fivetrace.com/iam_service/internal/application/managers"
	"github.com/golang-jwt/jwt/v5"
)

func ComputeSecretHash(clientId, clientSecret, username string) string {
	key := []byte(clientSecret)
	message := []byte(username + clientId)
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func GenerateJWT(claims managers.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(jwtSecret)
}
