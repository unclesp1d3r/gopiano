package gopiano

import "testing"

func TestStationAddFeedback_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationAddFeedback(t.Context(), "trackToken123", true)

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationAddMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationAddMusic(t.Context(), "musicToken789", "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationCreateStationTrack_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationCreateStationTrack(t.Context(), "trackToken123", "song")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationCreateStationMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationCreateStationMusic(t.Context(), "musicToken789")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteFeedback_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteFeedback(t.Context(), "feedbackID123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteMusic_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteMusic(t.Context(), "seedID123")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationDeleteStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = client.StationDeleteStation(t.Context(), "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetGenreStations_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetGenreStations(t.Context())

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetPlaylist_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetPlaylist(t.Context(), "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationGetStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationGetStation(t.Context(), "stationToken456", false)

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
		t.Context(),
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

	_, err = client.StationRenameStation(t.Context(), "stationToken456", "New Station Name")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}

func TestStationTransformSharedStation_MissingUserAuthToken(t *testing.T) {
	t.Parallel()

	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.StationTransformSharedStation(t.Context(), "stationToken456")

	// Verify error expectations
	assertMissingTokenError(t, err, "user authentication token missing", "AuthUserLogin")
}
