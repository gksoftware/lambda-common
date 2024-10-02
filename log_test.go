package lambda_common

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

func Test_LogRequestPretty(t *testing.T) {
	var e events.APIGatewayProxyRequest
	e.Body = "{\"age\":35,\"email\":\"a@b.com\",\"name\":\"Bob\"}"
	e.HTTPMethod = "GET"
	e.IsBase64Encoded = false
	LogRequestPretty(e)
}

func Test_LogRequest(t *testing.T) {
	var e events.APIGatewayProxyRequest
	e.Body = "{\"age\":35,\"email\":\"a@b.com\",\"name\":\"Bob\"}"
	e.HTTPMethod = "GET"
	e.IsBase64Encoded = false
	LogRequest(e)
}
