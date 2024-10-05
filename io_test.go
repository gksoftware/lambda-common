package lambda_common

import (
	"github.com/aws/aws-lambda-go/events"
	"strings"
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

func Test_ExtractCooikes_WithMultipleCookies(t *testing.T) {
	cookieString := "access_token=eyJr; id_token=HkiO; refresh_token=NNIiwiYWxnIjoi"

	cookies := *ExtractCookies(cookieString)

	if len(cookies) != 3 {
		t.Fatalf("Wanted %d cookies, but got %d", 3, len(cookies))
	}

	access_token := cookies["access_token"]
	if len(strings.TrimSpace(access_token)) == 0 {
		t.Fatalf("Wanted access_token, but got nil")
	}

	if access_token != "eyJr" {
		t.Fatalf("Wanted access_token=eyJr, but got %s", access_token)
	}

	refresh_token := cookies["refresh_token"]
	if len(strings.TrimSpace(refresh_token)) == 0 {
		t.Fatalf("Wanted refresh_token, but got nil")
	}

	if refresh_token != "NNIiwiYWxnIjoi" {
		t.Fatalf("Wanted access_token=NNIiwiYWxnIjoi, but got %s", refresh_token)
	}

	id_token := cookies["id_token"]
	if len(strings.TrimSpace(id_token)) == 0 {
		t.Fatalf("Wanted id_token, but got nil")
	}

	if id_token != "HkiO" {
		t.Fatalf("Wanted id_token=HkiO, but got %s", id_token)
	}
}

func Test_ExtractCooikes_WithSingleCookie(t *testing.T) {
	cookieString := "access_token=eyJr;"

	cookies := *ExtractCookies(cookieString)

	if len(cookies) != 1 {
		t.Fatalf("Wanted %d cookies, but got %d", 3, len(cookies))
	}

	access_token := cookies["access_token"]
	if len(strings.TrimSpace(access_token)) == 0 {
		t.Fatalf("Wanted access_token, but got nil")
	}

	if access_token != "eyJr" {
		t.Fatalf("Wanted access_token=eyJr, but got %s", access_token)
	}

}
