package requests

type (
	// BookmarkAddArtistBookmark represents the request data for bookmark.addArtistBookmark.
	BookmarkAddArtistBookmark trackAction
	// BookmarkAddSongBookmark represents the request data for bookmark.addSongBookmark.
	BookmarkAddSongBookmark trackAction
)

// ExplainTrack represents the request data for track.explainTrack.
type ExplainTrack struct {
	TrackToken    string `json:"trackToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// MusicSearch represents the request data for music.search.
type MusicSearch struct {
	SearchText    string `json:"searchText"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}
