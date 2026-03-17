package requests

type userTokenGeneric struct {
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type (
	// UserGetBookmarks represents the request data for user.getBookmarks.
	UserGetBookmarks userTokenGeneric
	// UserGetStationListChecksum represents the request data for user.getStationListChecksum.
	UserGetStationListChecksum userTokenGeneric
	// UserCanSubscribe represents the request data for user.canSubscribe.
	UserCanSubscribe userTokenGeneric
)

// UserCreateUser represents the request data for user.createUser.
type UserCreateUser struct {
	AccountType      string `json:"accountType"`
	BirthYear        int    `json:"birthYear"`
	CountryCode      string `json:"countryCode"`
	EmailOptin       bool   `json:"emailOptin"`
	Gender           string `json:"gender"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	Password         string `json:"password"`
	RegisteredType   string `json:"registeredType"`
	SyncTime         int    `json:"syncTime"`
	Username         string `json:"username"`
	ZipCode          int    `json:"zip"`
}

// UserEmailPassword represents the request data for user.emailPassword.
type UserEmailPassword struct {
	PartnerAuthToken string `json:"partnerAuthToken"`
	SyncTime         int    `json:"syncTime"`
	Username         string `json:"username"`
}

// UserGetStationList represents the request data for user.getStationList.
type UserGetStationList struct {
	IncludeStationArtURL bool   `json:"includeStationArtUrl,omitempty"`
	SyncTime             int    `json:"syncTime"`
	UserAuthToken        string `json:"userAuthToken"`
}

// UserSetQuickMix represents the request data for user.setQuickMix.
type UserSetQuickMix struct {
	QuickMixStationIDs []string `json:"quickMixStationIds"`
	SyncTime           int      `json:"syncTime"`
	UserAuthToken      string   `json:"userAuthToken"`
}

type trackAction struct {
	TrackToken    string `json:"trackToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type (
	// UserSleepSong represents the request data for user.sleepSong.
	UserSleepSong trackAction
)
