package gopiano

import "testing"

func TestMiscMethods_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		call func(*testing.T, *Client) error
	}{
		{
			name: "ExplainTrack",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.ExplainTrack(t.Context(), "trackToken123")
				return err
			},
		},
		{
			name: "MusicSearch",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.MusicSearch(t.Context(), "search query")
				return err
			},
		},
		{
			name: "BookmarkAddArtistBookmark",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.BookmarkAddArtistBookmark(t.Context(), "trackToken123")
				return err
			},
		},
		{
			name: "BookmarkAddSongBookmark",
			call: func(t *testing.T, c *Client) error { //nolint:thelper // table-driven test closure, not a helper
				_, err := c.BookmarkAddSongBookmark(t.Context(), "trackToken123")
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
