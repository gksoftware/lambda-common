package lambda_common

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

const (
	ErrorFailedToUnmarshal           = 100
	ErrorFailedToMarshal             = 200
	ErrorFailedToHashSecret          = 300
	ErrorCognitoResourceNotFound     = 400
	ErrorCognitoInvalidParameter     = 401
	ErrorCognitoUserLambdaValidation = 402
	ErrorCognitoNotAuthorized        = 403
	ErrorCognitoInvalidPassword      = 404
	ErrorCognitoUsernameExists       = 405
	ErrorCognitoTooManyRequests      = 406
	ErrorCognitoLimitExceeded        = 407
	ErrorCognitoForbidden            = 408
	ErrorCognitoUnexpected           = 499
)

func (e ApiError) Error() string {
	js, err := json.Marshal(e.Body)
	if err != nil {
		js = []byte("Error marshalling body")
	}

	return fmt.Sprintf("Id: %d, StatusCode: %d, Error: %s, Body: %s", e.Id, e.StatusCode, e.Err, string(js))
}

func CreateApiError(id int, status int, message string, err error) ApiError {
	var e ApiError
	e.Id = id
	e.StatusCode = status
	e.Err = err
	e.Body.Id = id
	e.Body.StatusCode = status
	e.Body.Message = message
	return e
}

func CreateResponseError(e error) (events.APIGatewayProxyResponse, error) {
	var apiError ApiError
	if errors.As(e, &apiError) {
		body, err := json.Marshal(apiError.Body)
		if err != nil {
			body = []byte("Error marshalling body")
		}

		return events.APIGatewayProxyResponse{
			StatusCode: apiError.StatusCode,
			Body:       string(body),
		}, nil
	} else {
		var b ApiErrorBody
		b.StatusCode = http.StatusInternalServerError
		b.Message = "Unknown server error"

		body, err := json.Marshal(b)
		if err != nil {
			body = []byte("Error marshalling body")
		}
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       string(body),
		}, nil
	}
}
