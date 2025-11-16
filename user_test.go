package gopiano

import (
	"testing"
)

func TestUserCreateUser_MissingPartnerAuthToken(t *testing.T) {
	t.Parallel()

	// Create a client without calling AuthPartnerLogin()
	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Attempt to create a user without partner authentication
	_, err = client.UserCreateUser(
		"user@example.com",
		"password",
		"male",
		"US",
		90210,
		1990,
		false,
	)

	// Verify error expectations
	assertMissingTokenError(t, err, "partner authentication token missing", "AuthPartnerLogin")
}

func TestUserEmailPassword_MissingPartnerAuthToken(t *testing.T) {
	t.Parallel()

	// Create a client without calling AuthPartnerLogin()
	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Attempt to resend registration email without partner authentication
	err = client.UserEmailPassword("user@example.com")

	// Verify error expectations
	assertMissingTokenError(t, err, "partner authentication token missing", "AuthPartnerLogin")
}

func TestUserCanSubscribe_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	// Create a client and authenticate partner but not user
	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Note: We can't actually call AuthPartnerLogin() in a unit test without network,
	// but we can test the validation logic by checking the error when userAuthToken is empty
	// For this test, we'll simulate having partner auth but not user auth
	// by creating a client and not calling AuthUserLogin() or UserCreateUser()

	// Attempt to check subscription status without user authentication
	_, err = client.UserCanSubscribe()

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestUserGetBookmarks_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.UserGetBookmarks()

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestUserGetStationList_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.UserGetStationList(false)

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestUserGetStationListChecksum_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.UserGetStationListChecksum()

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestUserSetQuickMix_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.UserSetQuickMix([]string{"station1"})

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestUserSleepSong_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.UserSleepSong("trackToken123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}
