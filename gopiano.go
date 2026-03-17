/*
Package gopiano provides a thin wrapper library around Pandora.com's unofficial, reverse-engineered legacy JSON API (v5).

This client API has been reverse engineered and documentation is available at
https://6xq.net/pandora-apidoc/json/ and https://6xq.net/pandora-apidoc/rest/.

The package provides a Client struct with a myriad of methods which interact with the
Pandora JSON API's own methods. Each method returns a struct of the parsed JSON data and an error.
All of the responses that these methods return can be found in the responses subpackage. There
is also a requests subpackage but mostly you don't need to bother with those; they get instantiated
by these client methods.

# Authentication Flow

All API interactions require a two-step authentication process:

1. Step 1: Call AuthPartnerLogin() to establish partner session
  - This obtains partnerAuthToken, partnerID, and syncTime
  - Required before any other API methods

2. Step 2: Call either AuthUserLogin() for existing users OR UserCreateUser() for new accounts
  - This obtains userAuthToken and userID
  - Required before calling user-specific methods

Only after both steps can you call other API methods that require user authentication.

# Quick Start

	ctx := context.Background()

	client, err := gopiano.NewClient(gopiano.AndroidClient)
	if err != nil {
		log.Fatal(err)
	}

	// Step 1: Partner login (required first)
	_, err = client.AuthPartnerLogin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: User login (for existing users)
	_, err = client.AuthUserLogin(ctx, "user@example.com", "password")
	if err != nil {
		log.Fatal(err)
	}

	// Now you can call other methods
	stations, err := client.UserGetStationList(ctx, false)

# Important Notes

- This wraps an unofficial, legacy Pandora API that may be deprecated
- US-only restrictions due to licensing (requires US IP address)
- Potential for rate limiting on frequent requests
- Error code 0 (INTERNAL) typically indicates authentication or validation issues
- Official Pandora API now uses OAuth2 + GraphQL

# Concurrency

A Client is safe for concurrent use by multiple goroutines. The Client uses
an internal mutex (sync.RWMutex) to protect authentication state (tokens, time
offset) during reads and writes.

For best performance in highly concurrent scenarios, consider creating separate
Client instances per goroutine to avoid lock contention.

# Disclaimer

This is a reference implementation for educational and research purposes. Users must have valid Pandora account credentials and are responsible for ensuring they have legal rights to access the Pandora API and comply with Pandora's Terms of Service. This software is provided "as-is" without warranty.
*/
package gopiano

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/unclesp1d3r/gopiano/responses"
	"golang.org/x/crypto/blowfish" //nolint:staticcheck // required by Pandora API
)

// Blowfish encryption constants.
const (
	// blowfishBlockSize is the size of a Blowfish block in bytes.
	blowfishBlockSize = 8
	// hexEncodedBlockSize is the size of a hex-encoded Blowfish block.
	hexEncodedBlockSize = 16
)

// emailRegex is a simple email validation pattern.
var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
) //nolint:gochecknoglobals // compiled once for efficiency

// ClientDescription describes a particular type of client to emulate.
type ClientDescription struct {
	// DeviceModel is the device identifier sent to the Pandora API (e.g., "android-generic").
	DeviceModel string
	// Username is the partner-level username used during partner authentication.
	Username string
	// Password is the partner-level password used during partner authentication.
	Password string
	// BaseURL is the base URL for the Pandora API (e.g., "tuner.pandora.com/services/json/").
	BaseURL string
	// EncryptKey is the Blowfish key used to encrypt outgoing request payloads.
	EncryptKey string
	// DecryptKey is the Blowfish key used to decrypt incoming response payloads (e.g., syncTime).
	DecryptKey string
	// Version is the Pandora API protocol version (e.g., "5").
	Version string
}

// AndroidClient is the data for the Android client.
//
// Deprecated: AndroidClient is a mutable global variable. Use DefaultAndroidClient() instead,
// which returns a fresh copy and avoids accidental mutation of shared state.
//
// SECURITY NOTE: These are partner-level credentials for the unofficial Pandora API,
// not user credentials. They are publicly documented and required for API communication.
// User credentials (email/password) are transmitted separately and securely over HTTPS.
// These keys are used for Blowfish encryption of request payloads as required by the API protocol.
var AndroidClient = ClientDescription{ //nolint:gochecknoglobals,gosec // exported by design; partner credentials are public, not user secrets
	DeviceModel: "android-generic",
	Username:    "android",
	Password:    "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
	BaseURL:     "tuner.pandora.com/services/json/",
	EncryptKey:  "6#26FRL$ZWD",
	DecryptKey:  "R=U!LH$O2B#",
	Version:     "5",
}

