//go:build integration

package gopiano

import "testing"

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

	// TODO: Not sure if this is a valid cred, but it shouldn't be in the repo. Looks like a real user. We need to replace this with a test user.
	_, err = client.AuthUserLogin("mellowcellofellow@gmail.com", "Great8")
	if err != nil {
		t.Fatalf("failed to authenticate user: %v", err)
	}

	return client
}

func Test_AuthPartnerLogin_1(t *testing.T) {
	client := newClient(t)
	response, err := client.AuthPartnerLogin()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_AuthUserLogin_1(t *testing.T) {
	client := newClient(t)

	_, err := client.AuthPartnerLogin()
	if err != nil {
		t.Fatalf("failed to authenticate partner: %v", err)
	}

	// TODO: Not sure if this is a valid cred, but it shouldn't be in the repo. Looks like a real user. We need to replace this with a test user.
	response, err := client.AuthUserLogin("mellowcellofellow@gmail.com", "Great8")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserCanSubscribe_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserCanSubscribe()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserBetBookmarks_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserGetBookmarks()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationList_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserGetStationList(true)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationListChecksum_1(t *testing.T) {
	client := newAuthenticatedClient(t)
	response, err := client.UserGetStationListChecksum()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}
