# AGENTS

Guidelines for AI coding assistants working on gopiano, a thin wrapper around Pandora.com's JSON API. This file is the single source of truth for assistant behavior in this repository.

## Project Overview

gopiano is a thin wrapper library around Pandora.com's reverse-engineered JSON API. The library provides a Go client for interacting with Pandora's API endpoints. The Pandora API documentation is available at <https://6xq.net/pandora-apidoc/json/> and <https://6xq.net/pandora-apidoc/rest/>.

## Project Context

- **Purpose**: Thin wrapper library around Pandora.com's reverse-engineered JSON API
- **API Documentation**: <https://6xq.net/pandora-apidoc/json/> and <https://6xq.net/pandora-apidoc/rest/>
- **Module Path**: `github.com/unclesp1d3r/gopiano`
- **Go Version**: 1.24+ (read-only module download mode)
- **Status**: Alpha quality; needs proper tests and error handling improvements

## Architecture

The project follows a three-layer architecture:

1. **Main Package** (`gopiano`): Contains the `Client` struct and all client methods organized by domain (`auth.go`, `station.go`, `user.go`, `misc.go`)
2. **Requests Package** (`requests`): Contains structs for JSON marshaling that mirror Pandora API request formats exactly
3. **Responses Package** (`responses`): Contains structs for JSON unmarshaling that mirror Pandora API response formats exactly

This structure ensures that the library remains a thin wrapper, not an abstraction, maintaining fidelity to the Pandora API.

## Codebase Structure

- **Core Client**: `gopiano.go` - Client struct, encryption/decryption, `PandoraCall`, `BlowfishCall`, generic helpers
- **Request Types**: `requests/` - Split by domain: `auth.go`, `station.go`, `user.go`, `misc.go`
- **Response Types**: `responses/` - Split by domain: `errors.go`, `common.go`, `auth.go`, `station.go`, `user.go`
- **Feature Files**: `auth.go`, `station.go`, `user.go`, `misc.go` - Client methods organized by domain
- **Tests**: `*_test.go` files - table-driven tests for missing-token validation

## Key Components

### Client Struct

The `Client` struct (defined in `gopiano.go`) is the main entry point for all API interactions. It contains:

- `description`: `ClientDescription` for device emulation
- `http`: HTTP client for making requests
- `encrypter`/`decrypter`: Blowfish ciphers for encryption/decryption
- `timeOffset`: Time offset for sync time calculations
- `partnerAuthToken`/`partnerID`: Partner authentication credentials
- `userAuthToken`/`userID`: User authentication credentials

### ClientDescription

The `ClientDescription` struct describes a particular type of client to emulate, including:

- `DeviceModel`: Device model identifier (e.g., "android-generic")
- `Username`/`Password`: Partner credentials
- `BaseURL`: Base URL for API endpoints
- `EncryptKey`/`DecryptKey`: Blowfish encryption keys
- `Version`: API version

### AndroidClient

The `AndroidClient` global variable provides a pre-configured `ClientDescription` for Android device emulation. `DefaultAndroidClient()` returns a fresh copy and is preferred over the mutable global.

### Encryption

The library uses Blowfish ECB (Electronic Codebook) mode encryption for API communication:

- **Encryption**: `encrypt()` method encrypts request data using Blowfish ECB mode
- **Decryption**: `decrypt()` method decrypts response data using Blowfish ECB mode
- **Implementation**: Uses `golang.org/x/crypto/blowfish` package (marked as deprecated but required by Pandora API)
- **Keys**: Encryption/decryption keys are provided by `ClientDescription`

### API Interaction Patterns

The library provides two public methods and two internal generic helpers for API interaction:

- **`PandoraCall()`**: Basic HTTP POST method for unencrypted API calls. Handles URL construction, authentication tokens, and response parsing.
- **`BlowfishCall()`**: Wrapper around `PandoraCall()` that first encrypts the request body using Blowfish encryption before sending.
- **`blowfishCallJSON[Resp]()`**: Internal generic helper that marshals, encrypts, calls, and unmarshals in one step. Used by most API methods.
- **`blowfishCallVoid()`**: Internal helper for API methods that return only an error with no response body.

These methods handle:

- URL construction with query parameters (method, partner_id, user_id, auth_token)
- HTTP request creation with proper headers (User-Agent: "gopiano", Content-Type: "text/plain")
- Response parsing and error handling
- Pandora API error detection and conversion to `*PandoraError` types

### Functional Options

`NewClient` accepts optional `Option` arguments:

- `WithHTTPClient(hc)` - Replace the default HTTP client
- `WithTimeout(d)` - Set timeout on the default HTTP client

### Error Handling

The library uses a custom error type for Pandora API errors:

- **`PandoraError`**: Defined in `responses/errors.go`, implements the `error` interface with pointer receiver
- **`ErrorCodeMap`**: Maps Pandora API error codes to human-readable error messages
- **Error Detection**: `PandoraCall()` uses `bytes.Contains` to check for `"stat":"fail"` and converts to `*PandoraError`
- **Usage**: Use `var pe *responses.PandoraError` with `errors.As(err, &pe)` to check for API errors

