package gopiano

import "testing"

func TestUserMethods_MissingPartnerAuthToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		call func(*testing.T, *Client) error
	}{
		{
			name: "UserCreateUser",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.UserCreateUser(
					t.Context(),
					"user@example.com",
					"password",
					"male",
					"US",
					90210,
					1990,
					false,
				)
				return err
			},
		},
		{
			name: "UserEmailPassword",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				return c.UserEmailPassword(t.Context(), "user@example.com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, err := NewClient(AndroidClient)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			err = tt.call(t, client)

			assertMissingTokenError(t, err, "partner authentication token missing", "AuthPartnerLogin")
		})
	}
}

func TestUserMethods_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		call func(*testing.T, *Client) error
	}{
		{
			name: "UserCanSubscribe",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.UserCanSubscribe(t.Context())
				return err
			},
		},
		{
			name: "UserGetBookmarks",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.UserGetBookmarks(t.Context())
				return err
			},
		},
		{
			name: "UserGetStationList",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.UserGetStationList(t.Context(), false)
				return err
			},
		},
		{
			name: "UserGetStationListChecksum",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.UserGetStationListChecksum(t.Context())
				return err
			},
		},
		{
			name: "UserSetQuickMix",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				return c.UserSetQuickMix(t.Context(), []string{"station1"})
			},
		},
		{
			name: "UserSleepSong",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				return c.UserSleepSong(t.Context(), "trackToken123")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			client, err := NewClient(AndroidClient)
			if err != nil {
				t.Fatalf("Failed to create client: %v", err)
			}

			err = tt.call(t, client)

			assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
		})
	}
}
