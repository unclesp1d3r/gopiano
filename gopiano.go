/*
Package gopiano provides a thin wrapper library around the Pandora.com client API.

This client API has been reverse engineered and documentation is available at
http://pan-do-ra-api.wikia.com/wiki/Json/5.

The package provides a Client struct with a myriad of methods which interact with the
Pandora JSON API's own methods. Each method returns a struct of the parsed JSON data and an error.
All of the responses that these methods return can be found in the responses subpackage. There
is also a requests subpackage but mostly you don't need to bother with those; they get instantiated
by these client methods.
*/
package gopiano

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cellofellow/gopiano/responses"
	"golang.org/x/crypto/blowfish" //nolint:staticcheck // required by Pandora API
)

// ClientDescription describes a particular type of client to emulate.
type ClientDescription struct {
	DeviceModel string
	Username    string
	Password    string
	BaseURL     string
	EncryptKey  string
	DecryptKey  string
	Version     string
}

// AndroidClient is the data for the Android client.
var AndroidClient = ClientDescription{ //nolint:gochecknoglobals // exported by design
	DeviceModel: "android-generic",
	Username:    "android",
	Password:    "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
	BaseURL:     "tuner.pandora.com/services/json/",
	EncryptKey:  "6#26FRL$ZWD",
	DecryptKey:  "R=U!LH$O2B#",
	Version:     "5",
}

// Client represents a Pandora client.
type Client struct {
	description      ClientDescription
	http             *http.Client
	encrypter        *blowfish.Cipher
	decrypter        *blowfish.Cipher
	timeOffset       time.Duration
	partnerAuthToken string
	partnerID        string
	userAuthToken    string
	userID           string
}

// NewClient creates a new Client with specified ClientDescription.
func NewClient(d ClientDescription) (*Client, error) {
	client := new(http.Client)
	encrypter, err := blowfish.NewCipher([]byte(d.EncryptKey))
	if err != nil {
		return nil, err
	}
	decrypter, err := blowfish.NewCipher([]byte(d.DecryptKey))
	if err != nil {
		return nil, err
	}
	return &Client{
		description: d,
		http:        client,
		encrypter:   encrypter,
		decrypter:   decrypter,
	}, nil
}

// Blowfish encrypts a string in ECB mode.
// Many methods of the Pandora API take their JSON data as Blowfish encrypted data.
// The key for the encryption is provided by the ClientDescription.
func (c *Client) encrypt(data string) string {
	chunks := make([]string, 0)
	for i := 0; i < len(data); i += 8 {
		var buf [8]byte
		var crypt [8]byte
		copy(buf[:], data[i:])
		c.encrypter.Encrypt(crypt[:], buf[:])
		encoded := hex.EncodeToString(crypt[:])
		chunks = append(chunks, encoded)
	}
	return strings.Join(chunks, "")
}

// Blowfish decrypts a string in ECB mode.
// Some data returned from the Pandora API is encrypted. This decrypts it.
// The key for the decryption is provided by the ClientDescription.
func (c *Client) decrypt(data string) (string, error) {
	chunks := make([]string, 0)
	for i := 0; i < len(data); i += 16 {
		var buf [16]byte
		var decoded, decrypted [8]byte
		copy(buf[:], data[i:])
		_, err := hex.Decode(decoded[:], buf[:])
		if err != nil {
			return "", err
		}
		c.decrypter.Decrypt(decrypted[:], decoded[:])
		chunks = append(chunks, strings.Trim(string(decrypted[:]), "\x00"))
	}
	return strings.Join(chunks, ""), nil
}

// PandoraCall is the basic function to send an HTTP POST to pandora.com.
// Arguments: protocol is either "https://" or "http://", method is whatever must be in
// the "method" url argument and specifies the remote procedure to call, body is an io.Reader
// to be passed directly into http.Post, and data is to be passed to json.Unmarshal to parse
// the JSON response.
func (c *Client) PandoraCall(protocol, method string, body io.Reader, data interface{}) error {
	urlArgs := url.Values{
		"method": {method},
	}

	if c.partnerID != "" {
		urlArgs.Add("partner_id", c.partnerID)
	}
	if c.userID != "" {
		urlArgs.Add("user_id", c.userID)
	}
	if c.partnerAuthToken != "" && c.userAuthToken == "" {
		urlArgs.Add("auth_token", c.partnerAuthToken)
	} else if c.userAuthToken != "" {
		urlArgs.Add("auth_token", c.userAuthToken)
	}
	callURL := protocol + c.description.BaseURL + "?" + urlArgs.Encode()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, callURL, body)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "gopiano")
	req.Header.Add("Content-Type", "text/plain")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var errResp responses.PandoraError
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(responseBody, &errResp)
	if err != nil {
		return err
	}

	if errResp.Stat == "fail" {
		if message, ok := responses.ErrorCodeMap[errResp.Code]; ok {
			errResp.Message = message
		}
		return errResp
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return err
	}
	return nil
}

// BlowfishCall first encrypts the body before calling PandoraCall.
// Arguments are identical to PandoraCall.
func (c *Client) BlowfishCall(protocol, method string, body io.Reader, data interface{}) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	encrypted := strings.NewReader(c.encrypt(string(bodyBytes)))
	return c.PandoraCall(protocol, method, encrypted, data)
}

// GetSyncTime calculates the SyncTime for each call based on the timeOffset.
func (c *Client) GetSyncTime() int {
	return int(time.Now().Add(c.timeOffset).Unix())
}
