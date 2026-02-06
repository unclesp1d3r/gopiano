package gopiano

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/unclesp1d3r/gopiano/responses"
)

func TestAuthUserLogin_MissingPartnerAuthToken(t *testing.T) {
	t.Parallel()

	// Create a client without calling AuthPartnerLogin()
	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Attempt to login a user without partner authentication
	_, err = client.AuthUserLogin(context.Background(), "user@example.com", "password")

	// Verify error expectations
	assertMissingTokenError(t, err, "partner authentication token missing", "AuthPartnerLogin")
}

// assertMissingTokenError verifies that an error is a standard Go error (not PandoraError)
// and contains the expected substrings for missing token errors.
func assertMissingTokenError(t *testing.T, err error, expectedTokenMsg, expectedMethod string) {
	t.Helper()

	// Verify that an error is returned
	if err == nil {
		t.Fatal("Expected error when token is missing, got nil")
	}

	// Verify that it's NOT a PandoraError (should be a standard Go error)
	var pandoraErr *responses.PandoraError
	if errors.As(err, &pandoraErr) {
		t.Errorf("Expected standard Go error, got PandoraError: %v", err)
	}

	// Verify the error message contains the expected token message
	if err.Error() == "" || !strings.Contains(err.Error(), expectedTokenMsg) {
		t.Errorf("Expected error message to contain %q, got: %q", expectedTokenMsg, err.Error())
	}

	// Verify the error message mentions the expected authentication method
	if !strings.Contains(err.Error(), expectedMethod) {
		t.Errorf("Expected error message to mention %q, got: %q", expectedMethod, err.Error())
	}
}
