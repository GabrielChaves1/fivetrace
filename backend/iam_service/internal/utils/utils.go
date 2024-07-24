package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeSecretHash(clientId, clientSecret, username string) string {
	key := []byte(clientSecret)
	message := []byte(username + clientId)
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}