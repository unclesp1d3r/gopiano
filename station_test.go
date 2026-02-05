package gopiano

import (
	"context"
	"testing"
)

func TestStationAddFeedback_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationAddFeedback(context.Background(), "trackToken123", true)

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationAddMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationAddMusic(context.Background(), "musicToken789", "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationCreateStationTrack_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationCreateStationTrack(context.Background(), "trackToken123", "song")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationCreateStationMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationCreateStationMusic(context.Background(), "musicToken789")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteFeedback_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteFeedback(context.Background(), "feedbackID123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteMusic(context.Background(), "seedID123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteStation(context.Background(), "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetGenreStations_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetGenreStations(context.Background())

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetPlaylist_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetPlaylist(context.Background(), "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetStation(context.Background(), "stationToken456", false)

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationShareStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationShareStation(
		context.Background(),
		"stationID123",
		"stationToken456",
		[]string{"email@example.com"},
	)

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationRenameStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationRenameStation(context.Background(), "stationToken456", "New Station Name")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationTransformSharedStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationTransformSharedStation(context.Background(), "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}