// DefaultAndroidClient returns a fresh copy of the Android client description.
// Use this instead of the AndroidClient global variable to avoid accidental mutation of shared state.
func DefaultAndroidClient() ClientDescription {
	return ClientDescription{ //nolint:gosec // partner credentials are public, not user secrets
		DeviceModel: "android-generic",
		Username:    "android",
		Password:    "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
		BaseURL:     "tuner.pandora.com/services/json/",
		EncryptKey:  "6#26FRL$ZWD",
		DecryptKey:  "R=U!LH$O2B#",
		Version:     "5",
	}
}

// Client represents a Pandora client.
//
// The Client is safe for concurrent use by multiple goroutines. It maintains
// authentication state protected by an internal mutex (sync.RWMutex).
// For best performance in highly concurrent scenarios, consider creating
// separate Client instances per goroutine to avoid lock contention.
type Client struct {
	// mu protects all mutable fields below
	mu sync.RWMutex

	// Immutable after construction (no mutex needed for these)
	description ClientDescription
	// http uses the default TLS configuration; no certificate pinning is performed.
	http      *http.Client
	encrypter *blowfish.Cipher
	decrypter *blowfish.Cipher

	// Mutable state (protected by mu)
	timeOffset       time.Duration
	partnerAuthToken string
	partnerID        string
	userAuthToken    string
	userID           string
}

// Option configures a Client. Use With* functions to create Options.
type Option func(*Client)

// WithHTTPClient replaces the default HTTP client entirely.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.http = hc }
}

// WithTimeout sets the timeout on the default HTTP client.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) { c.http.Timeout = d }
}

// NewClient creates a new Client with specified ClientDescription and optional configuration.
func NewClient(d ClientDescription, opts ...Option) (*Client, error) {
	client := &http.Client{
		Timeout: 30 * time.Second, //nolint:mnd // reasonable default timeout
		Transport: &http.Transport{
			TLSHandshakeTimeout:   10 * time.Second, //nolint:mnd // reasonable TLS timeout
			ResponseHeaderTimeout: 15 * time.Second, //nolint:mnd // reasonable header timeout
			IdleConnTimeout:       90 * time.Second, //nolint:mnd // standard idle timeout
			MaxIdleConns:          10,               //nolint:mnd // reasonable pool size
			MaxIdleConnsPerHost:   10,               //nolint:mnd // single host API
		},
	}
	encrypter, err := blowfish.NewCipher([]byte(d.EncryptKey))
	if err != nil {
		return nil, err
	}
	decrypter, err := blowfish.NewCipher([]byte(d.DecryptKey))
	if err != nil {
		return nil, err
	}
	c := &Client{
		description: d,
		http:        client,
		encrypter:   encrypter,
		decrypter:   decrypter,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

// encrypt encrypts a string using Blowfish in ECB mode.
// Many methods of the Pandora API take their JSON data as Blowfish encrypted data.
// The key for the encryption is provided by the ClientDescription.
//
// SECURITY: ECB mode is weak because identical plaintext blocks produce identical
// ciphertext blocks, enabling pattern analysis. This is a Pandora API protocol
// requirement, not a design choice.
func (c *Client) encrypt(data string) string {
	if data == "" {
		return ""
	}

	// Pre-calculate capacity: each 8-byte block becomes 16 hex chars
	numBlocks := (len(data) + blowfishBlockSize - 1) / blowfishBlockSize
	var result strings.Builder
	result.Grow(numBlocks * hexEncodedBlockSize)

	for i := 0; i < len(data); i += blowfishBlockSize {
		var buf [blowfishBlockSize]byte
		var crypt [blowfishBlockSize]byte
		copy(buf[:], data[i:])
		c.encrypter.Encrypt(crypt[:], buf[:])
		result.WriteString(hex.EncodeToString(crypt[:]))
	}
	return result.String()
}

// decrypt decrypts a string using Blowfish in ECB mode.
// Some data returned from the Pandora API is encrypted. This decrypts it.
// The key for the decryption is provided by the ClientDescription.
//
// SECURITY: ECB mode is weak because identical plaintext blocks produce identical
// ciphertext blocks, enabling pattern analysis. This is a Pandora API protocol
// requirement, not a design choice.
func (c *Client) decrypt(data string) (string, error) {
	if data == "" {
		return "", nil
	}

	// Pre-calculate capacity: each 16 hex chars becomes up to 8 bytes of text
	numBlocks := len(data) / hexEncodedBlockSize
	var result strings.Builder
	result.Grow(numBlocks * blowfishBlockSize)

	for i := 0; i < len(data); i += hexEncodedBlockSize {
		var buf [hexEncodedBlockSize]byte
		var decoded, decrypted [blowfishBlockSize]byte
		copy(buf[:], data[i:])
		_, err := hex.Decode(decoded[:], buf[:])
		if err != nil {
			return "", err
		}
		c.decrypter.Decrypt(decrypted[:], decoded[:])
		result.WriteString(strings.TrimRight(string(decrypted[:]), "\x00"))
	}
	return result.String(), nil
}

// PandoraCall is the basic function to send an HTTP POST to pandora.com.
// Arguments: ctx is the context for request cancellation and deadlines,
// protocol is either "https://" or "http://", method is whatever must be in
// the "method" url argument and specifies the remote procedure to call, body is an io.Reader
// to be passed directly into http.Post, and data is to be passed to json.Unmarshal to parse
// the JSON response.
//
// SECURITY: Using "http://" transmits data in plaintext, including auth tokens in URL
// query parameters. Always use "https://" to protect credentials in transit.
func (c *Client) PandoraCall(ctx context.Context, protocol, method string, body io.Reader, data any) error {
	// Check for context cancellation before starting
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Auth tokens are passed as URL query parameters because the Pandora API
	// protocol requires them there (the encrypted request body carries payload
	// data, not authentication). Always use "https://" so tokens are protected.
	urlArgs := url.Values{
		"method": {method},
	}

	// Thread-safe read of authentication state
	c.mu.RLock()
	partnerID := c.partnerID
	userID := c.userID
	partnerAuthToken := c.partnerAuthToken
	userAuthToken := c.userAuthToken
	c.mu.RUnlock()

	if partnerID != "" {
		urlArgs.Add("partner_id", partnerID)
	}
	if userID != "" {
		urlArgs.Add("user_id", userID)
	}
	if partnerAuthToken != "" && userAuthToken == "" {
		urlArgs.Add("auth_token", partnerAuthToken)
	} else if userAuthToken != "" {
		urlArgs.Add("auth_token", userAuthToken)
	}
	callURL := protocol + c.description.BaseURL + "?" + urlArgs.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, callURL, body)
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
	const maxResponseSize = 1 << 20 // 1 MB
	responseBody, err := io.ReadAll(io.LimitReader(resp.Body, maxResponseSize))
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
		// Provide additional troubleshooting guidance for error code 0 (INTERNAL)
		if errResp.Code == 0 {
			guidance := responses.GetErrorGuidance(errResp.Code)
			if guidance != "" {
				errResp.Message = errResp.Message + ". " + guidance
			}
		}
		return &errResp
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		return err
	}
	return nil
}

