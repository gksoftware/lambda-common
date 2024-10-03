package lambda_common

import (
	"fmt"
	"net/http"
	"testing"
)

func Test(t *testing.T) {

}

func test_error_id(want, got int, t *testing.T) {
	if want != got {
		t.Fatalf("Invalid error id. Wanted: %d, Got: %d", want, got)
	}
}

func test_error_status_code(want, got int, t *testing.T) {
	if want != got {
		t.Fatalf("Invalid status code. Wanted: %d, Got: %d", want, got)
	}
}

func Test_HandleCognitoError_WithResourceNotFoundException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("ResourceNotFoundException"))
	test_error_id(ErrorCognitoResourceNotFoundException, e.Id, t)
	test_error_status_code(http.StatusNotFound, e.StatusCode, t)
}

func Test_HandleCognitoError_WithInvalidParameterException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("InvalidParameterException"))
	test_error_id(ErrorCognitoInvalidParameterException, e.Id, t)
	test_error_status_code(http.StatusBadRequest, e.StatusCode, t)
}

func Test_HandleCognitoError_WithUserLambdaValidationException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("UserLambdaValidationException"))
	test_error_id(ErrorCognitoUserLambdaValidationException, e.Id, t)
	test_error_status_code(http.StatusBadRequest, e.StatusCode, t)
}

func Test_HandleCognitoError_WithNotAuthorizedException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("NotAuthorizedException"))
	test_error_id(ErrorCognitoNotAuthorizedException, e.Id, t)
	test_error_status_code(http.StatusForbidden, e.StatusCode, t)
}

func Test_HandleCognitoError_WithInvalidPasswordException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("InvalidPasswordException"))
	test_error_id(ErrorCognitoInvalidPasswordException, e.Id, t)
	test_error_status_code(http.StatusBadRequest, e.StatusCode, t)
}

func Test_HandleCognitoError_WithUsernameExistsException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("UsernameExistsException"))
	test_error_id(ErrorCognitoUsernameExistsException, e.Id, t)
	test_error_status_code(http.StatusConflict, e.StatusCode, t)
}

func Test_HandleCognitoError_WithTooManyRequestsException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("TooManyRequestsException"))
	test_error_id(ErrorCognitoTooManyRequestsException, e.Id, t)
	test_error_status_code(http.StatusTooManyRequests, e.StatusCode, t)
}

func Test_HandleCognitoError_WithLimitExceededException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("LimitExceededException"))
	test_error_id(ErrorCognitoLimitExceededException, e.Id, t)
	test_error_status_code(http.StatusTooManyRequests, e.StatusCode, t)
}

func Test_HandleCognitoError_WithForbiddenException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("ForbiddenException"))
	test_error_id(ErrorCognitoForbiddenException, e.Id, t)
	test_error_status_code(http.StatusForbidden, e.StatusCode, t)
}

func Test_HandleCognitoError_WithUnexpectedException(t *testing.T) {
	e := HandleCognitoError(fmt.Errorf("UnexpectedException"))
	test_error_id(ErrorCognitoUnexpectedException, e.Id, t)
	test_error_status_code(http.StatusInternalServerError, e.StatusCode, t)
}
