package gopiano

import (
	"testing"
)

func TestStationAddFeedback_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationAddFeedback("trackToken123", true)

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationAddMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationAddMusic("musicToken789", "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationCreateStationTrack_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationCreateStationTrack("trackToken123", "song")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationCreateStationMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationCreateStationMusic("musicToken789")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteFeedback_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteFeedback("feedbackID123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteMusic("seedID123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteStation("stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetGenreStations_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetGenreStations()

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetPlaylist_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetPlaylist("stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetStation("stationToken456", false)

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationShareStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationShareStation("stationID123", "stationToken456", []string{"email@example.com"})

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationRenameStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationRenameStation("stationToken456", "New Station Name")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationTransformSharedStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationTransformSharedStation("stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}
