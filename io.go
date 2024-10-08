package lambda_common

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"strings"
)

func UnmarshalRequest(req events.APIGatewayProxyRequest, payload any) error {
	if err := json.Unmarshal([]byte(req.Body), &payload); err != nil {
		return CreateApiError(ErrorFailedToUnmarshal, http.StatusBadRequest, "Invalid payload", err)
	}
	return nil
}

func MarshalResponse(payload any) (string, error) {
	js, err := json.Marshal(payload)
	if err != nil {
		return "", CreateApiError(ErrorFailedToMarshal, http.StatusInternalServerError, "Invalid response payload", err)
	}
	return string(js), nil
}

func ExtractCookies(cookieString string) *map[string]string {
	var m = make(map[string]string)

	cookies := strings.Split(cookieString, ";")
	for _, c := range cookies {
		if len(c) == 0 {
			continue
		}
		parts := strings.Split(strings.Trim(c, " "), "=")
		m[parts[0]] = parts[1]
	}

	return &m
}
