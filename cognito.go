package lambda_common

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleCognitoError(err error) ApiError {
	e := err.Error()

	if strings.HasPrefix(e, "ResourceNotFound") {
		return CreateApiError(ErrorCognitoResourceNotFound, http.StatusNotFound, "User registration resource not found", err)
	} else if strings.HasPrefix(e, "InvalidParameter") {
		return CreateApiError(ErrorCognitoInvalidParameter, http.StatusBadRequest, "Invalid parameters", err)
	} else if strings.HasPrefix(e, "UserLambdaValidation") {
		return CreateApiError(ErrorCognitoUserLambdaValidation, http.StatusBadRequest, "Validation error", err)
	} else if strings.HasPrefix(e, "NotAuthorized") {
		return CreateApiError(ErrorCognitoNotAuthorized, http.StatusUnauthorized, "Not authorized", err)
	} else if strings.HasPrefix(e, "InvalidPassword") {
		return CreateApiError(ErrorCognitoInvalidPassword, http.StatusBadRequest, "Invalid password", err)
	} else if strings.HasPrefix(e, "UsernameExists") {
		return CreateApiError(ErrorCognitoUsernameExists, http.StatusConflict, "User already exists", err)
	} else if strings.HasPrefix(e, "TooManyRequests") {
		return CreateApiError(ErrorCognitoTooManyRequests, http.StatusTooManyRequests, "Too many requests", err)
	} else if strings.HasPrefix(e, "LimitExceeded") {
		return CreateApiError(ErrorCognitoLimitExceeded, http.StatusTooManyRequests, "Limit exceeded", err)
	} else if strings.HasPrefix(e, "Forbidden") {
		return CreateApiError(ErrorCognitoForbidden, http.StatusForbidden, "Forbidden", err)
	} else if strings.HasPrefix(e, "UserNotFound") {
		return CreateApiError(ErrorCognitoUnauthorized, http.StatusUnauthorized, "Unauthorized", err)
	} else {
		fmt.Println(err)
		return CreateApiError(ErrorCognitoUnexpected, http.StatusInternalServerError, "Unexpected error", err)
	}

}
