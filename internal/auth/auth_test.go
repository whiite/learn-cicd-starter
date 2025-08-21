package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKeyNoAuth(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err == nil {
		t.Fatalf("expected error got key")
	}
}

func TestGetApiKeyAuthHeader(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey testing")
	_, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("expected key, got error: %v", err)
	}
}
