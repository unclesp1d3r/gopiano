# gopiano Documentation

Welcome to the gopiano documentation! This library provides a Go wrapper for Pandora's music streaming API.

## Documentation Index

| Document                              | Description                                           |
| ------------------------------------- | ----------------------------------------------------- |
| [Getting Started](GETTING_STARTED.md) | Quick start guide, installation, and common tasks     |
| [API Reference](API.md)               | Complete API documentation with all methods and types |
| [Architecture](ARCHITECTURE.md)       | System design, diagrams, and internal workings        |

## Quick Links

### Installation

```bash
go get github.com/unclesp1d3r/gopiano
```

### Minimal Example

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

    client, _ := gopiano.NewClient(gopiano.AndroidClient)
    client.AuthPartnerLogin(ctx)    // Step 1: Partner auth
    client.AuthUserLogin(ctx, "email", "password")  // Step 2: User auth

    stations, _ := client.UserGetStationList(ctx, false)
    for _, s := range stations.Result.Stations {
        fmt.Println(s.StationName)
    }
}
```

### Key Concepts

1. **Two-step authentication** - `AuthPartnerLogin()` must be called first, then `AuthUserLogin()`
2. **Context support** - All methods accept `context.Context` for cancellation/timeouts
3. **Thin wrapper** - Library mirrors the Pandora API exactly
4. **US-only** - Requires US IP address due to licensing

## API Overview

| Category           | Methods                                                                                                                                  |
| ------------------ | ---------------------------------------------------------------------------------------------------------------------------------------- |
| **Authentication** | `AuthPartnerLogin`, `AuthUserLogin`                                                                                                      |
| **User**           | `UserGetStationList`, `UserGetBookmarks`, `UserCreateUser`, `UserSetQuickMix`, `UserSleepSong`                                           |
| **Stations**       | `StationGetPlaylist`, `StationGetStation`, `StationCreateStation*`, `StationAddFeedback`, `StationRenameStation`, `StationDeleteStation` |
| **Music**          | `MusicSearch`, `ExplainTrack`, `BookmarkAdd*`                                                                                            |

## Error Handling

```go
if err != nil {
    if pe, ok := err.(responses.PandoraError); ok {
        fmt.Printf("API Error %d: %s\n", pe.Code, pe.Message)
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}
```

## Important Notes

- This wraps an **unofficial, reverse-engineered API** that may change or be deprecated
- Requires a **US IP address** due to Pandora's licensing restrictions
- The `Client` is **NOT thread-safe** for concurrent use
- For educational and personal use only - comply with Pandora's Terms of Service

## External Resources

- [Pandora API Documentation (unofficial)](https://6xq.net/pandora-apidoc/json/)
- [Pandora API REST Documentation](https://6xq.net/pandora-apidoc/rest/)
- [gopiano GitHub Repository](https://github.com/unclesp1d3r/gopiano)

## License

BSD License - See [LICENSE](../LICENSE) for details.
