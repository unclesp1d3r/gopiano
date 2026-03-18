package requests

type (
	// StationGetGenreStations represents the request data for station.getGenreStations.
	StationGetGenreStations userTokenGeneric
	// StationGetGenreStationsChecksum represents the request data for station.getGenreStationsChecksum.
	StationGetGenreStationsChecksum userTokenGeneric
)

// StationCreateStation represents the request data for station.createStation.
type StationCreateStation struct {
	MusicToken    string `json:"musicToken,omitempty"`
	TrackToken    string `json:"trackToken,omitempty"`
	MusicType     string `json:"musicType,omitempty"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationDeleteStation represents the request data for station.deleteStation.
type StationDeleteStation struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationAddFeedback represents the request data for station.addFeedback.
type StationAddFeedback struct {
	TrackToken    string `json:"trackToken"`
	IsPositive    bool   `json:"isPositive"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationDeleteFeedback represents the request data for station.deleteFeedback.
type StationDeleteFeedback struct {
	FeedbackID    string `json:"feedbackId"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationAddMusic represents the request data for station.addMusic.
type StationAddMusic struct {
	MusicToken    string `json:"musicToken"`
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationDeleteMusic represents the request data for station.deleteMusic.
type StationDeleteMusic struct {
	SeedID        string `json:"seedId"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationGetPlaylist represents the request data for station.getPlaylist.
type StationGetPlaylist struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationGetStation represents the request data for station.getStation.
type StationGetStation struct {
	StationToken              string `json:"stationToken"`
	IncludeExtendedAttributes bool   `json:"includeExtendedAttributes,omitempty"`
	SyncTime                  int    `json:"syncTime"`
	UserAuthToken             string `json:"userAuthToken"`
}

// StationShareStation represents the request data for station.shareStation.
type StationShareStation struct {
	StationID     string   `json:"stationId"`
	StationToken  string   `json:"stationToken"`
	Emails        []string `json:"emails"`
	SyncTime      int      `json:"syncTime"`
	UserAuthToken string   `json:"userAuthToken"`
}

// StationRenameStation represents the request data for station.renameStation.
type StationRenameStation struct {
	StationToken  string `json:"stationToken"`
	StationName   string `json:"stationName"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationTransformSharedStation represents the request data for station.transformSharedStation.
type StationTransformSharedStation struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}
