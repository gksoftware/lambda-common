package lambda_common

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"testing"
)

func Test_Error(t *testing.T) {
	var err ApiError
	err.Id = 100
	err.StatusCode = 404
	err.Body.Id = 100
	err.Body.StatusCode = 404
	err.Body.Message = "testing"
	err.Err = fmt.Errorf("Message: %s", "testing")

	want := "Id: 100, StatusCode: 404, Error: Message: testing, Body: {\"Id\":100,\"StatusCode\":404,\"Message\":\"testing\"}"
	got := err.Error()

	if want != got {
		t.Fatalf("Wanted: %s, Got: %s", want, got)
	}
}

func Test_Error_Without_Body(t *testing.T) {
	var err ApiError
	err.Id = 100
	err.StatusCode = 404
	err.Err = fmt.Errorf("Message: %s", "testing")

	want := "Id: 100, StatusCode: 404, Error: Message: testing, Body: {\"Id\":0,\"StatusCode\":0,\"Message\":\"\"}"
	got := err.Error()

	if want != got {
		t.Fatalf("Wanted: %s, Got: %s", want, got)
	}
}

func Test_CreateApiError(t *testing.T) {
	var want ApiError
	want.Id = 200
	want.StatusCode = 500
	want.Err = fmt.Errorf("Message: %s", "test")
	want.Body.Id = 200
	want.Body.StatusCode = 500
	want.Body.Message = "Test message"

	got := CreateApiError(200, 500, "Test message", fmt.Errorf("Message: %s", "test"))

	if want.Id != got.Id {
		t.Fatalf("Wanted %s=%d, but got %d", "Id", want.Id, got.Id)
	}
	if want.StatusCode != got.StatusCode {
		t.Fatalf("Wanted %s=%d, but got %d", "StatusCode", want.StatusCode, got.StatusCode)
	}
	if want.Err.Error() != got.Err.Error() {
		t.Fatalf("Wanted %s=%s, but got %s", "Err", want.Err.Error(), got.Err.Error())
	}
	if want.Body.Id != got.Body.Id {
		t.Fatalf("Wanted %s=%d, but got %d", "Body.Id", want.Body.Id, got.Body.Id)
	}
	if want.Body.Message != got.Body.Message {
		t.Fatalf("Wanted %s=%s, but got %s", "Body.Message", want.Body.Message, got.Body.Message)
	}
	if want.Body.StatusCode != got.Body.StatusCode {
		t.Fatalf("Wanted %s=%d, but got %d", "Body.StatusCode", want.Body.StatusCode, got.Body.StatusCode)
	}
}

func Test_ResponseError_WithApiError(t *testing.T) {
	var err ApiError
	err.Id = 200
	err.StatusCode = 500
	err.Err = fmt.Errorf("Message: %s", "test")
	err.Body.Id = 200
	err.Body.StatusCode = 500
	err.Body.Message = "Test message"

	var want events.APIGatewayProxyResponse
	want.Body = "{\"Id\":200,\"StatusCode\":500,\"Message\":\"Test message\"}"
	want.StatusCode = 500

	got, _ := CreateResponseError(err)
	if want.Body != got.Body {
		t.Fatalf("Wanted %s=%s, but got %s", "Body", want.Body, got.Body)
	}
	if want.StatusCode != got.StatusCode {
		t.Fatalf("Wanted %s=%d, but got %d", "StatusCode", want.StatusCode, got.StatusCode)
	}
}

func Test_ResponseError_WithNonApiError(t *testing.T) {
	err := fmt.Errorf("Message: %s", "test")

	var want events.APIGatewayProxyResponse
	want.Body = "{\"Id\":0,\"StatusCode\":500,\"Message\":\"Unknown server error\"}"
	want.StatusCode = http.StatusInternalServerError

	got, _ := CreateResponseError(err)
	if want.Body != got.Body {
		t.Fatalf("Wanted %s=%s, but got %s", "Body", want.Body, got.Body)
	}
	if want.StatusCode != got.StatusCode {
		t.Fatalf("Wanted %s=%d, but got %d", "StatusCode", want.StatusCode, got.StatusCode)
	}
}
