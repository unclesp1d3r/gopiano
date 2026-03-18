# Pandora.com API Wrapper Client

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->

[![All Contributors](https://img.shields.io/badge/all_contributors-4-orange.svg?style=flat-square)](#contributors-)

<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![Go Report Card](https://goreportcard.com/badge/github.com/unclesp1d3r/gopiano)](https://goreportcard.com/report/github.com/unclesp1d3r/gopiano)
[![GoDoc](https://pkg.go.dev/badge/github.com/unclesp1d3r/gopiano.svg)](https://pkg.go.dev/github.com/unclesp1d3r/gopiano)
[![License](https://img.shields.io/badge/license-BSD-blue.svg)](LICENSE)
[![Security Policy](https://img.shields.io/badge/security-policy-blue.svg)](SECURITY.md)

A very thin wrapper around Pandora.com's unofficial, reverse-engineered legacy JSON API (v5).

> **⚠️ Important**: This library wraps Pandora's unofficial legacy JSON API which may be deprecated. The official Pandora API now uses OAuth2 + GraphQL. Account creation and other features may be restricted or unavailable. See [API Status and Limitations](#api-status-and-limitations) for details.

## Disclaimer

**Reference Implementation**: This library is a reference implementation for educational and research purposes, demonstrating interaction with Pandora's unofficial, reverse-engineered API. It is not production-ready and is provided "as-is" without warranty.

**Valid Credentials Required**: Users must have valid Pandora account credentials (username and password) to use this library. The library does not provide credentials or account access.

**Legal Access Rights**: Users are responsible for ensuring they have legal rights to access the Pandora API and must comply with Pandora's Terms of Service. This library is not affiliated with, endorsed by, or connected to Pandora Media, LLC or its affiliates.

**Unofficial API Warning**: This library wraps an unofficial, reverse-engineered API (not reverse-engineered by the maintainers of this library, and likely not approved by Pandora) that may be deprecated, restricted, or unavailable at any time. Users assume all risks associated with using an unofficial API, including but not limited to account suspension, service interruption, or legal consequences.

**No Warranty**: This software is provided "as-is" without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose, and noninfringement. In no event shall the authors or copyright holders be liable for any claim, damages, or other liability arising from the use of this software.

## Getting Started

### Authentication Flow

All API interactions require a two-step authentication process:

1. **Step 1**: Call `AuthPartnerLogin()` to establish partner session

   - This obtains `partnerAuthToken`, `partnerID`, and `syncTime`
   - **Required before any other API methods**

2. **Step 2**: Call either `AuthUserLogin()` for existing users OR `UserCreateUser()` for new accounts

   - This obtains `userAuthToken` and `userID`
   - Required before calling user-specific methods

Only after both steps can you call other API methods that require user authentication.

**Important**: All client methods that require `partnerAuthToken` (such as `AuthUserLogin`, `UserCreateUser`, `UserEmailPassword`) or `userAuthToken` (such as `UserGetStationList`, `StationGetPlaylist`, `MusicSearch`, and other user/station methods) now perform client-side validation. When called without the required prior authentication, these methods return standard Go `error` values (not `responses.PandoraError`) with descriptive messages. This is a consistent, library-wide behavior that prevents cryptic server-side errors by failing fast with clear guidance.

### Basic Example

```go
package main

import (
 "context"
 "log"

 "github.com/unclesp1d3r/gopiano"
)

func main() {
 // Create a client
 client, err := gopiano.NewClient(gopiano.AndroidClient)
 if err != nil {
  log.Fatal(err)
 }

 // Step 1: Partner login (REQUIRED FIRST)
 _, err = client.AuthPartnerLogin(context.Background())
 if err != nil {
  log.Fatal(err)
 }

 // Step 2: User login (for existing users)
 _, err = client.AuthUserLogin(context.Background(), "user@example.com", "password")
 if err != nil {
  log.Fatal(err)
 }

 // Now you can call other methods
 stations, err := client.UserGetStationList(context.Background(), false)
 if err != nil {
  log.Fatal(err)
 }

 log.Printf("User has %d stations", len(stations.Result.Stations))
}
```

See the [examples/](examples/) directory for complete, runnable examples.

## Creating New Users

The `UserCreateUser` function allows you to create new Pandora user accounts. This requires careful attention to prerequisites and parameter validation.

### Prerequisites

- **Must call `AuthPartnerLogin()` first** - This is now validated and will return a clear client-side error if missing. **Important**: Calling `UserCreateUser` without first calling `AuthPartnerLogin()` will now return a standard Go `error` (not a `responses.PandoraError`) with a descriptive message. This is a breaking change from previous behavior where a server-side `PandoraError` (typically error code 0 INTERNAL) would be returned.
- **US IP address required** - Due to licensing restrictions, account creation only works from US IPs
- **Valid parameters** - All parameters must meet Pandora's strict requirements

### Complete Example

```go
package main

import (
 "context"
 "log"

 "github.com/unclesp1d3r/gopiano"
)

func main() {
 client, err := gopiano.NewClient(gopiano.AndroidClient)
 if err != nil {
  log.Fatal(err)
 }

 // Step 1: Partner login (REQUIRED)
 _, err = client.AuthPartnerLogin(context.Background())
 if err != nil {
  log.Fatalf("Partner login failed: %v", err)
 }

 // Step 2: Create new user
 userResp, err := client.UserCreateUser(context.Background(),
  "user@example.com",  // username: must be valid email
  "SecurePassword123",  // password
  "male",               // gender: must be "male" or "female"
  "US",                 // countryCode: must be "US"
  90210,                // zipCode: must be valid US ZIP
  1990,                 // birthYear: must meet age requirements
  false,                // emailOptin: marketing email preference
 )
 if err != nil {
  log.Fatalf("User creation failed: %v", err)
 }

 log.Printf("User created: %s (ID: %s)",
  userResp.Result.Username,
  userResp.Result.UserID)
}
```

### Parameter Requirements

- **username**: Must be a valid email address format
- **gender**: Must be exactly `"male"` or `"female"` (case-sensitive)
- **countryCode**: Must be `"US"` (API restriction)
- **zipCode**: Must be a valid 5-digit US ZIP code
- **birthYear**: Must meet minimum age requirements (typically 13+)
- **emailOptin**: Boolean indicating marketing email preference

### Common Pitfalls

- **Missing `AuthPartnerLogin()`**: Now caught by client-side validation with a clear error message. **Note**: This returns a standard Go `error`, not a `responses.PandoraError`. The error message will be: "partner authentication token missing: must call AuthPartnerLogin() first to establish a partner session before creating a user". This is a breaking change from previous versions where a server-side `PandoraError` would be returned.
- **Invalid email format**: Username must be a valid email address
- **Wrong gender value**: Must be exactly `"male"` or `"female"` (not `"M"`, `"F"`, etc.)
- **Non-US IP**: Account creation only works from US IP addresses
- **Rate limiting**: Too many requests may trigger rate limiting
- **Username already exists**: Email address must be unique

See [examples/create_user/](examples/create_user/) for a complete example with detailed comments.

## Troubleshooting

### "Pandora Error: 0 INTERNAL"

This generic error code typically indicates one of several issues:

1. **Missing authentication**: Ensure you called `AuthPartnerLogin()` before making the API call. **Note**: Methods requiring authentication (such as `AuthUserLogin`, `UserCreateUser`, `UserGetStationList`, `StationGetPlaylist`, etc.) now perform client-side validation and return standard Go `error` values (not `responses.PandoraError`) when called without proper authentication, preventing this error from occurring due to missing tokens.
2. **Invalid parameters**: Verify all parameters meet the requirements (see above)
3. **Geographic restrictions**: Ensure you're calling from a US IP address
4. **Rate limiting**: Wait before retrying if making frequent requests
5. **API restrictions**: The legacy API may have additional restrictions

The library now provides enhanced error messages with troubleshooting guidance when error code 0 occurs.

### Authentication Issues

**Problem**: Methods fail with authentication errors

**Solution**:

- Verify you called `AuthPartnerLogin()` first
- Check that your client was properly initialized with `NewClient(gopiano.AndroidClient)`
- Ensure you're calling from a US IP address

**Note**: All methods requiring authentication (e.g., `AuthUserLogin`, `UserCreateUser`, `UserGetStationList`, `StationGetPlaylist`, `MusicSearch`) now perform client-side validation. When called without the required `partnerAuthToken` or `userAuthToken`, they return standard Go `error` values (not `responses.PandoraError`) with descriptive messages indicating which authentication method to call first. This prevents cryptic "Pandora Error: 0 INTERNAL" errors by failing fast with clear guidance.

### Parameter Validation Errors

**Problem**: `UserCreateUser` fails with parameter errors

**Solution**:

- Verify email format is valid
- Ensure gender is exactly `"male"` or `"female"`
- Check that countryCode is `"US"`
- Validate ZIP code is a valid 5-digit US ZIP
- Confirm birth year meets age requirements

### Geographic Restrictions

**Problem**: API calls fail with licensing restriction errors

**Solution**:

- The API requires a US IP address due to licensing restrictions
- Verify your IP address location
- VPN users: Ensure your VPN is set to a US location

### Rate Limiting

**Problem**: Requests start failing after multiple calls

**Solution**:

- Implement exponential backoff retry logic
- Reduce request frequency
- Wait between batches of requests

For detailed troubleshooting guidance, see [TROUBLESHOOTING.md](TROUBLESHOOTING.md).

## API Status and Limitations

### Legacy API Warning

This library wraps Pandora's **unofficial legacy JSON API (v5)** which:

- Is reverse-engineered and not officially supported by Pandora
- May be deprecated or restricted at any time
- Has limited documentation and support
- May have rate limiting or other restrictions

### Official API Migration

Pandora's official API now uses:

- **OAuth2** for authentication
- **GraphQL** for API interactions

If you need official support or access to newer features, consider migrating to the official API.

### Known Limitations

- **US-only**: Requires US IP address due to licensing restrictions
- **Account creation**: May be restricted or unavailable
- **Rate limiting**: Frequent requests may be rate-limited
- **Error code 0**: Generic "INTERNAL" error provides limited diagnostic information
- **No official support**: This is an unofficial API with no official support channel

### Recommendation

Before using this library for production applications:

1. Verify the legacy API still supports your use case
2. Test thoroughly with your specific requirements
3. Consider migrating to the official OAuth2 + GraphQL API if available
4. Implement robust error handling and retry logic

## Usage and Hacking

### Using as a Dependency

Add gopiano to your Go module:

```sh
go get github.com/unclesp1d3r/gopiano
```

Then import it in your code:

```go
import "github.com/unclesp1d3r/gopiano"
```

### Development

To contribute or hack on gopiano, clone the repository:

```sh
git clone https://github.com/unclesp1d3r/gopiano.git
cd gopiano
```

The project uses Go modules, so you can work on it from any directory. See [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

## Testing

The library includes comprehensive regression tests to prevent authentication-related errors. These tests verify:

1. **Methods requiring partner authentication** fail with descriptive errors when `AuthPartnerLogin()` hasn't been called
2. **Methods requiring user authentication** fail with descriptive errors when `AuthUserLogin()` or `UserCreateUser()` hasn't been called
3. **Error messages are standard Go errors** (not `PandoraError`) for client-side validation failures
4. **Error messages include guidance** on which authentication method to call first

These tests can be run with `just test` or `go test ./...` and are separate from integration tests (which require actual Pandora credentials and use the `//go:build integration` build tag). The regression tests ensure that misuse scenarios (calling methods without proper authentication) fail fast with clear, actionable error messages instead of making API calls that return cryptic "Pandora Error: 0 INTERNAL" errors.

For more details on authentication flow and common errors, see [TROUBLESHOOTING.md](TROUBLESHOOTING.md).

## Project Status

> **⚠️ Alpha Quality**: This is *alpha quality code* - use at your own risk. Known issues include:
>
> - Limited test coverage
> - Error handling improvements needed
> - Wraps an unofficial API that may be deprecated
> - Account creation and other features may be restricted
>
> **Important**: Before using this library, please review the [Disclaimer](#disclaimer) section above. This is a reference implementation requiring valid Pandora credentials and legal API access rights. Users are responsible for compliance with Pandora's Terms of Service.

The project is actively being improved with ongoing work on:

- Proper tests
- Proper error handling
- Enhanced validation and error messages
- Comprehensive documentation

Contributions are welcome! See [CONTRIBUTING.md](CONTRIBUTING.md) for how to help improve the project.

## Breaking Changes

### PandoraError Pointer Semantics

`PandoraError` now uses pointer semantics. If you check for Pandora API errors, update your code:

**Before (no longer works):**

```go
var pe responses.PandoraError
if errors.As(err, &pe) { ... }
```

**After:**

```go
var pe *responses.PandoraError
if errors.As(err, &pe) { ... }
```

### Default HTTP Timeout

`NewClient` now sets a 30-second HTTP timeout by default. Use `WithHTTPClient` to customize:

```go
client, err := gopiano.NewClient(gopiano.AndroidClient, gopiano.WithHTTPClient(&http.Client{
    Timeout: 60 * time.Second,
}))
```

## Documentation & Community

- **[TROUBLESHOOTING.md](TROUBLESHOOTING.md)** - Common issues and solutions
- **[examples/](examples/)** - Runnable code examples
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - How to contribute to this project
- **[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)** - Our community standards
- **[SECURITY.md](SECURITY.md)** - Security policy and vulnerability reporting
- **[GOVERNANCE.md](GOVERNANCE.md)** - Project governance and decision-making
- **[CONTRIBUTORS.md](CONTRIBUTORS.md)** - Recognition of our contributors
- **[AGENTS.md](AGENTS.md)** - AI agent instructions for project context

### Security

**Please report security vulnerabilities privately** per our [Security Policy](SECURITY.md). Do not report security issues via public GitHub issues.
