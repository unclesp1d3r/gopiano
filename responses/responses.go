//nolint:gochecknoglobals // part of public API

// Package responses provides structs used with json.Unmarshal in processing responses from the Pandora API.
package responses

import (
	"fmt"
	"time"
)

// ErrorCodeMap maps Pandora API error codes to their string names.
//
// Error code 0 (INTERNAL) is a generic error that often indicates authentication issues,
// invalid parameters, or rate limiting. Error codes in the 1000+ range are more specific
// and actionable. When receiving error code 0, check authentication flow and parameter
// validation first.
var ErrorCodeMap = map[int]string{ //nolint:gochecknoglobals // part of public API
	0:    "INTERNAL",
	1:    "MAINTENCANCE_MODE",
	2:    "URL_PARAM_MISSING_METHOD",
	3:    "URL_PARAM_MISSING_AUTH_TOKEN",
	4:    "URL_PARAM_MISING_PARTNER_ID",
	5:    "URL_PARAM_MISSING_USER_ID",
	6:    "SECURE_PROTOCOL_REQUIRED",
	7:    "CERTIFICATE_REQUIRED",
	8:    "PARAMATER_TYPE_MISMATCH",
	9:    "PARAMETER_MISSING",
	10:   "PARAMETER_VALUE_INVALID",
	11:   "API_VERSION_NOT_SUPPORTED",
	12:   "LICENSING_RESTRICTIONS",
	13:   "INSUFFICIENT_CONNECTIVITY",
	14:   "UNKNOWN_METHOD_NAME",
	15:   "WRONG_PROTOCOL",
	1000: "READ_ONLY_MODE",
	1001: "INVALID_AUTH_TOKEN",
	1002: "INVALID_PARTNER_LOGIN",
	1003: "LISTENER_NOT_AUTHORIZED",
	1004: "USER_NOT_AUTHORIZED",
	1005: "MAX_STATIONS_REACHED",
	1006: "STATION_DOES_NOT_EXIST",
	1007: "COMPLIMENTARY_PERIOD_ALREADY_IN_USE",
	1008: "CALL_NOT_ALLOWED",
	1009: "DEVICE_NOT_FOUND",
	1010: "PARTNER_NOT_AUTHORIZED",
	1011: "INVALID_USERNAME",
	1012: "INVALID_PASSWORD",
	1013: "USERNAME_ALREADY_EXISTS",
	1014: "DEVICE_ALREADY_ASSOCIATED_TO_ACCOUNT",
	1015: "UPGRADE_DEVICE_MODEL_INVALID",
	1018: "EXPLICIT_PIN_INCORRECT",
	1020: "EXPLICIT_PIN_MALFORMED",
	1023: "DEVICE_MODEL_INVALID",
	1024: "ZIP_CODE_INVALID",
	1025: "BIRTH_YEAR_INVALID",
	1026: "BIRTH_YEAR_TOO_YOUNG",
	1027: "INVALID_COUNTRY_CODE or INVALID_GENDER",
	1034: "DEVICE_DISABLED",
	1035: "DAILY_TRIAL_LIMIT_REACHED",
	1036: "INVALID_SPONSOR",
	1037: "USER_ALREADY_USED_TRIAL",
	1039: "PLAYLIST_EXCEEDED",
}

