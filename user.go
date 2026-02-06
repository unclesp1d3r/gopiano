package gopiano

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/unclesp1d3r/gopiano/requests"
	"github.com/unclesp1d3r/gopiano/responses"
)

// UserCanSubscribe returns whether a user is subscribed or can subscribe
// to the premium Pandora One service.
// Calls API method "user.canSubscribe".
func (c *Client) UserCanSubscribe(ctx context.Context) (*responses.UserCanSubscribe, error) {
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
	err = c.BlowfishCall(ctx, "https://", "user.canSubscribe", requestDataReader, &resp)
	if err != nil {
		return nil, fmt.Errorf("check subscription status: %w", err)
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
	ctx context.Context,
	username, password, gender, countryCode string,
	zipCode, birthYear int,
	emailOptin bool,
) (*responses.UserCreateUser, error) {
	// Validate inputs
	if err := validateEmail(username); err != nil {
		return nil, fmt.Errorf("invalid username: %w", err)
	}
	if password == "" {
		return nil, errors.New("password is required")
	}
	if gender != "male" && gender != "female" {
		return nil, fmt.Errorf("gender must be 'male' or 'female', got: %s", gender)
	}
	if countryCode != "US" {
		return nil, fmt.Errorf("country code must be 'US', got: %s", countryCode)
	}
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
	err = c.BlowfishCall(ctx, "https://", "user.createUser", requestDataReader, &resp)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	// Set user data onto client for later use (thread-safe).
	c.mu.Lock()
	c.userAuthToken = resp.Result.UserAuthToken
	c.userID = resp.Result.UserID
	c.mu.Unlock()

	return &resp, nil
}

// UserEmailPassword resends the registration email.
// Calls API method "user.emailPassword".
func (c *Client) UserEmailPassword(ctx context.Context, username string) error {
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
	var resp any
	if err = c.BlowfishCall(ctx, "https://", "user.emailPassword", requestDataReader, &resp); err != nil {
		return fmt.Errorf("email password: %w", err)
	}
	return nil
}

// UserGetBookmarks returns the user's bookmarked artists and songs.
// Also see BookmarkAddArtistBookmark and BookmarkAddSongBookmark.
// Calls API method "user.getBookmarks".
func (c *Client) UserGetBookmarks(ctx context.Context) (*responses.UserGetBookmarks, error) {
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
	err = c.BlowfishCall(ctx, "https://", "user.getBookmarks", requestDataReader, &resp)
	if err != nil {
		return nil, fmt.Errorf("get bookmarks: %w", err)
	}
	return &resp, nil
}

// UserGetStationList gets the list of a user's stations.
// Calls API method "user.getStationList".
func (c *Client) UserGetStationList(
	ctx context.Context,
	includeStationArtURL bool,
) (*responses.UserGetStationList, error) {
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
	err = c.BlowfishCall(ctx, "https://", "user.getStationList", requestDataReader, &resp)
	if err != nil {
		return nil, fmt.Errorf("get station list: %w", err)
	}
	return &resp, nil
}

// UserGetStationListChecksum returns the checksum of the user's station list.
// Calls API method "user.getStationListChecksum".
func (c *Client) UserGetStationListChecksum(ctx context.Context) (*responses.UserGetStationListChecksum, error) {
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
	err = c.BlowfishCall(ctx, "https://", "user.getStationListChecksum", requestDataReader, &resp)
	if err != nil {
		return nil, fmt.Errorf("get station list checksum: %w", err)
	}
	return &resp, nil
}

// UserSetQuickMix selects the stations that should be in the special QuickMix station.
// Calls API method "user.setQuickMix".
func (c *Client) UserSetQuickMix(ctx context.Context, stationIDs []string) error {
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
	var resp any
	if err = c.BlowfishCall(ctx, "https://", "user.setQuickMix", requestDataReader, &resp); err != nil {
		return fmt.Errorf("set quick mix: %w", err)
	}
	return nil
}

// UserSleepSong marks a song to not be played again for 1 month.
// Calls API method "user.sleepSong".
func (c *Client) UserSleepSong(ctx context.Context, trackToken string) error {
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
	var resp any
	if err = c.BlowfishCall(ctx, "https://", "user.sleepSong", requestDataReader, &resp); err != nil {
		return fmt.Errorf("sleep song: %w", err)
	}
	return nil
}
