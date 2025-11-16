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
// Argument username must be in the form of an email address. gender must be either "male" or "female".
// countryCode must be "US".
// Calls API method "user.createUser".
func (c *Client) UserCreateUser(
	username, password, gender, countryCode string,
	zipCode, birthYear int,
	emailOptin bool,
) (*responses.UserCreateUser, error) {
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
