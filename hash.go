package lambda_common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
)

func GenerateSecretHash(username, clientId, clientSecret string) (string, error) {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	_, err := mac.Write([]byte(username + clientId))
	if err != nil {
		return "", CreateApiError(ErrorFailedToHashSecret, http.StatusInternalServerError, "Internal server error", err)
	}

	secretHash := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return secretHash, nil
}
