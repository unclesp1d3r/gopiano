package gopiano

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	"github.com/unclesp1d3r/gopiano/requests"
	"github.com/unclesp1d3r/gopiano/responses"
)

// AuthPartnerLogin establishes a Partner session with provided
// API username and password and receives a PartnerAuthToken, PartnerID and SyncTime
// which are stored for later calls.
// Calls API method "auth.partnerLogin".
func (c *Client) AuthPartnerLogin() (*responses.AuthPartnerLogin, error) {
	requestData := requests.AuthPartnerLogin{
		Username:    c.description.Username,
		Password:    c.description.Password,
		Version:     c.description.Version,
		DeviceModel: c.description.DeviceModel,
		IncludeURLs: true,
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.AuthPartnerLogin
	err = c.PandoraCall("https://", "auth.partnerLogin", requestDataReader, &resp)
	if err != nil {
		// TODO Handle error
		return nil, err
	}

	syncTime, err := c.decrypt(resp.Result.SyncTime)
	if err != nil {
		return nil, err
	}
	resp.Result.SyncTime = syncTime[4:14]
	i, err := strconv.ParseInt(resp.Result.SyncTime, 10, 32)
	if err != nil {
		return nil, err
	}

	// Set partner data onto client for later use.
	c.timeOffset = time.Until(time.Unix(i, 0))
	c.partnerAuthToken = resp.Result.PartnerAuthToken
	c.partnerID = resp.Result.PartnerID

	return &resp, nil
}

// AuthUserLogin logs in a username and password pair.
// Receives the UserAuthToken which is used in subsequent calls.
// You must call AuthPartnerLogin first, and then either this method or UserCreateUser
// before you proceed.
// Calls API method "auth.userLogin".
func (c *Client) AuthUserLogin(username, password string) (*responses.AuthUserLogin, error) {
	requestData := requests.AuthUserLogin{
		PartnerAuthToken: c.partnerAuthToken,
		LoginType:        "user",
		Username:         username,
		Password:         password,
		SyncTime:         c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.AuthUserLogin
	err = c.BlowfishCall("https://", "auth.userLogin", requestDataReader, &resp)
	if err != nil {
		// TODO Handle error
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.Result.UserAuthToken
	c.userID = resp.Result.UserID

	return &resp, nil
}