// BlowfishCall first encrypts the body before calling PandoraCall.
// Arguments are identical to PandoraCall.
//
// SECURITY: Using "http://" transmits data in plaintext, including auth tokens in URL
// query parameters. Always use "https://" to protect credentials in transit.
func (c *Client) BlowfishCall(ctx context.Context, protocol, method string, body io.Reader, data any) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	encrypted := strings.NewReader(c.encrypt(string(bodyBytes)))
	return c.PandoraCall(ctx, protocol, method, encrypted, data)
}

// GetSyncTime calculates the SyncTime for each call based on the timeOffset.
func (c *Client) GetSyncTime() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return int(time.Now().Add(c.timeOffset).Unix())
}

// validateEmail checks if the provided email address has a valid format.
func validateEmail(email string) error {
	if email == "" {
		return errors.New("email address is required")
	}
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// getPartnerAuthToken returns the partner auth token under the read lock, or an error if missing.
// This is used by methods that require partner authentication (e.g., AuthUserLogin, UserCreateUser).
func (c *Client) getPartnerAuthToken(operation string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.partnerAuthToken == "" {
		return "", fmt.Errorf(
			"partner authentication token missing: must call AuthPartnerLogin() first to establish a partner session before %s",
			operation,
		)
	}
	return c.partnerAuthToken, nil
}

// getUserAuthToken returns the user auth token under the read lock, or an error if missing.
// This is used by methods that require user authentication (e.g., UserGetStationList, StationGetPlaylist).
func (c *Client) getUserAuthToken(operation string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.userAuthToken == "" {
		return "", fmt.Errorf(
			"user authentication token missing: must call AuthUserLogin() or UserCreateUser() first to establish a user session before %s",
			operation,
		)
	}
	return c.userAuthToken, nil
}