### Naming Conventions

Structs are named after their corresponding API methods:

- Request structs: Match API method names (e.g., `AuthPartnerLogin`, `UserGetStationList`)
- Response structs: Match API method names (e.g., `AuthPartnerLogin`, `UserGetStationList`)
- Client methods: Match API method names with appropriate prefixes (e.g., `AuthPartnerLogin()`, `UserGetStationList()`)

## Development Workflow

1. **Before Making Changes**:

   - Run `just lint` or `golangci-lint run ./...` (config in `.golangci.yml`) - enforces formatting and linting
   - Run `just test` or `go test ./...` to verify existing tests pass
   - Use `just test-run TestName` or `go test -run TestName ./path` to target specific tests

2. **After Making Changes**:

   - Run `just lint` or `golangci-lint run ./...` to ensure code quality
   - Run `just build` or `go build ./...` to verify compilation
   - Run `just test` or `go test ./...` to verify tests pass
   - Update README.md if adding new commands or client capabilities

## Code Patterns

### Adding New API Methods

When adding a new Pandora API method:

1. **Add Request Struct** (in the appropriate `requests/*.go` domain file):

   ```go
   // MethodName represents the request data for api.methodName.
   type MethodName struct {
       FieldName     string `json:"fieldName"` // Match Pandora API exactly
       UserAuthToken string `json:"userAuthToken"`
       SyncTime      int    `json:"syncTime"`
   }
   ```

2. **Add Response Struct** (in the appropriate `responses/*.go` domain file):

   ```go
   // MethodName represents the response from api.methodName.
   type MethodName struct {
       Result struct {
           // ... fields matching Pandora API exactly
       } `json:"result"`
   }
   ```

3. **Add Client Method** (in appropriate file: `auth.go`, `station.go`, `user.go`, or `misc.go`):

   ```go
   // MethodName does X. Calls API method "api.methodName".
   func (c *Client) MethodName(ctx context.Context, ...) (*responses.MethodName, error) {
       userAuthToken, err := c.getUserAuthToken("doing X")
       if err != nil {
           return nil, err
       }
       requestData := requests.MethodName{
           // ... populate fields
           UserAuthToken: userAuthToken,
           SyncTime:      c.GetSyncTime(),
       }
       resp, err := blowfishCallJSON[responses.MethodName](ctx, c, "api.methodName", requestData)
       if err != nil {
           return nil, fmt.Errorf("do X: %w", err)
       }
       return resp, nil
   }
   ```

### API Struct Conventions

- **Mirror Pandora JSON exactly**: Field names, types, and JSON tags must match API documentation
- **Maintain JSON tags**: When editing `requests/` or `responses/` packages, preserve existing JSON tags
- **Use omitempty**: For optional fields that may not be sent
- **Preserve typos**: If Pandora API has typos (e.g., "IncludeDemographics"), keep them with `//nolint:tagliatelle` comment

### Error Handling

- **Wrap errors**: Use `fmt.Errorf("context: %w", err)` for error wrapping
- **Handle immediately**: Don't defer error handling; check and return errors immediately
- **Avoid panics**: No panics outside tests
- **API errors**: Pandora API errors are returned as `*responses.PandoraError` (pointer). Use `errors.As(err, &pe)` with `var pe *responses.PandoraError`

### Authentication Flow

1. Call `AuthPartnerLogin()` to get partner credentials
2. Call `AuthUserLogin(username, password)` to get user credentials
3. Subsequent calls use `getUserAuthToken()`/`getPartnerAuthToken()` to read tokens under lock
4. `PandoraCall` reads partner/user IDs under lock for URL construction

## Code Quality Standards

### Formatting

- **Tool**: `gofmt`/`gofumpt` (enforced by golangci-lint)
- **Indentation**: Tabs (not spaces)
- **Line Length**: Keep under ~120 characters
- **Imports**: Standard lib, third-party, local (sorted by `goimports`)
- **Unused imports**: Remove them (caught by `goimports`)

### Naming

- **Exported**: PascalCase with doc comments
- **Unexported**: camelCase, concise
- **Package-level**: Document all exported functions, types, and variables

### Types

