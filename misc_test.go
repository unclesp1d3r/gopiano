package gopiano

import (
	"testing"
)

func TestExplainTrack_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.ExplainTrack("trackToken123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestMusicSearch_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.MusicSearch("search query")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestBookmarkAddArtistBookmark_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.BookmarkAddArtistBookmark("trackToken123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestBookmarkAddSongBookmark_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.BookmarkAddSongBookmark("trackToken123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}
