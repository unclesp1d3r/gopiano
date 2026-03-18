package responses

// StationAddFeedback represents the response from station.addFeedback.
type StationAddFeedback struct {
	Result FeedbackResponse `json:"result"`
}

// StationAddMusic represents the response from station.addMusic.
type StationAddMusic struct {
	Result struct {
		ArtistName  string       `json:"artistName"`
		DateCreated DateResponse `json:"dateCreated"`
		SeedID      string       `json:"seedId"`
	} `json:"result"`
}

// GenreStation represents a single station within a genre category.
type GenreStation struct {
	StationToken string `json:"stationToken"`
	StationName  string `json:"stationName"`
	StationID    string `json:"stationId"`
}

// GenreCategory represents a category of genre stations.
type GenreCategory struct {
	CategoryName string         `json:"categoryName"`
	Stations     []GenreStation `json:"stations"`
}

// StationGetGenreStations represents the response from station.getGenreStations.
type StationGetGenreStations struct {
	Result struct {
		Categories []GenreCategory `json:"categories"`
	} `json:"result"`
}

// StationGetGenreStationsChecksum represents the response from station.getGenreStationsChecksum.
type StationGetGenreStationsChecksum struct {
	Result struct {
		Checksum string `json:"checksum"`
	} `json:"result"`
}

// AudioURL represents an audio URL entry in a playlist item's AudioURLMap.
type AudioURL struct {
	Bitrate  string `json:"bitrate"`
	Encoding string `json:"encoding"`
	AudioURL string `json:"audioUrl"`
	Protocol string `json:"protocol"`
}

// PlaylistItem represents a single track in a station playlist.
type PlaylistItem struct {
	TrackToken             string              `json:"trackToken"`
	ArtistName             string              `json:"artistName"`
	AlbumName              string              `json:"albumName"`
	AmazonAlbumURL         string              `json:"amazonAlbumUrl"`
	SongExplorerURL        string              `json:"songExplorerUrl"`
	AlbumArtURL            string              `json:"albumArtUrl"`
	ArtistDetailURL        string              `json:"artistDetailUrl"`
	AudioURLMap            map[string]AudioURL `json:"audioUrlMap"`
	ITunesSongURL          string              `json:"itunesSongUrl"`
	AmazonAlbumAsin        string              `json:"amazonAlbumAsin"`
	AmazonAlbumDigitalAsin string              `json:"amazonAlbumDigitalAsin"`
	ArtistExplorerURL      string              `json:"artistExplorerUrl"`
	SongName               string              `json:"songName"`
	AlbumDetailURL         string              `json:"albumDetailUrl"`
	SongDetailURL          string              `json:"songDetailUrl"`
	StationID              string              `json:"stationId"`
	SongRating             int                 `json:"songRating"`
	TrackGain              string              `json:"trackGain"`
	AlbumExplorerURL       string              `json:"albumExplorerUrl"`
	AllowFeedback          bool                `json:"allowFeedback"`
	AmazonSongDigitalAsin  string              `json:"amazonSongDigitalAsin"`
	NowPlayingStationAdURL string              `json:"nowPlayingStationAdUrl"`
	AdToken                string              `json:"adToken"`
}

// StationGetPlaylist represents the response from station.getPlaylist.
type StationGetPlaylist struct {
	Result struct {
		Items []PlaylistItem `json:"items"`
	} `json:"result"`
}

// ExplainTrack represents the response from track.explainTrack.
type ExplainTrack struct {
	Result struct {
		Explanations []struct {
			FocustTraitName string `json:"focusTraitName"`
			FocusTraitID    string `json:"focustTraitId"`
		} `json:"explanations"`
	} `json:"result"`
}
