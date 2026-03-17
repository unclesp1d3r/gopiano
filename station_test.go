package gopiano

import "testing"

func TestStationMethods_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		call func(*testing.T, *Client) error
	}{
		{
			name: "StationAddFeedback",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationAddFeedback(t.Context(), "trackToken123", true)
				return err
			},
		},
		{
			name: "StationAddMusic",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationAddMusic(t.Context(), "musicToken789", "stationToken456")
				return err
			},
		},
		{
			name: "StationCreateStationTrack",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationCreateStationTrack(t.Context(), "trackToken123", "song")
				return err
			},
		},
		{
			name: "StationCreateStationMusic",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationCreateStationMusic(t.Context(), "musicToken789")
				return err
			},
		},
		{
			name: "StationDeleteFeedback",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				return c.StationDeleteFeedback(t.Context(), "feedbackID123")
			},
		},
		{
			name: "StationDeleteMusic",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				return c.StationDeleteMusic(t.Context(), "seedID123")
			},
		},
		{
			name: "StationDeleteStation",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				return c.StationDeleteStation(t.Context(), "stationToken456")
			},
		},
		{
			name: "StationGetGenreStations",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationGetGenreStations(t.Context())
				return err
			},
		},
		{
			name: "StationGetPlaylist",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationGetPlaylist(t.Context(), "stationToken456")
				return err
			},
		},
		{
			name: "StationGetStation",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationGetStation(t.Context(), "stationToken456", false)
				return err
			},
		},
		{
			name: "StationShareStation",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				return c.StationShareStation(
					t.Context(),
					"stationID123",
					"stationToken456",
					[]string{"email@example.com"},
				)
			},
		},
		{
			name: "StationRenameStation",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationRenameStation(t.Context(), "stationToken456", "New Station Name")
				return err
			},
		},
		{
			name: "StationTransformSharedStation",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.StationTransformSharedStation(t.Context(), "stationToken456")
				return err
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
