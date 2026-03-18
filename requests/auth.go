// Package requests provides structs for use with json.Marshal when sending requests to the Pandora API.
package requests

// AuthPartnerLogin represents the request data for auth.partnerLogin.
type AuthPartnerLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`
	IncludeURLs bool   `json:"includeUrls,omitempty"`
}

// AuthUserLogin represents the request data for auth.userLogin.
type AuthUserLogin struct {
	PartnerAuthToken              string `json:"partnerAuthToken"`
	Username                      string `json:"username"`
	Password                      string `json:"password"`
	LoginType                     string `json:"loginType"` // Should always be "user"
	SyncTime                      int    `json:"syncTime"`
	IncludeAdAttributes           bool   `json:"includeAdAttributes,omitempty"`
	IncludeDemographics           bool   `json:"IncludeDemographics,omitempty"`   //nolint:tagliatelle // matches Pandora API
	IncludePandoraOneInfo         bool   `json:"includePandoraOneInfo,omitempty"` // Appears to do nothing.
	IncludeStationArtURL          bool   `json:"includeStationArtUrl,omitempty"`
	IncludeSubscriptionExpiration bool   `json:"includeSubscriptionExpiration,omitempty"`
	ReturnCapped                  bool   `json:"returnCapped,omitempty"`
	ReturnGenreStations           bool   `json:"returnGenreStations,omitempty"`
	ReturnStationList             bool   `json:"returnStationList,omitempty"`
}
