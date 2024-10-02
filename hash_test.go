package lambda_common

import (
	"strings"
	"testing"
)

func Test_GenerateSecretHash(t *testing.T) {

	hash, err := GenerateSecretHash("username", "clientId", "secret")
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	if len(strings.TrimSpace(hash)) == 0 {
		t.Fatalf("Hash should not be empty")
	}
}
