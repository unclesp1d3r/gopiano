package gopiano

import (
	"bytes"
	"encoding/json"

	"github.com/cellofellow/gopiano/requests"
	"github.com/cellofellow/gopiano/responses"
)

// StationAddFeedback adds feedback (thumbs up or down, or star or ban if you prefer) to a song.
// Argument trackToken is the token identifying a track. Obtained from Client.StationGetPlaylist
// Argument isPositive is a bool which if true is a "star" and if false is a "ban".
// Calls API method "station.addFeedback".
func (c *Client) StationAddFeedback(trackToken string, isPositive bool) (*responses.StationAddFeedback, error) {
	requestData := requests.StationAddFeedback{
		TrackToken:    trackToken,
		IsPositive:    isPositive,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationAddFeedback
	err = c.BlowfishCall("http://", "station.addFeedback", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationAddMusic adds an additional music seed to an existing station.
// Argument musicToken is obtained from Client.MusicSearch
// Argument stationToken is obtained from Client.UserGetStationList
// Calls API method "station.addMusic".
func (c *Client) StationAddMusic(musicToken, stationToken string) (*responses.StationAddMusic, error) {
	requestData := requests.StationAddMusic{
		MusicToken:    musicToken,
		StationToken:  stationToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationAddMusic
	err = c.BlowfishCall("http://", "station.addMusic", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationCreateStationTrack creates a new station from a specified track.
// Argument trackToken is a token of a song or artist obtained from Client.StationGetPlaylist.
// Argument musicType is either "song" or "artist" specifying the type of track being used.
// Calls API method "station.createStation".
func (c *Client) StationCreateStationTrack(trackToken, musicType string) (*responses.StationCreateStation, error) {
	requestData := requests.StationCreateStation{
		TrackToken:    trackToken,
		MusicType:     musicType,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationCreateStation
	err = c.BlowfishCall("http://", "station.createStation", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationCreateStationMusic creates a new station from a music search result.
// Argument musicToken is obtained from Client.MusicSearch.
// Calls API method "station.createStation".
func (c *Client) StationCreateStationMusic(musicToken string) (*responses.StationCreateStation, error) {
	requestData := requests.StationCreateStation{
		MusicToken:    musicToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationCreateStation
	err = c.BlowfishCall("http://", "station.createStation", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationDeleteFeedback deletes feedback (thumbs up/down) on a track's feedback ID.
// Calls API method "station.deleteFeedback".
func (c *Client) StationDeleteFeedback(feedbackID string) error {
	requestData := requests.StationDeleteFeedback{
		FeedbackID:    feedbackID,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp interface{}
	return c.BlowfishCall("http://", "station.deleteFeedback", requestDataReader, &resp)
}

// StationDeleteMusic removes seed music identified by a seedID from a station.
// Calls API method "station.deleteMusic".
func (c *Client) StationDeleteMusic(seedID string) error {
	requestData := requests.StationDeleteMusic{
		SeedID:        seedID,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp interface{}
	return c.BlowfishCall("http://", "station.deleteMusic", requestDataReader, &resp)
}

// StationDeleteStation removes a station identified by a stationToken.
// Calls API method "station.deleteStation".
func (c *Client) StationDeleteStation(stationToken string) error {
	requestData := requests.StationDeleteStation{
		StationToken:  stationToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp interface{}
	return c.BlowfishCall("http://", "station.deleteStation", requestDataReader, &resp)
}

// StationGetGenreStations retrieves a list of predefined "genre stations".
// Calls API method "station.getGenreStations".
func (c *Client) StationGetGenreStations() (*responses.StationGetGenreStations, error) {
	requestData := requests.StationGetGenreStations{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationGetGenreStations
	err = c.BlowfishCall("http://", "station.getGenreStations", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationGetPlaylist retrieves a playlist for a specified token.
// Argument stationToken is obtained from UserGetStationList.
// Note: an error response with code 0 may mean you've called getPlaylist too much.
// Calls API method "station.getPlaylist".
func (c *Client) StationGetPlaylist(stationToken string) (*responses.StationGetPlaylist, error) {
	requestData := requests.StationGetPlaylist{
		StationToken:  stationToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationGetPlaylist
	err = c.BlowfishCall("https://", "station.getPlaylist", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationGetStation retrieves station details.
// Argument stationToken is obtained from Client.UserGetStationList
// Argument includeExtendedAttributes will include music seed and feedback IDs in response.
// Calls API method "station.getStation".
func (c *Client) StationGetStation(
	stationToken string,
	includeExtendedAttributes bool,
) (*responses.StationGetStation, error) {
	requestData := requests.StationGetStation{
		StationToken:              stationToken,
		IncludeExtendedAttributes: includeExtendedAttributes,
		UserAuthToken:             c.userAuthToken,
		SyncTime:                  c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationGetStation
	err = c.BlowfishCall("http://", "station.getStation", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationShareStation shares a station with provided email addresses.
// Arguments stationID and stationToken obtained from Client.UserGetStationList
// Argument emails is a list of email addresses.
// Calls API method "station.shareStation".
func (c *Client) StationShareStation(stationID, stationToken string, emails []string) error {
	requestData := requests.StationShareStation{
		StationToken:  stationToken,
		StationID:     stationID,
		Emails:        emails,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp interface{}
	return c.BlowfishCall("http://", "station.shareStation", requestDataReader, &resp)
}

// StationRenameStation sets a new name for a station.
// Calls API method "station.renameStation".
func (c *Client) StationRenameStation(stationToken, stationName string) (*responses.StationRenameStation, error) {
	requestData := requests.StationRenameStation{
		StationToken:  stationToken,
		StationName:   stationName,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationRenameStation
	err = c.BlowfishCall("http://", "station.renameStation", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StationTransformSharedStation copies a shared station and creates a user-editable station.
// Calls API method "station.transformSharedStation".
func (c *Client) StationTransformSharedStation(stationToken string) (*responses.StationTransformSharedStation, error) {
	requestData := requests.StationTransformSharedStation{
		StationToken:  stationToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationTransformSharedStation
	err = c.BlowfishCall("http://", "station.transformSharedStation", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
