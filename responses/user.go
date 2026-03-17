package responses

// UserCanSubscribe represents the response from user.canSubscribe.
type UserCanSubscribe struct {
	Result struct {
		CanSubscribe bool `json:"canSubscribe"`
		IsSubscriber bool `json:"isSubscriber"`
	} `json:"result"`
}

// UserCreateUser represents the response from user.createUser.
type UserCreateUser AuthUserLogin

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

// BookmarkAddArtistBookmark represents the response from bookmark.addArtistBookmark.
type BookmarkAddArtistBookmark struct {
	Result ArtistBookmark `json:"result"`
}

// BookmarkAddSongBookmark represents the response from bookmark.addSongBookmark.
type BookmarkAddSongBookmark struct {
	Result SongBookmark `json:"result"`
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