// PandoraError represents an API error response from Pandora.
type PandoraError struct {
	Stat    string `json:"stat"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e PandoraError) Error() string {
	return fmt.Sprintf("Pandora Error: %d %s", e.Code, e.Message)
}

// GetErrorGuidance returns troubleshooting tips for common error codes.
// This function provides contextual help to diagnose the root cause of API errors,
// especially for generic error code 0 (INTERNAL).
func GetErrorGuidance(code int) string {
	switch code {
	case 0:
		return "Troubleshooting: Check that authentication prerequisites are met (partner login completed), verify all required parameters are valid, ensure calling from a US IP address, consider rate limiting if making frequent requests, note that this is an unofficial API that may have restrictions"
	default:
		return ""
	}
}

// DateResponse is used repeatedly in places where Pandora returns a JSON object
// called dateCreated.
// Most of the data is rubish without a little processing but you can use GetDate()
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
func (d DateResponse) GetDate() time.Time {
	return time.Date(1900+d.Year, time.Month(d.Month), d.Date, d.Hours, d.Minutes, d.Seconds,
		d.Nanos, time.FixedZone("Local Time", d.TimezoneOffset*60)) //nolint:mnd // 60 seconds per minute
}

// AuthPartnerLogin represents the response from auth.partnerLogin.
type AuthPartnerLogin struct {
	Result struct {
		SyncTime         string `json:"syncTime"`
		StationSkipLimit int    `json:"stationSkipLimit"`
		PartnerAuthToken string `json:"partnerAuthToken"`
		PartnerID        string `json:"partnerId"`
		StationSkipUnit  string `json:"stationSkipUnit"`
		DeviceProperties struct {
			VideoAdRefreshInterval int `json:"videoAdRefreshInterval"`
			VideoAdUniqueInterval  int `json:"videoAdUniqueInterval"`
			AdRefreshInterval      int `json:"adRefreshInterval"`
			VideoAdStartInterval   int `json:"videoAdStartInterval"`
		} `json:"deviceProperties"`
		Urls struct {
			AutoComplete string `json:"autoComplete"`
		} `json:"urls"`
	}
}

// AuthUserLogin represents the response from auth.userLogin.
type AuthUserLogin struct {
	Result struct {
		CanListen                   bool   `json:"canListen"`
		HasAudioAds                 bool   `json:"hasAudioAds"`
		IsCapped                    bool   `json:"isCapped,omitempty"`
		ListeningTimeoutAlertMsgUri string `json:"listeningTimeoutAlertMsgUri"` //nolint:revive,staticcheck // matches Pandora API
		ListeningTimeoutMinutes     string `json:"listeningTimeoutMinutes"`
		MaxStationsAllowed          int    `json:"maxStationsAllowed"`
		MinimumAdRefreshInterval    int    `json:"minimumAdRefreshInterval"`
		NowPlayingURL               string `json:"nowPlayingUrl"`
		SplashScreenAdURL           string `json:"splashScreenAdUrl"`
		StationCreationAdURL        string `json:"stationCreationAdUrl"`
		UserAuthToken               string `json:"userAuthToken"`
		UserID                      string `json:"userId"`
		UserProfileURL              string `json:"userProfileUrl"`
		Username                    string `json:"username"`
		VideoAdURL                  string `json:"videoAdUrl"`
	} `json:"result"`
}

// UserCanSubscribe represents the response from user.canSubscribe.
type UserCanSubscribe struct {
	Result struct {
		CanSubscribe bool `json:"canSubscribe"`
		IsSubscriber bool `json:"isSubscriber"`
	} `json:"result"`
}

// UserCreateUser represents the response from user.createUser.
type UserCreateUser AuthUserLogin

// ArtistBookmark represents an artist bookmark in API responses.
type ArtistBookmark struct {
	ArtURL        string       `json:"artUrl"`
	ArtistName    string       `json:"artistName"`
	BookmarkToken string       `json:"bookmarkToken"`
	DateCreated   DateResponse `json:"dateCreated"`
	MusicToken    string       `json:"musicToken"`
}

// BookmarkAddArtistBookmark represents the response from bookmark.addArtistBookmark.
type BookmarkAddArtistBookmark struct {
	Result ArtistBookmark `json:"result"`
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

// BookmarkAddSongBookmark represents the response from bookmark.addSongBookmark.
type BookmarkAddSongBookmark struct {
	Result SongBookmark `json:"result"`
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

// Len returns the number of stations in the list.
func (s StationList) Len() int {
	return len(s)
}

func (s StationList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StationList) Less(i, j int) bool {
	return s[i].StationName < s[j].StationName
}

// UserGetBookmarks represents the response from user.getBookmarks.
type UserGetBookmarks struct {
	Result struct {
		Artists []ArtistBookmark `json:"artists"`
		Songs   []SongBookmark   `json:"songs"`
	} `json:"result"`
}

// UserGetStationList represents the response from user.getStationList.
type UserGetStationList struct {
	Result struct {
		Stations StationList `json:"stations"`
		Checksum string      `json:"checksum"`
	} `json:"result"`
}

// UserGetStationListChecksum represents the response from user.getStationListChecksum.
type UserGetStationListChecksum struct {
	Result struct {
		Checksum string `json:"checksum"`
	} `json:"result"`
}

// MusicSearch represents the response from music.search.
type MusicSearch struct {
	Result struct {
		NearMatchesAvailable bool   `json:"nearMatchesAvailable"`
		Explanation          string `json:"explanation"`
		Songs                []struct {
			ArtistName string `json:"artistName"`
			MusicToken string `json:"musicToken"`
			SongName   string `json:"songName"`
			Score      int    `json:"score"`
		} `json:"songs"`
		Artists []struct {
			ArtistName  string `json:"artistName"`
			MusicToken  string `json:"musicToken"`
			LikelyMatch bool   `json:"likelyMatch"`
			Score       int    `json:"score"`
		} `json:"artists"`
	} `json:"result"`
}

// FeedbackResponse represents feedback data in API responses.
type FeedbackResponse struct {
	ArtistName  string       `json:"artistName"`
	SongName    string       `json:"songName"`
	DateCreated DateResponse `json:"dateCreated"`
	FeedbackID  string       `json:"feedbackId"`
	IsPositive  bool         `json:"isPositive"`
}

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

// StationGetGenreStations represents the response from station.getGenreStations.
type StationGetGenreStations struct {
	Result struct {
		Categories []struct {
			CategoryName string `json:"categoryName"`
			Stations     []struct {
				StationToken string `json:"stationToken"`
				StationName  string `json:"stationName"`
				StationID    string `json:"stationId"`
			}
		} `json:"categories"`
	} `json:"result"`
}

// StationGetGenreStationsChecksum represents the response from station.getGenreStationsChecksum.
type StationGetGenreStationsChecksum struct {
	Result struct {
		Checksum string `json:"checksum"`
	} `json:"result"`
}

// StationGetPlaylist represents the response from station.getPlaylist.
type StationGetPlaylist struct {
	Result struct {
		Items []struct {
			TrackToken      string `json:"trackToken"`
			ArtistName      string `json:"artistName"`
			AlbumName       string `json:"albumName"`
			AmazonAlbumURL  string `json:"amazonAlbumUrl"`
			SongExplorerURL string `json:"songExplorerUrl"`
			AlbumArtURL     string `json:"albumArtUrl"`
			ArtistDetailURL string `json:"artistDetailUrl"`
			AudioURLMap     map[string]struct {
				Bitrate  string `json:"bitrate"`
				Encoding string `json:"encoding"`
				AudioURL string `json:"audioUrl"`
				Protocol string `json:"protocol"`
			} `json:"audioUrlMap"`
			ITunesSongURL          string `json:"itunesSongUrl"`
			AmazonAlbumAsin        string `json:"amazonAlbumAsin"`
			AmazonAlbumDigitalAsin string `json:"amazonAlbumDigitalAsin"`
			ArtistExplorerURL      string `json:"artistExplorerUrl"`
			SongName               string `json:"songName"`
			AlbumDetailURL         string `json:"albumDetailUrl"`
			SongDetailURL          string `json:"songDetailUrl"`
			StationID              string `json:"stationId"`
			SongRating             int    `json:"songRating"`
			TrackGain              string `json:"trackGain"`
			AlbumExplorerURL       string `json:"albumExplorerUrl"`
			AllowFeedback          bool   `json:"allowFeedback"`
			AmazonSongDigitalAsin  string `json:"amazonSongDigitalAsin"`
			NowPlayingStationAdURL string `json:"nowPlayingStationAdUrl"`
			AdToken                string `json:"adToken"`
		} `json:"items"`
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
