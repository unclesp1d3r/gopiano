package responses

import (
	"cmp"
	"slices"
	"time"
)

// DateResponse is used repeatedly in places where Pandora returns a JSON object
// called dateCreated.
// Most of the data is rubbish without a little processing but you can use GetDate()
// and also Time is just a nice UNIX epoch.
type DateResponse struct {
	Nanos          int `json:"nano"`
	Seconds        int `json:"seconds"`
	Year           int `json:"year"`
	Month          int `json:"month"`
	Hours          int `json:"hours"`
	Time           int `json:"time"`
	Date           int `json:"date"`
	Minutes        int `json:"minutes"`
	Day            int `json:"day"`
	TimezoneOffset int `json:"timezoneOffset"`
}

// GetDate converts the DateResponse to a time.Time object.
// The 1900+Year offset follows Java's deprecated Date.getYear() convention,
// which returns year minus 1900. The Month+1 offset follows Java's deprecated
// Date.getMonth() convention, which returns 0-indexed months (0=January).
// The Pandora API uses these Java conventions.
func (d DateResponse) GetDate() time.Time {
	return time.Date(1900+d.Year, time.Month(d.Month+1), d.Date, d.Hours, d.Minutes, d.Seconds,
		d.Nanos, time.FixedZone("Local Time", d.TimezoneOffset*60)) //nolint:mnd // 60 seconds per minute
}

// FeedbackResponse represents feedback data in API responses.
type FeedbackResponse struct {
	ArtistName  string       `json:"artistName"`
	SongName    string       `json:"songName"`
	DateCreated DateResponse `json:"dateCreated"`
	FeedbackID  string       `json:"feedbackId"`
	IsPositive  bool         `json:"isPositive"`
}

// ArtistBookmark represents an artist bookmark in API responses.
type ArtistBookmark struct {
	ArtURL        string       `json:"artUrl"`
	ArtistName    string       `json:"artistName"`
	BookmarkToken string       `json:"bookmarkToken"`
	DateCreated   DateResponse `json:"dateCreated"`
	MusicToken    string       `json:"musicToken"`
}

// SongBookmark represents a song bookmark in API responses.
type SongBookmark struct {
	AlbumName     string       `json:"albumName"`
	ArtURL        string       `json:"artUrl"`
	ArtistName    string       `json:"artistName"`
	BookmarkToken string       `json:"bookmarkToken"`
	DateCreated   DateResponse `json:"dateCreated"`
	MusicToken    string       `json:"musicToken"`
	SampleGain    string       `json:"sampleGain"`
	SampleURL     string       `json:"sampleUrl"`
	SongName      string       `json:"songName"`
}

// Station represents a Pandora station in API responses.
type Station struct {
	SuppressVideoAds bool         `json:"suppressVideoAds"`
	StationID        string       `json:"stationId"`
	AllowAddMusic    bool         `json:"allowAddMusic"`
	DateCreated      DateResponse `json:"dateCreated"`
	StationDetailURL string       `json:"stationDetailUrl"`
	ArtURL           string       `json:"artUrl"`
	RequiresCleanAds bool         `json:"requiresCleanAds"`
	StationToken     string       `json:"stationToken"`
	StationName      string       `json:"stationName"`
	Music            struct {
		Songs []struct {
			SeedID      string       `json:"seedId"`
			ArtistName  string       `json:"artistName"`
			SongName    string       `json:"songName"`
			DateCreated DateResponse `json:"dateCreated"`
		} `json:"songs"`
		Artists []struct {
			SeedID      string       `json:"seedId"`
			ArtistName  string       `json:"artistName"`
			DateCreated DateResponse `json:"dateCreated"`
		} `json:"artists"`
	} `json:"music"`
	IsShared           bool     `json:"isShared"`
	AllowDelete        bool     `json:"allowDelete"`
	Genre              []string `json:"genre"`
	IsQuickMix         bool     `json:"isQuickMix"`
	AllowRename        bool     `json:"allowRename"`
	StationSharingURL  string   `json:"stationSharingUrl"`
	QuickMixStationIDs []string `json:"quickMixStationIds"`
	Feedback           struct {
		ThumbsDown []FeedbackResponse `json:"thumbsDown"`
		ThumbsUp   []FeedbackResponse `json:"thumbsUp"`
	} `json:"feedback"`
}

// StationList represents a list of stations that implements sort.Interface.
type StationList []Station

// SortByName sorts the station list alphabetically by station name.
func (s StationList) SortByName() {
	slices.SortFunc([]Station(s), func(a, b Station) int {
		return cmp.Compare(a.StationName, b.StationName)
	})
}

// Deprecated: Use SortByName instead. Len implements sort.Interface.
func (s StationList) Len() int {
	return len(s)
}

// Deprecated: Use SortByName instead. Swap implements sort.Interface.
func (s StationList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Deprecated: Use SortByName instead. Less implements sort.Interface.
func (s StationList) Less(i, j int) bool {
	return s[i].StationName < s[j].StationName
}

// StationResponse represents a generic station response wrapper.
type StationResponse struct {
	Result Station `json:"result"`
}

type (
	// StationCreateStation represents the response from station.createStation.
	StationCreateStation StationResponse
	// StationGetStation represents the response from station.getStation.
	StationGetStation StationResponse
	// StationRenameStation represents the response from station.renameStation.
	StationRenameStation StationResponse
	// StationTransformSharedStation represents the response from station.transformSharedStation.
	StationTransformSharedStation StationResponse
)
