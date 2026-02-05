# gopiano API Reference

Complete API reference for the gopiano library - a Go wrapper for Pandora's JSON API.

## Table of Contents

- [Client](#client)
- [Authentication](#authentication)
- [User Methods](#user-methods)
- [Station Methods](#station-methods)
- [Music & Bookmarks](#music--bookmarks)
- [Types](#types)
- [Error Handling](#error-handling)

---

## Client

### NewClient

Creates a new Pandora API client.

```go
func NewClient(d ClientDescription) (*Client, error)
```

**Parameters:**

| Name | Type                | Description                                                                   |
| ---- | ------------------- | ----------------------------------------------------------------------------- |
| `d`  | `ClientDescription` | Client configuration including device model, credentials, and encryption keys |

**Returns:**

| Type      | Description                           |
| --------- | ------------------------------------- |
| `*Client` | Configured client ready for API calls |
| `error`   | Error if cipher initialization fails  |

**Example:**

```go
client, err := gopiano.NewClient(gopiano.AndroidClient)
if err != nil {
    log.Fatal(err)
}
```

### ClientDescription

Configuration structure for client initialization.

```go
type ClientDescription struct {
    DeviceModel string  // Device identifier (e.g., "android-generic")
    Username    string  // Partner API username
    Password    string  // Partner API password
    BaseURL     string  // API base URL
    EncryptKey  string  // Blowfish encryption key
    DecryptKey  string  // Blowfish decryption key
    Version     string  // API version
}
```

### AndroidClient

Pre-configured client description for Android device emulation.

```go
var AndroidClient = ClientDescription{
    DeviceModel: "android-generic",
    Username:    "android",
    Password:    "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
    BaseURL:     "tuner.pandora.com/services/json/",
    EncryptKey:  "6#26FRL$ZWD",
    DecryptKey:  "R=U!LH$O2B#",
    Version:     "5",
}
```

---

## Authentication

### AuthPartnerLogin

Establishes a partner session with the Pandora API. **This must be called first before any other API methods.**

```go
func (c *Client) AuthPartnerLogin(ctx context.Context) (*responses.AuthPartnerLogin, error)
```

**Parameters:**

| Name  | Type              | Description                           |
| ----- | ----------------- | ------------------------------------- |
| `ctx` | `context.Context` | Context for cancellation and timeouts |

**Returns:**

| Type                          | Description                     |
| ----------------------------- | ------------------------------- |
| `*responses.AuthPartnerLogin` | Partner authentication response |
| `error`                       | Error if authentication fails   |

**Side Effects:**

- Sets `partnerAuthToken` on the client
- Sets `partnerID` on the client
- Sets `timeOffset` for sync time calculations

**Example:**

```go
ctx := context.Background()
resp, err := client.AuthPartnerLogin(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Partner ID: %s\n", resp.Result.PartnerID)
```

### AuthUserLogin

Authenticates an existing Pandora user.

```go
func (c *Client) AuthUserLogin(ctx context.Context, username, password string) (*responses.AuthUserLogin, error)
```

**Prerequisites:** `AuthPartnerLogin()` must be called first.

**Parameters:**

| Name       | Type              | Description                           |
| ---------- | ----------------- | ------------------------------------- |
| `ctx`      | `context.Context` | Context for cancellation and timeouts |
| `username` | `string`          | User's email address                  |
| `password` | `string`          | User's password                       |

**Returns:**

| Type                       | Description                   |
| -------------------------- | ----------------------------- |
| `*responses.AuthUserLogin` | User authentication response  |
| `error`                    | Error if authentication fails |

**Side Effects:**

- Sets `userAuthToken` on the client
- Sets `userID` on the client

**Example:**

```go
resp, err := client.AuthUserLogin(ctx, "user@example.com", "password")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("User ID: %s\n", resp.Result.UserID)
fmt.Printf("Can Listen: %v\n", resp.Result.CanListen)
```

---

## User Methods

### UserCreateUser

Creates a new Pandora user account.

```go
func (c *Client) UserCreateUser(
    ctx context.Context,
    username, password, gender, countryCode string,
    zipCode, birthYear int,
    emailOptin bool,
) (*responses.UserCreateUser, error)
```

**Prerequisites:** `AuthPartnerLogin()` must be called first.

**Parameters:**

| Name          | Type              | Description              | Constraints                |
| ------------- | ----------------- | ------------------------ | -------------------------- |
| `ctx`         | `context.Context` | Context for cancellation | -                          |
| `username`    | `string`          | Email address            | Must be valid email        |
| `password`    | `string`          | Account password         | -                          |
| `gender`      | `string`          | User's gender            | Must be "male" or "female" |
| `countryCode` | `string`          | Country code             | Must be "US"               |
| `zipCode`     | `int`             | US ZIP code              | Valid 5-digit ZIP          |
| `birthYear`   | `int`             | Year of birth            | Must meet age requirements |
| `emailOptin`  | `bool`            | Marketing email opt-in   | -                          |

**Side Effects:**

- Sets `userAuthToken` on the client
- Sets `userID` on the client

### UserCanSubscribe

Checks if a user can subscribe to Pandora One premium service.

```go
func (c *Client) UserCanSubscribe(ctx context.Context) (*responses.UserCanSubscribe, error)
```

**Prerequisites:** User must be authenticated.

### UserGetBookmarks

Returns the user's bookmarked artists and songs.

```go
func (c *Client) UserGetBookmarks(ctx context.Context) (*responses.UserGetBookmarks, error)
```

**Prerequisites:** User must be authenticated.

### UserGetStationList

Gets the list of a user's stations.

```go
func (c *Client) UserGetStationList(ctx context.Context, includeStationArtURL bool) (*responses.UserGetStationList, error)
```

**Prerequisites:** User must be authenticated.

**Parameters:**

| Name                   | Type              | Description                             |
| ---------------------- | ----------------- | --------------------------------------- |
| `ctx`                  | `context.Context` | Context for cancellation                |
| `includeStationArtURL` | `bool`            | Whether to include station artwork URLs |

**Example:**

```go
stations, err := client.UserGetStationList(ctx, true)
if err != nil {
    log.Fatal(err)
}
for _, station := range stations.Result.Stations {
    fmt.Printf("Station: %s (%s)\n", station.StationName, station.StationToken)
}
```

### UserGetStationListChecksum

Returns the checksum of the user's station list (useful for caching).

```go
func (c *Client) UserGetStationListChecksum(ctx context.Context) (*responses.UserGetStationListChecksum, error)
```

### UserSetQuickMix

Sets which stations should be included in the QuickMix station.

```go
func (c *Client) UserSetQuickMix(ctx context.Context, stationIDs []string) error
```

**Parameters:**

| Name         | Type              | Description                    |
| ------------ | ----------------- | ------------------------------ |
| `ctx`        | `context.Context` | Context for cancellation       |
| `stationIDs` | `[]string`        | List of station IDs to include |

### UserSleepSong

Marks a song to not be played again for 1 month.

```go
func (c *Client) UserSleepSong(ctx context.Context, trackToken string) error
```

### UserEmailPassword

Resends the registration email to a user.

```go
func (c *Client) UserEmailPassword(ctx context.Context, username string) error
```

---

## Station Methods

### StationGetPlaylist

Retrieves a playlist for a specified station.

```go
func (c *Client) StationGetPlaylist(ctx context.Context, stationToken string) (*responses.StationGetPlaylist, error)
```

**Parameters:**

| Name           | Type              | Description                             |
| -------------- | ----------------- | --------------------------------------- |
| `ctx`          | `context.Context` | Context for cancellation                |
| `stationToken` | `string`          | Station token from `UserGetStationList` |

**Note:** Error code 0 may indicate rate limiting if called too frequently.

**Example:**

```go
playlist, err := client.StationGetPlaylist(ctx, stationToken)
if err != nil {
    log.Fatal(err)
}
for _, item := range playlist.Result.Items {
    fmt.Printf("Track: %s - %s\n", item.ArtistName, item.SongName)
}
```

### StationGetStation

Retrieves detailed station information.

```go
func (c *Client) StationGetStation(ctx context.Context, stationToken string, includeExtendedAttributes bool) (*responses.StationGetStation, error)
```

**Parameters:**

| Name                        | Type              | Description                         |
| --------------------------- | ----------------- | ----------------------------------- |
| `ctx`                       | `context.Context` | Context for cancellation            |
| `stationToken`              | `string`          | Station token                       |
| `includeExtendedAttributes` | `bool`            | Include music seed and feedback IDs |

### StationGetGenreStations

Retrieves a list of predefined genre stations.

```go
func (c *Client) StationGetGenreStations(ctx context.Context) (*responses.StationGetGenreStations, error)
```

### StationCreateStationTrack

Creates a new station from a track.

```go
func (c *Client) StationCreateStationTrack(ctx context.Context, trackToken, musicType string) (*responses.StationCreateStation, error)
```

**Parameters:**

| Name         | Type              | Description               |
| ------------ | ----------------- | ------------------------- |
| `ctx`        | `context.Context` | Context for cancellation  |
| `trackToken` | `string`          | Track token from playlist |
| `musicType`  | `string`          | Either "song" or "artist" |

### StationCreateStationMusic

Creates a new station from a music search result.

```go
func (c *Client) StationCreateStationMusic(ctx context.Context, musicToken string) (*responses.StationCreateStation, error)
```

### StationAddMusic

Adds an additional music seed to an existing station.

```go
func (c *Client) StationAddMusic(ctx context.Context, musicToken, stationToken string) (*responses.StationAddMusic, error)
```

### StationDeleteMusic

Removes a seed from a station.

```go
func (c *Client) StationDeleteMusic(ctx context.Context, seedID string) error
```

### StationAddFeedback

Adds feedback (thumbs up/down) to a song.

```go
func (c *Client) StationAddFeedback(ctx context.Context, trackToken string, isPositive bool) (*responses.StationAddFeedback, error)
```

**Parameters:**

| Name         | Type              | Description                                        |
| ------------ | ----------------- | -------------------------------------------------- |
| `ctx`        | `context.Context` | Context for cancellation                           |
| `trackToken` | `string`          | Track token from playlist                          |
| `isPositive` | `bool`            | `true` = thumbs up/star, `false` = thumbs down/ban |

### StationDeleteFeedback

Removes feedback from a track.

```go
func (c *Client) StationDeleteFeedback(ctx context.Context, feedbackID string) error
```

### StationDeleteStation

Removes a station.

```go
func (c *Client) StationDeleteStation(ctx context.Context, stationToken string) error
```

### StationRenameStation

Renames a station.

```go
func (c *Client) StationRenameStation(ctx context.Context, stationToken, stationName string) (*responses.StationRenameStation, error)
```

### StationShareStation

Shares a station with email recipients.

```go
func (c *Client) StationShareStation(ctx context.Context, stationID, stationToken string, emails []string) error
```

### StationTransformSharedStation

Copies a shared station to create a user-editable version.

```go
func (c *Client) StationTransformSharedStation(ctx context.Context, stationToken string) (*responses.StationTransformSharedStation, error)
```

---

## Music & Bookmarks

### MusicSearch

Searches for music to create stations or add seeds.

```go
func (c *Client) MusicSearch(ctx context.Context, searchText string) (*responses.MusicSearch, error)
```

**Example:**

```go
results, err := client.MusicSearch(ctx, "The Beatles")
if err != nil {
    log.Fatal(err)
}
for _, artist := range results.Result.Artists {
    fmt.Printf("Artist: %s (token: %s)\n", artist.ArtistName, artist.MusicToken)
}
```

### ExplainTrack

Retrieves Music Genome Project attributes for a song.

```go
func (c *Client) ExplainTrack(ctx context.Context, trackToken string) (*responses.ExplainTrack, error)
```

### BookmarkAddArtistBookmark

Bookmarks an artist.

```go
func (c *Client) BookmarkAddArtistBookmark(ctx context.Context, trackToken string) (*responses.BookmarkAddArtistBookmark, error)
```

### BookmarkAddSongBookmark

Bookmarks a song.

```go
func (c *Client) BookmarkAddSongBookmark(ctx context.Context, trackToken string) (*responses.BookmarkAddSongBookmark, error)
```

---

## Types

### Response Types

All response types are in the `github.com/unclesp1d3r/gopiano/responses` package.

#### Station

```go
type Station struct {
    StationID          string       `json:"stationId"`
    StationToken       string       `json:"stationToken"`
    StationName        string       `json:"stationName"`
    StationDetailURL   string       `json:"stationDetailUrl"`
    ArtURL             string       `json:"artUrl"`
    DateCreated        DateResponse `json:"dateCreated"`
    AllowAddMusic      bool         `json:"allowAddMusic"`
    AllowDelete        bool         `json:"allowDelete"`
    AllowRename        bool         `json:"allowRename"`
    IsShared           bool         `json:"isShared"`
    IsQuickMix         bool         `json:"isQuickMix"`
    RequiresCleanAds   bool         `json:"requiresCleanAds"`
    SuppressVideoAds   bool         `json:"suppressVideoAds"`
    Genre              []string     `json:"genre"`
    QuickMixStationIDs []string     `json:"quickMixStationIds"`
    Music              struct { ... }
    Feedback           struct { ... }
}
```

#### StationList

Implements `sort.Interface` for sorting stations by name.

```go
type StationList []Station

func (s StationList) Len() int
func (s StationList) Swap(i, j int)
func (s StationList) Less(i, j int) bool
```

#### DateResponse

Custom date format from Pandora API with conversion method.

```go
type DateResponse struct {
    Nanos, Seconds, Year, Month, Hours, Time, Date, Minutes, Day, TimezoneOffset int
}

func (d DateResponse) GetDate() time.Time
```

---

## Error Handling

### PandoraError

API errors are returned as `responses.PandoraError`:

```go
type PandoraError struct {
    Stat    string `json:"stat"`    // "fail"
    Code    int    `json:"code"`    // Error code
    Message string `json:"message"` // Error message
}

func (e PandoraError) Error() string
```

### Error Codes

| Code | Name                     | Description                            |
| ---- | ------------------------ | -------------------------------------- |
| 0    | INTERNAL                 | Generic error (check auth, params, IP) |
| 1    | MAINTENANCE_MODE         | API is under maintenance               |
| 6    | SECURE_PROTOCOL_REQUIRED | HTTPS required                         |
| 9    | PARAMETER_MISSING        | Required parameter missing             |
| 10   | PARAMETER_VALUE_INVALID  | Invalid parameter value                |
| 12   | LICENSING_RESTRICTIONS   | Geographic restriction                 |
| 1001 | INVALID_AUTH_TOKEN       | Authentication token invalid           |
| 1002 | INVALID_PARTNER_LOGIN    | Partner credentials invalid            |
| 1005 | MAX_STATIONS_REACHED     | Station limit reached                  |
| 1006 | STATION_DOES_NOT_EXIST   | Station not found                      |
| 1011 | INVALID_USERNAME         | Invalid email format                   |
| 1012 | INVALID_PASSWORD         | Invalid password                       |
| 1013 | USERNAME_ALREADY_EXISTS  | Email already registered               |
| 1039 | PLAYLIST_EXCEEDED        | Rate limit on playlist requests        |

### Client-Side Validation

Methods validate authentication state before making API calls:

```go
// Returns standard Go error if not authenticated
resp, err := client.UserGetStationList(ctx, false)
if err != nil {
    // Could be validation error or API error
    if pe, ok := err.(responses.PandoraError); ok {
        // API error with code
        fmt.Printf("API Error %d: %s\n", pe.Code, pe.Message)
    } else {
        // Validation error (e.g., missing auth)
        fmt.Printf("Error: %v\n", err)
    }
}
```

### Error Guidance

For error code 0 (INTERNAL), use `GetErrorGuidance()`:

```go
guidance := responses.GetErrorGuidance(0)
// Returns: "Troubleshooting: Check that authentication prerequisites are met..."
```

---

## Complete Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/unclesp1d3r/gopiano"
)

func main() {
    ctx := context.Background()

    // Create client
    client, err := gopiano.NewClient(gopiano.AndroidClient)
    if err != nil {
        log.Fatal(err)
    }

    // Step 1: Partner login (required first)
    _, err = client.AuthPartnerLogin(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Step 2: User login
    _, err = client.AuthUserLogin(ctx, "user@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }

    // Get stations
    stations, err := client.UserGetStationList(ctx, false)
    if err != nil {
        log.Fatal(err)
    }

    // Get playlist for first station
    if len(stations.Result.Stations) > 0 {
        station := stations.Result.Stations[0]
        playlist, err := client.StationGetPlaylist(ctx, station.StationToken)
        if err != nil {
            log.Fatal(err)
        }

        for _, track := range playlist.Result.Items {
            fmt.Printf("%s - %s\n", track.ArtistName, track.SongName)
        }
    }
}
```
