package auth

import (
	"net/http"
	"testing"
)

// TestGetAPIKey tests the functionality of GetAPIKey function.
func TestGetAPIKey(t *testing.T) {
	expected := "your-expected-api-key"
	headers := make(http.Header)

	headers.Add("Authorization", "ApiKey "+expected)
	// Test case 1: No Authorization header included
	_, error := GetAPIKey(headers)
	if error != nil {
		t.Errorf("Expected no error, got %v", error)
	}
}