- **Prefer explicit structs**: Avoid `any`/`interface{}` unless type assertions are intentional
- **Constants**: Define constants for protocol-level values (reject magic numbers)
- **No GOPATH**: Keep GOPATH usage minimal; use modules
- **Module path**: When adding or editing API structs, use the module path `github.com/unclesp1d3r/gopiano` - request structs go in the appropriate `requests/*.go` domain file and response structs go in the appropriate `responses/*.go` domain file (see [Project Context](#project-context) for module details)

### Context

- **All public methods**: Accept `context.Context` as first parameter
- **Context propagation**: `PandoraCall`, `BlowfishCall`, and generic helpers all propagate context
- **Cancellation**: `PandoraCall` checks `ctx.Done()` before starting the HTTP request

### Concurrency

- **Guard shared state**: Use channels or sync primitives
- **Avoid data races**: Be explicit about synchronization

### Logging

- **Structured logs**: Prefer structured logging
- **No fmt.Println**: Avoid in library code (except tests/examples)

### Security

- **No plaintext credentials**: Don't store credentials in plaintext
- **Use crypto helpers**: Use `golang.org/x/crypto` helpers where available
- **Blowfish encryption**: Required for many API calls (handled by `BlowfishCall`)

## Testing

- **Style**: Table-driven tests preferred
- **Deterministic**: Keep tests deterministic
- **Parallel**: Use `t.Parallel()` when safe
- **Coverage**: This repo currently lacks proper tests; prioritize coverage when adding features
- **Test files**: `*_test.go` alongside source files
- **Integration Tests**: Integration tests use the build tag `//go:build integration` as seen in `gopiano_test.go`. These tests require valid Pandora credentials and make actual API calls. Run with `just test-integration` or `go test -tags=integration ./...`

## Common Operations

### Adding a New Station Method

1. Check Pandora API docs for method signature
2. Add request struct to `requests/station.go`
3. Add response struct to `responses/station.go`
4. Add method to `station.go` using `blowfishCallJSON` or `blowfishCallVoid` helpers
5. Use `getUserAuthToken()` to obtain token under lock
6. Include `SyncTime: c.GetSyncTime()` for authenticated calls
7. Add input validation for required string parameters
8. Add doc comment describing what the method does

### Modifying Existing Methods

- **Preserve API compatibility**: Don't change request/response structs unless matching API changes
- **Update doc comments**: Keep comments accurate
- **Maintain consistency**: Follow patterns in `auth.go`, `station.go`, `user.go`, `misc.go`

## Things to Avoid

- **Don't change JSON tags**: They must match Pandora API exactly
- **Don't add dependencies**: Keep the library thin; only `golang.org/x/crypto` is allowed
- **Don't break API compatibility**: This is a wrapper, not an abstraction
- **Don't skip error handling**: Always check and return errors
- **Don't use magic numbers**: Define constants for protocol values
- **Don't ignore linter**: Fix all `golangci-lint` warnings before committing

## File Organization

- **Keep related code together**: `auth.go` for auth, `station.go` for stations, etc.
- **Consistent exports**: Keep `misc.go`, `station.go`, `user.go` consistent with exported behavior
- **Doc comments**: Add doc comments for all public functions

## Dependencies

- **golang.org/x/crypto/blowfish**: Required for Blowfish encryption. This package is marked as deprecated but is required by the Pandora API. Use `//nolint:staticcheck` when importing.

## Code Quality

The project uses `golangci-lint` for code quality enforcement:

- **Configuration**: `.golangci.yml` defines linting rules and settings
- **Nolint Directives**: Use nolint directives sparingly with clear justification comments
- **Common Nolints**:
  - `//nolint:staticcheck` for required but deprecated packages
  - `//nolint:tagliatelle` for JSON tag names that match API typos
  - `//nolint:gochecknoglobals` for intentionally exported global variables
  - `//nolint:gosec` for `json.Marshal` of structs containing partner credentials or encrypted passwords
  - `//nolint:mnd` for magic numbers that are reasonable defaults (e.g., timeout durations)

## Documentation Standards

- **Package Documentation**: All packages should have package-level documentation
- **Exported Symbols**: All exported types, functions, and methods must have godoc comments
- **Go Conventions**: Follow standard Go documentation conventions
- **Examples**: Include usage examples in doc comments when helpful

## Build and Test Commands

Use these commands from the module root. The project includes a `justfile` with convenient recipes - use `just <recipe>` or the underlying Go commands directly.

```bash
# Build
just build
# or
go build ./...

# Run all tests
just test
# or
go test ./...

# Run all tests with race detector and verbose output
just test-race
# or
go test -race -v ./...

# Run specific test by name pattern (unit or integration)
just test-run TestName
# or
go test -run '^TestName$' ./...

# Run all integration tests (requires Pandora credentials and hits live endpoints)
just test-integration
# or
go test -tags=integration ./...

# Run a specific integration test
just test-integration-run Test_AuthPartnerLogin_1
# or
go test -tags=integration -run '^Test_AuthPartnerLogin_1$' ./...

# Lint and format check
just lint
# or
golangci-lint run ./...

# Lint with autofix (formatting/imports and simple fixes)
just lint-fix
# or
golangci-lint run --fix ./...

# Run all checks and tests (CI)
just ci-check

# Test coverage
just test-coverage           # Generate coverage report
just test-coverage-html      # View HTML coverage report
just test-coverage-func      # Show function-level coverage
```

See the `justfile` for all available recipes. Run `just` without arguments to see the full list.

## Agent Rules <!-- tessl-managed -->

@.tessl/RULES.md follow the [instructions](.tessl/RULES.md)
