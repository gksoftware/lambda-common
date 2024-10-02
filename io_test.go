package lambda_common

import (
	"github.com/aws/aws-lambda-go/events"
	"testing"
)

type test_struct struct {
	Age   int    `json:"age"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func Test_UnmarshalRequest(t *testing.T) {
	var want test_struct
	want.Age = 35
	want.Email = "a@b.com"
	want.Name = "Bob"

	var got test_struct
	var req events.APIGatewayProxyRequest
	req.Body = "{\"Age\":35,\"Name\":\"Bob\",\"Email\":\"a@b.com\"}"

	if err := UnmarshalRequest(req, &got); err != nil {
		t.Fatalf("Unexpected error")
	}

	if want.Age != got.Age {
		t.Fatalf("Wanted %s=%d, but got %d", "Age", want.Age, got.Age)
	}

	if want.Email != got.Email {
		t.Fatalf("Wanted %s=%s, but got %s", "Email", want.Email, got.Email)
	}

	if want.Name != got.Name {
		t.Fatalf("Wanted %s=%s, but got %s", "Name", want.Name, got.Name)
	}

}

func Test_MarshalResponse(t *testing.T) {
	var payload test_struct
	payload.Age = 35
	payload.Email = "a@b.com"
	payload.Name = "Bob"

	want := "{\"age\":35,\"email\":\"a@b.com\",\"name\":\"Bob\"}"

	got, err := MarshalResponse(&payload)
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	if want != got {
		t.Fatalf("Wanted %s, but got %s", want, got)
	}

}
