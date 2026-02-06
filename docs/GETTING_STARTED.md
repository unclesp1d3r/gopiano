# Getting Started with gopiano

This guide will help you get started with the gopiano library for interacting with Pandora's music streaming API.

## Prerequisites

- Go 1.24 or later
- A valid Pandora account (free or premium)
- A US IP address (required by Pandora's licensing)

## Installation

```bash
go get github.com/unclesp1d3r/gopiano
```

## Quick Start

Here's a minimal example to authenticate and list your stations:

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

    // Create a new client
    client, err := gopiano.NewClient(gopiano.AndroidClient)
    if err != nil {
        log.Fatal(err)
    }

    // Step 1: Partner login (REQUIRED FIRST)
    _, err = client.AuthPartnerLogin(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Step 2: User login
    _, err = client.AuthUserLogin(ctx, "your-email@example.com", "your-password")
    if err != nil {
        log.Fatal(err)
    }

    // Now you can use the API!
    stations, err := client.UserGetStationList(ctx, false)
    if err != nil {
        log.Fatal(err)
    }

    for _, station := range stations.Result.Stations {
        fmt.Printf("Station: %s\n", station.StationName)
    }
}
```

## Understanding the Authentication Flow

gopiano requires a **two-step authentication process**:

```text
┌─────────────────────────────────────────────────────────────┐
│                    Authentication Flow                       │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│   1. AuthPartnerLogin()  ─────►  Partner Session            │
│      • Must be called FIRST                                 │
│      • Establishes partner credentials                      │
│      • Sets time synchronization                            │
│                                                             │
│   2. AuthUserLogin()  ────────►  User Session               │
│      OR                                                     │
│      UserCreateUser()                                       │
│      • Requires partner session                             │
│      • Establishes user credentials                         │
│                                                             │
│   3. API Methods  ────────────►  Full Access                │
│      • UserGetStationList()                                 │
│      • StationGetPlaylist()                                 │
│      • MusicSearch()                                        │
│      • etc.                                                 │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

## Common Tasks

### Getting a Playlist

```go
// After authentication...

// Get your stations
stations, err := client.UserGetStationList(ctx, false)
if err != nil {
    log.Fatal(err)
}

if len(stations.Result.Stations) > 0 {
    // Get playlist for the first station
    station := stations.Result.Stations[0]
    playlist, err := client.StationGetPlaylist(ctx, station.StationToken)
    if err != nil {
        log.Fatal(err)
    }

    for _, track := range playlist.Result.Items {
        fmt.Printf("%s - %s\n", track.ArtistName, track.SongName)
        fmt.Printf("  Album: %s\n", track.AlbumName)

        // Access audio URLs
        for quality, audio := range track.AudioURLMap {
            fmt.Printf("  Audio (%s): %s\n", quality, audio.AudioURL)
        }
    }
}
```

### Searching for Music

```go
results, err := client.MusicSearch(ctx, "The Beatles")
if err != nil {
    log.Fatal(err)
}

fmt.Println("Artists:")
for _, artist := range results.Result.Artists {
    fmt.Printf("  %s (score: %d)\n", artist.ArtistName, artist.Score)
}

fmt.Println("Songs:")
for _, song := range results.Result.Songs {
    fmt.Printf("  %s - %s (score: %d)\n", song.ArtistName, song.SongName, song.Score)
}
```

### Creating a Station

```go
// Search for an artist
results, err := client.MusicSearch(ctx, "Radiohead")
if err != nil {
    log.Fatal(err)
}

// Create station from first artist result
if len(results.Result.Artists) > 0 {
    artist := results.Result.Artists[0]
    station, err := client.StationCreateStationMusic(ctx, artist.MusicToken)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created station: %s\n", station.Result.StationName)
}
```

### Rating Songs (Thumbs Up/Down)

```go
// Get a playlist first
playlist, err := client.StationGetPlaylist(ctx, stationToken)
if err != nil {
    log.Fatal(err)
}

if len(playlist.Result.Items) > 0 {
    track := playlist.Result.Items[0]

    // Thumbs up (positive feedback)
    _, err = client.StationAddFeedback(ctx, track.TrackToken, true)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Liked: %s\n", track.SongName)

    // Or thumbs down (negative feedback)
    // _, err = client.StationAddFeedback(ctx, track.TrackToken, false)
}
```

### Managing Bookmarks

```go
// Bookmark a song
_, err = client.BookmarkAddSongBookmark(ctx, trackToken)
if err != nil {
    log.Fatal(err)
}

// Bookmark an artist
_, err = client.BookmarkAddArtistBookmark(ctx, trackToken)
if err != nil {
    log.Fatal(err)
}

// Get all bookmarks
bookmarks, err := client.UserGetBookmarks(ctx)
if err != nil {
    log.Fatal(err)
}

fmt.Println("Bookmarked Artists:")
for _, artist := range bookmarks.Result.Artists {
    fmt.Printf("  %s\n", artist.ArtistName)
}

fmt.Println("Bookmarked Songs:")
for _, song := range bookmarks.Result.Songs {
    fmt.Printf("  %s - %s\n", song.ArtistName, song.SongName)
}
```

## Using Context for Timeouts

All API methods accept a `context.Context`, allowing you to set timeouts:

```go
import (
    "context"
    "time"
)

// Create a context with a 10-second timeout
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

// The API call will be cancelled if it takes more than 10 seconds
stations, err := client.UserGetStationList(ctx, false)
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        log.Println("Request timed out")
    } else {
        log.Fatal(err)
    }
}
```

## Error Handling

gopiano provides detailed error information:

```go
import (
    "errors"

    "github.com/unclesp1d3r/gopiano/responses"
)

stations, err := client.UserGetStationList(ctx, false)
if err != nil {
    // Check if it's a Pandora API error (use errors.As for wrapped errors)
    var pandoraErr responses.PandoraError
    if errors.As(err, &pandoraErr) {
        fmt.Printf("Pandora Error Code: %d\n", pandoraErr.Code)
        fmt.Printf("Message: %s\n", pandoraErr.Message)

        // Get guidance for error code 0
        if pandoraErr.Code == 0 {
            guidance := responses.GetErrorGuidance(0)
            fmt.Printf("Guidance: %s\n", guidance)
        }
    } else {
        // Client-side validation error or network error
        fmt.Printf("Error: %v\n", err)
    }
}
```

### Common Error Codes

| Code | Meaning               | Solution                       |
| ---- | --------------------- | ------------------------------ |
| 0    | Internal/Generic      | Check auth, params, IP address |
| 1001 | Invalid auth token    | Re-authenticate                |
| 1002 | Invalid partner login | Check client configuration     |
| 1005 | Max stations reached  | Delete a station first         |
| 1006 | Station doesn't exist | Verify station token           |
| 1011 | Invalid username      | Use valid email format         |
| 1012 | Invalid password      | Check password                 |
| 1039 | Playlist exceeded     | Wait before requesting again   |

## Important Limitations

### Geographic Restrictions

The Pandora API is only available from US IP addresses due to music licensing restrictions. If you're outside the US, you'll need a VPN.

### Rate Limiting

The API may rate-limit frequent requests, especially for `StationGetPlaylist`. If you receive error code 0 or 1039, wait before retrying.

### Concurrency

The `Client` is **NOT thread-safe**. For concurrent use:

```go
// Option 1: Create separate clients per goroutine
go func() {
    client, err := gopiano.NewClient(gopiano.AndroidClient)
    if err != nil {
        log.Printf("Failed to create client: %v", err)
        return
    }
    if _, err := client.AuthPartnerLogin(ctx); err != nil {
        log.Printf("Partner login failed: %v", err)
        return
    }
    if _, err := client.AuthUserLogin(ctx, email, password); err != nil {
        log.Printf("User login failed: %v", err)
        return
    }
    // Use client...
}()

// Option 2: Protect with mutex
var mu sync.Mutex
mu.Lock()
stations, err := client.UserGetStationList(ctx, false)
mu.Unlock()
if err != nil {
    log.Fatal(err)
}
```

### API Status

This library wraps Pandora's **unofficial, legacy JSON API (v5)**. It may be deprecated at any time. Use for educational and personal purposes only.

## Next Steps

- Read the [API Reference](API.md) for complete method documentation
- See the [Architecture Guide](ARCHITECTURE.md) to understand how the library works
- Check out the [examples](../examples/) directory for more code samples

## Troubleshooting

### "partner authentication token missing"

You forgot to call `AuthPartnerLogin()` first. This must always be the first API call.

### "user authentication token missing"

You forgot to call `AuthUserLogin()` (or `UserCreateUser()`) after `AuthPartnerLogin()`.

### Error code 0 (INTERNAL)

This generic error can mean:

- Authentication not completed properly
- Invalid parameters
- Not calling from US IP address
- Rate limiting

### "dial tcp: connection refused"

Network issue - check your internet connection and firewall settings.

### Empty playlist

Some stations may return empty playlists if rate-limited. Wait and try again.
