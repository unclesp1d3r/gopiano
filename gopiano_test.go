//go:build integration

package gopiano

import (
	"os"
	"testing"
)

// newClient creates a fresh Client instance for testing.
func newClient(t *testing.T) *Client {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client
}

// newAuthenticatedClient creates a fresh Client instance and authenticates it
// with both partner and user credentials. Returns a fully initialized client
// ready for authenticated API calls.
func newAuthenticatedClient(t *testing.T) *Client {
	t.Helper()
	client := newClient(t)

	_, err := client.AuthPartnerLogin()
	if err != nil {
		t.Fatalf("failed to authenticate partner: %v", err)
	}

	username := os.Getenv("PANDORA_USERNAME")
	if username == "" {
		t.Fatal("PANDORA_USERNAME environment variable is not set")
	}
	password := os.Getenv("PANDORA_PASSWORD")
	if password == "" {
		t.Fatal("PANDORA_PASSWORD environment variable is not set")
	}

	_, err = client.AuthUserLogin(username, password)
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	return client
}

func Test_AuthPartnerLogin_1(t *testing.T) {
	client := newClient(t)
	response, err := client.AuthPartnerLogin()
	if err != nil {
		t.Fatalf("AuthPartnerLogin failed: %v", err)
	}
	if response == nil {
		t.Fatal("AuthPartnerLogin returned nil response")
	}
	if response.Result.PartnerAuthToken == "" {
		t.Error("PartnerAuthToken is empty")
	}
	if response.Result.PartnerID == "" {
		t.Error("PartnerID is empty")
	}
}

func Test_AuthUserLogin_1(t *testing.T) {
	client := newClient(t)

	_, err := client.AuthPartnerLogin()
	if err != nil {
		t.Fatalf("failed to authenticate partner: %v", err)
	}

	username := os.Getenv("PANDORA_USERNAME")
	if username == "" {
		t.Fatal("PANDORA_USERNAME environment variable is not set")
	}
	password := os.Getenv("PANDORA_PASSWORD")
	if password == "" {
		t.Fatal("PANDORA_PASSWORD environment variable is not set")
	}

	response, err := client.AuthUserLogin(username, password)
	if err != nil {
		t.Fatalf("AuthUserLogin failed: %v", err)
	}
	if response == nil {
		t.Fatal("AuthUserLogin returned nil response")
	}
	if response.Result.UserAuthToken == "" {
		t.Error("UserAuthToken is empty")
	}
	if response.Result.UserID == "" {
		t.Error("UserID is empty")
	}
}

func Test_UserCanSubscribe_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserCanSubscribe()
	if err != nil {
		t.Fatalf("UserCanSubscribe failed: %v", err)
	}
	if response == nil {
		t.Fatal("UserCanSubscribe returned nil response")
	}
	// CanSubscribe and IsSubscriber are boolean fields that should always be present
	// regardless of their values, so we just verify the response structure is valid
}

func Test_UserGetBookmarks_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserGetBookmarks()
	if err != nil {
		t.Fatalf("UserGetBookmarks failed: %v", err)
	}
	if response == nil {
		t.Fatal("UserGetBookmarks returned nil response")
	}
	// Artists and Songs are slices that may be empty, but the Result structure should exist
	// We verify the response structure is valid without checking exact counts
}

func Test_UserGetStationList_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserGetStationList(true)
	if err != nil {
		t.Fatalf("UserGetStationList failed: %v", err)
	}
	if response == nil {
		t.Fatal("UserGetStationList returned nil response")
	}
	if response.Result.Checksum == "" {
		t.Error("Checksum is empty")
	}
	// Stations may be empty, but the slice should exist
}

func Test_UserGetStationListChecksum_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserGetStationListChecksum()
	if err != nil {
		t.Fatalf("UserGetStationListChecksum failed: %v", err)
	}
	if response == nil {
		t.Fatal("UserGetStationListChecksum returned nil response")
	}
	if response.Result.Checksum == "" {
		t.Error("Checksum is empty")
	}
}
