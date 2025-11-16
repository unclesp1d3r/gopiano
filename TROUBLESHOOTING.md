# Troubleshooting Guide

This guide covers common issues and their solutions when using the gopiano library.

## Table of Contents

- [Pandora Error: 0 INTERNAL](#pandora-error-0-internal)
- [Authentication Issues](#authentication-issues)
- [Parameter Validation Errors](#parameter-validation-errors)
- [Geographic Restrictions](#geographic-restrictions)
- [Rate Limiting](#rate-limiting)
- [API Deprecation Concerns](#api-deprecation-concerns)

## Pandora Error: 0 INTERNAL

### Most Common Causes

Error code 0 (INTERNAL) is a generic error that can indicate several issues:

1. **Missing `AuthPartnerLogin()` call** (most common)

   - The library now validates this client-side and provides a clear error message
   - **Important**: This validation returns a standard Go `error` (not a `responses.PandoraError`) with the message: "partner authentication token missing: must call AuthPartnerLogin() first to establish a partner session before creating a user"
   - This is a breaking change from previous versions where a server-side `PandoraError` (typically error code 0 INTERNAL) would be returned
   - Solution: Always call `AuthPartnerLogin()` before any other API methods

2. **Invalid parameters**

   - Email format, gender value, ZIP code, birth year, etc.
   - Solution: Verify all parameters meet the requirements (see below)

3. **Rate limiting**

   - Too many requests in a short time period
   - Solution: Implement exponential backoff and reduce request frequency

4. **Geographic restrictions**

   - Not calling from a US IP address
   - Solution: Ensure your IP address is in the US

5. **API restrictions**

   - The legacy API may have additional undocumented restrictions
   - Solution: Verify the API still supports your use case

### Step-by-Step Debugging Checklist

1. ✅ **Verify authentication flow**:

   ```go
   client, _ := gopiano.NewClient(gopiano.AndroidClient)
   _, err := client.AuthPartnerLogin()
   if err != nil {
       log.Fatal("Partner login failed:", err)
   }
   ```

2. ✅ **Check parameter validity**:

   - Username is a valid email address
   - Gender is exactly "male" or "female"
   - Country code is "US"
   - ZIP code is a valid 5-digit US ZIP
   - Birth year meets age requirements

3. ✅ **Verify IP address location**:

   - Use a service like `https://ipinfo.io` to check your IP location
   - Ensure you're using a US IP address

4. ✅ **Check for rate limiting**:

   - Wait 30-60 seconds between requests
   - Implement exponential backoff retry logic

5. ✅ **Review error message**:

   - The library now provides enhanced error messages with troubleshooting tips
   - Look for specific guidance in the error message

### Code Examples

#### Correct Usage

```go
client, _ := gopiano.NewClient(gopiano.AndroidClient)

// Step 1: Partner login (REQUIRED)
_, err := client.AuthPartnerLogin()
if err != nil {
    log.Fatal(err)
}

// Step 2: Create user (now validated)
_, err = client.UserCreateUser(
    "user@example.com",
    "password",
    "male",
    "US",
    90210,
    1990,
    false,
)
if err != nil {
    log.Fatal(err)
}
```

#### Incorrect Usage (Will Fail)

```go
client, _ := gopiano.NewClient(gopiano.AndroidClient)

// ❌ Missing AuthPartnerLogin() - will now return clear client-side error
_, err := client.UserCreateUser(
    "user@example.com",
    "password",
    "male",
    "US",
    90210,
    1990,
    false,
)
// Error: "partner authentication token missing: must call AuthPartnerLogin() first..."

// Note: This is a standard Go error, not a responses.PandoraError
// Previous versions would return a server-side PandoraError (error code 0 INTERNAL)
// This is a breaking change - update error handling if you were checking for PandoraError
```

## Authentication Issues

### Error Behavior Change

**Important**: As of the latest version, `UserCreateUser` performs client-side validation and returns a standard Go `error` (not a `responses.PandoraError`) when `AuthPartnerLogin()` has not been called. This is a breaking change from previous behavior where a server-side `PandoraError` would be returned.

**Previous behavior**: Calling `UserCreateUser` without `AuthPartnerLogin()` would result in a server round-trip and return a `responses.PandoraError` (typically error code 0 INTERNAL).

**Current behavior**: The validation happens client-side before making any API call, returning a descriptive standard Go `error` with the message: "partner authentication token missing: must call AuthPartnerLogin() first to establish a partner session before creating a user".

**Migration**: If your code relies on catching `responses.PandoraError` for this specific misuse scenario, update your error handling to check for standard Go errors instead. The error message clearly indicates the issue, so type assertions to `PandoraError` are no longer necessary for this case.

### Required Authentication Flow

All API interactions require a two-step authentication process:

1. **Partner Login** (Step 1):

   ```go
   _, err := client.AuthPartnerLogin()
   ```

   - Establishes partner session
   - Obtains `partnerAuthToken`, `partnerID`, and `syncTime`
   - **Required before any other API methods**

2. **User Login** (Step 2):

   ```go
   // For existing users:
   _, err := client.AuthUserLogin("user@example.com", "password")

   // OR for new users:
   _, err := client.UserCreateUser(...)
   ```

   - Establishes user session
   - Obtains `userAuthToken` and `userID`
   - Required before user-specific methods

### How to Verify Authentication State

The client stores authentication tokens internally. You can verify authentication by checking if methods that require authentication succeed:

```go
// After AuthPartnerLogin():
_, err := client.AuthPartnerLogin()
if err != nil {
    // Partner authentication failed
}

// After user authentication:
_, err := client.UserGetStationList(false)
if err != nil {
    // User authentication may be missing or invalid
}
```

### Token Expiration and Refresh

- Partner tokens: Typically valid for the session duration
- User tokens: May expire after periods of inactivity
- Refresh: Re-authenticate by calling `AuthPartnerLogin()` and `AuthUserLogin()` again

## Parameter Validation Errors

### Requirements for UserCreateUser Parameters

| Parameter     | Requirement                                     | Example                |
| ------------- | ----------------------------------------------- | ---------------------- |
| `username`    | Valid email address format                      | `"user@example.com"`   |
| `password`    | String (no specific format enforced by library) | `"SecurePassword123"`  |
| `gender`      | Exactly `"male"` or `"female"` (case-sensitive) | `"male"` or `"female"` |
| `countryCode` | Must be `"US"`                                  | `"US"`                 |
| `zipCode`     | Valid 5-digit US ZIP code                       | `90210`                |
| `birthYear`   | Integer meeting minimum age (typically 13+)     | `1990`                 |
| `emailOptin`  | Boolean                                         | `true` or `false`      |

### Common Validation Mistakes

1. **Invalid email format**:

   ```go
   // ❌ Wrong
   username := "notanemail"

   // ✅ Correct
   username := "user@example.com"
   ```

2. **Wrong gender value**:

   ```go
   // ❌ Wrong
   gender := "M"  // or "F", "Male", "MALE", etc.

   // ✅ Correct
   gender := "male"  // or "female" (exactly)
   ```

3. **Invalid country code**:

   ```go
   // ❌ Wrong
   countryCode := "USA"  // or "United States", etc.

   // ✅ Correct
   countryCode := "US"
   ```

4. **Invalid ZIP code**:

   ```go
   // ❌ Wrong
   zipCode := 123  // too short
   zipCode := 123456  // too long

   // ✅ Correct
   zipCode := 90210  // 5-digit US ZIP
   ```

5. **Invalid birth year**:

   ```go
   // ❌ Wrong (too young)
   birthYear := 2020  // if current year is 2024, user is only 4 years old

   // ✅ Correct
   birthYear := 1990  // meets minimum age requirement
   ```

### How to Test Parameters Before Calling the API

```go
// Validate email format
if !strings.Contains(username, "@") || !strings.Contains(username, ".") {
    return fmt.Errorf("username must be a valid email address")
}

// Validate gender
if gender != "male" && gender != "female" {
    return fmt.Errorf("gender must be exactly 'male' or 'female'")
}

// Validate country code
if countryCode != "US" {
    return fmt.Errorf("countryCode must be 'US'")
}

// Validate ZIP code
if zipCode < 10000 || zipCode > 99999 {
    return fmt.Errorf("zipCode must be a valid 5-digit US ZIP code")
}

// Validate birth year (example: must be at least 13 years old)
currentYear := time.Now().Year()
if currentYear - birthYear < 13 {
    return fmt.Errorf("birthYear must meet minimum age requirements")
}
```

## Geographic Restrictions

### US-Only Limitation Explanation

Pandora's API enforces geographic restrictions due to music licensing agreements. Account creation and many API features require a US IP address.

### How to Verify IP Address Location

Use an IP geolocation service to verify your IP address:

```bash
# Using curl
curl https://ipinfo.io/json

# Using wget
wget -qO- https://ipinfo.io/json
```

Look for `"country": "US"` in the response.

### VPN Considerations

If using a VPN:

1. **Ensure VPN is set to US location**: Connect to a US-based VPN server
2. **Verify IP location**: Check that your VPN IP shows as US
3. **Test connectivity**: Some VPNs may be blocked by Pandora
4. **Use residential IPs**: Datacenter IPs may be more likely to be blocked

### Workarounds

- Use a US-based server or VPS for API calls
- Use a reliable US VPN service
- Note: The library cannot bypass these restrictions; they are enforced by Pandora's servers

## Rate Limiting

### How to Detect Rate Limiting

Rate limiting may manifest as:

- Error code 0 (INTERNAL) after multiple successful requests
- Sudden failures after a period of successful requests
- Timeout errors
- Connection refused errors

### Recommended Retry Strategies

#### Exponential Backoff

```go
func retryWithBackoff(fn func() error, maxRetries int) error {
    backoff := time.Second
    for i := 0; i < maxRetries; i++ {
        err := fn()
        if err == nil {
            return nil
        }

        // Check if it's a rate limit error
        if isRateLimitError(err) {
            time.Sleep(backoff)
            backoff *= 2  // Exponential backoff
            continue
        }

        return err
    }
    return fmt.Errorf("max retries exceeded")
}

func isRateLimitError(err error) bool {
    // Check for error code 0 after multiple requests
    // or implement your own rate limit detection logic
    return false
}
```

#### Fixed Delay Between Requests

```go
// Wait between requests
time.Sleep(2 * time.Second)

// Make API call
_, err := client.UserCreateUser(...)
```

#### Jittered Backoff

```go
import (
    "math/rand"
    "time"
)

func jitteredBackoff(baseDelay time.Duration) time.Duration {
    jitter := time.Duration(rand.Intn(1000)) * time.Millisecond
    return baseDelay + jitter
}

// Usage
time.Sleep(jitteredBackoff(2 * time.Second))
```

### Best Practices

1. **Limit request frequency**: Wait 1-2 seconds between requests
2. **Batch operations**: Group related requests together
3. **Cache results**: Don't re-fetch data unnecessarily
4. **Monitor for rate limits**: Implement detection and automatic backoff
5. **Respect API limits**: Don't abuse the API with excessive requests

## API Deprecation Concerns

### Status of the Legacy JSON API

The Pandora legacy JSON API (v5) that this library wraps:

- Is **unofficial** and reverse-engineered
- May be **deprecated** or restricted at any time
- Has **limited documentation** and support
- May have **undocumented restrictions** or changes

### Migration Path to Official API

Pandora's official API uses:

- **OAuth2** for authentication
- **GraphQL** for API interactions
- **Official documentation** and support

If you need:

- Official support
- Access to newer features
- Production reliability guarantees
- Long-term API stability

Consider migrating to the official API.

### Alternative Approaches

1. **Use official Pandora API**: Migrate to OAuth2 + GraphQL
2. **Use Pandora web interface**: For manual operations
3. **Use Pandora mobile apps**: For end-user functionality
4. **Monitor API status**: Watch for deprecation notices

### Recommendations

Before using this library for production:

1. ✅ **Verify API availability**: Test that the API still works for your use case
2. ✅ **Test thoroughly**: Ensure all required features work as expected
3. ✅ **Implement error handling**: Handle API failures gracefully
4. ✅ **Monitor for changes**: Watch for API deprecation or restrictions
5. ✅ **Have a backup plan**: Prepare for potential API shutdown
6. ✅ **Consider alternatives**: Evaluate official API or other solutions

### Getting Help

- **GitHub Issues**: Report bugs and ask questions
- **Documentation**: Check README.md and code comments
- **Examples**: See examples/ directory for usage patterns
- **Community**: Check project discussions and contributions

---

For more information, see the [main README.md](README.md) and [examples/](examples/) directory.
