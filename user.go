package gopiano

import (
	"bytes"
	"encoding/json"

	"github.com/unclesp1d3r/gopiano/requests"
	"github.com/unclesp1d3r/gopiano/responses"
)

// UserCanSubscribe returns whether a user is subscribed or can subscribe
// to the premium Pandora One service.
// Calls API method "user.canSubscribe".
func (c *Client) UserCanSubscribe() (*responses.UserCanSubscribe, error) {
	if err := c.validateUserAuthToken("checking subscription status"); err != nil {
		return nil, err
	}
	requestData := requests.UserCanSubscribe{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserCanSubscribe
	err = c.BlowfishCall("http://", "user.canSubscribe", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UserCreateUser creates a new Pandora user.
//
// Prerequisite: Must call AuthPartnerLogin() first to obtain a partner authentication token.
// This function establishes user authentication and sets userAuthToken on the client.
//
// Parameter requirements:
//   - username must be a valid email address
//   - gender must be exactly "male" or "female"
//   - countryCode must be "US" (API restriction)
//   - zipCode must be a valid US ZIP code
//   - birthYear must meet minimum age requirements
//
// Known limitations:
//   - Requires US IP address due to licensing restrictions
//   - May fail with rate limiting if called too frequently
//   - Legacy API endpoint that may be deprecated
//
// See examples/create_user/ for a complete usage example.
//
// Calls API method "user.createUser".
func (c *Client) UserCreateUser(
	username, password, gender, countryCode string,
	zipCode, birthYear int,
	emailOptin bool,
) (*responses.UserCreateUser, error) {
	if err := c.validatePartnerAuthToken("creating a user"); err != nil {
		return nil, err
	}
	requestData := requests.UserCreateUser{
		PartnerAuthToken: c.partnerAuthToken,
		AccountType:      "registered",
		RegisteredType:   "user",
		Username:         username,
		Password:         password,
		Gender:           gender,
		ZipCode:          zipCode,
		CountryCode:      countryCode,
		BirthYear:        birthYear,
		EmailOptin:       emailOptin,
		SyncTime:         c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserCreateUser
	err = c.BlowfishCall("https://", "user.createUser", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.Result.UserAuthToken
	c.userID = resp.Result.UserID

	return &resp, nil
}

// UserEmailPassword resends the registration email.
// Calls API method "user.emailPassword".
func (c *Client) UserEmailPassword(username string) error {
	if err := c.validatePartnerAuthToken("resending registration email"); err != nil {
		return err
	}
	requestData := requests.UserEmailPassword{
		Username:         username,
		PartnerAuthToken: c.partnerAuthToken,
		SyncTime:         c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp interface{}
	return c.BlowfishCall("https://", "user.emailPassword", requestDataReader, &resp)
}

// UserGetBookmarks returns the user's bookmarked artists and songs.
// Also see BookmarkAddArtistBookmark and BookmarkAddSongBookmark.
// Calls API method "user.getBookmarks".
func (c *Client) UserGetBookmarks() (*responses.UserGetBookmarks, error) {
	if err := c.validateUserAuthToken("retrieving bookmarks"); err != nil {
		return nil, err
	}
	requestData := requests.UserGetBookmarks{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserGetBookmarks
	err = c.BlowfishCall("http://", "user.getBookmarks", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UserGetStationList gets the list of a user's stations.
// Calls API method "user.getStationList".
func (c *Client) UserGetStationList(includeStationArtURL bool) (*responses.UserGetStationList, error) {
	if err := c.validateUserAuthToken("getting station list"); err != nil {
		return nil, err
	}
	requestData := requests.UserGetStationList{
		UserAuthToken:        c.userAuthToken,
		SyncTime:             c.GetSyncTime(),
		IncludeStationArtURL: includeStationArtURL,
	}

	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.UserGetStationList
	err = c.BlowfishCall("http://", "user.getStationList", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UserGetStationListChecksum returns the checksum of the user's station list.
// Calls API method "user.getStationListChecksum".
func (c *Client) UserGetStationListChecksum() (*responses.UserGetStationListChecksum, error) {
	if err := c.validateUserAuthToken("getting station list checksum"); err != nil {
		return nil, err
	}
	requestData := requests.UserGetStationListChecksum{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.UserGetStationListChecksum
	err = c.BlowfishCall("http://", "user.getStationListChecksum", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UserSetQuickMix selects the stations that should be in the special QuickMix station.
// Calls API method "user.setQuickMix".
func (c *Client) UserSetQuickMix(stationIDs []string) error {
	if err := c.validateUserAuthToken("setting QuickMix"); err != nil {
		return err
	}
	requestData := requests.UserSetQuickMix{
		QuickMixStationIDs: stationIDs,
		UserAuthToken:      c.userAuthToken,
		SyncTime:           c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp interface{}
	return c.BlowfishCall("https://", "user.setQuickMix", requestDataReader, &resp)
}

// UserSleepSong marks a song to not be played again for 1 month.
// Calls API method "user.sleepSong".
func (c *Client) UserSleepSong(trackToken string) error {
	if err := c.validateUserAuthToken("sleeping a song"); err != nil {
		return err
	}
	requestData := requests.UserSleepSong{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp interface{}
	return c.BlowfishCall("https://", "user.sleepSong", requestDataReader, &resp)
}
