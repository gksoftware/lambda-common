package lambda_common

import (
	"fmt"
	"net/http"
	"strings"
)

func HandleCognitoError(err error) ApiError {
	e := err.Error()

	if strings.HasPrefix(e, "ResourceNotFoundException") {
		return CreateApiError(ErrorCognitoResourceNotFoundException, http.StatusNotFound, "User registration resource not found", err)
	} else if strings.HasPrefix(e, "InvalidParameterException") {
		return CreateApiError(ErrorCognitoInvalidParameterException, http.StatusBadRequest, "Invalid parameters", err)
	} else if strings.HasPrefix(e, "UserLambdaValidationException") {
		return CreateApiError(ErrorCognitoUserLambdaValidationException, http.StatusBadRequest, "Validation error", err)
	} else if strings.HasPrefix(e, "NotAuthorizedException") {
		return CreateApiError(ErrorCognitoNotAuthorizedException, http.StatusForbidden, "Not authorized", err)
	} else if strings.HasPrefix(e, "InvalidPasswordException") {
		return CreateApiError(ErrorCognitoInvalidPasswordException, http.StatusBadRequest, "Invalid password", err)
	} else if strings.HasPrefix(e, "UsernameExistsException") {
		return CreateApiError(ErrorCognitoUsernameExistsException, http.StatusConflict, "User already exists", err)
	} else if strings.HasPrefix(e, "TooManyRequestsException") {
		return CreateApiError(ErrorCognitoTooManyRequestsException, http.StatusTooManyRequests, "Too many requests", err)
	} else if strings.HasPrefix(e, "LimitExceededException") {
		return CreateApiError(ErrorCognitoLimitExceededException, http.StatusTooManyRequests, "Limit exceeded", err)
	} else if strings.HasPrefix(e, "ForbiddenException") {
		return CreateApiError(ErrorCognitoForbiddenException, http.StatusForbidden, "Forbidden", err)
	} else {
		fmt.Println(err)
		return CreateApiError(ErrorCognitoUnexpectedException, http.StatusInternalServerError, "Unexpected error", err)
	}
}
