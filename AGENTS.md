# AGENTS

Guidelines for AI coding assistants working on gopiano, a thin wrapper around Pandora.com's JSON API.

## Project Context

- **Purpose**: Thin wrapper library around Pandora.com's reverse-engineered JSON API
- **API Documentation**: <https://6xq.net/pandora-apidoc/json/> and <https://6xq.net/pandora-apidoc/rest/>
- **Module Path**: `github.com/cellofellow/gopiano`
- **Go Version**: 1.24+ (read-only module download mode)
- **Status**: Alpha quality; needs proper tests and error handling improvements

## Codebase Structure

- **Core Client**: `gopiano.go` - Client struct, encryption/decryption, `PandoraCall`, `BlowfishCall`
- **Request Types**: `requests/requests.go` - Structs for JSON marshaling (mirror Pandora API exactly)
- **Response Types**: `responses/responses.go` - Structs for JSON unmarshaling (mirror Pandora API exactly)
- **Feature Files**: `auth.go`, `station.go`, `user.go`, `misc.go` - Client methods organized by domain
- **Tests**: `*_test.go` files (currently minimal coverage)

## Development Workflow

1. **Before Making Changes**:
   - Run `golangci-lint run ./...` (config in `.golangci.yml`) - enforces formatting and linting
   - Run `go test ./...` to verify existing tests pass
   - Use `go test -run TestName ./path` to target specific tests

2. **After Making Changes**:
   - Run `golangci-lint run ./...` to ensure code quality
   - Run `go build ./...` to verify compilation
   - Run `go test ./...` to verify tests pass
   - Update README.md if adding new commands or client capabilities

## Code Patterns

### Adding New API Methods

When adding a new Pandora API method:

1. **Add Request Struct** (in `requests/requests.go`):

   ```go
   // MethodName represents the request data for api.methodName.
   type MethodName struct {
       FieldName string `json:"fieldName"` // Match Pandora API exactly
       SyncTime  int    `json:"syncTime"`  // Required for authenticated calls
   }
   ```

2. **Add Response Struct** (in `responses/responses.go`):

   ```go
   // MethodName represents the response from api.methodName.
   type MethodName struct {
       Stat string `json:"stat"` // "ok" or "fail"
       // ... other fields matching Pandora API exactly
   }
   ```

3. **Add Client Method** (in appropriate file: `auth.go`, `station.go`, `user.go`, or `misc.go`):

   ```go
   // MethodName does X. Calls API method "api.methodName".
   func (c *Client) MethodName(...) (*responses.MethodName, error) {
       requestData := requests.MethodName{
           // ... populate fields
           SyncTime: c.GetSyncTime(), // For authenticated calls
       }
       requestDataEncoded, err := json.Marshal(requestData)
       if err != nil {
           return nil, err
       }
       requestDataReader := bytes.NewReader(requestDataEncoded)

       var resp responses.MethodName
       // Use BlowfishCall for encrypted requests, PandoraCall for plain
       err = c.BlowfishCall("http://", "api.methodName", requestDataReader, &resp)
       if err != nil {
           return nil, err
       }
       return &resp, nil
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
- **API errors**: Pandora API errors are returned as `responses.PandoraError` (implements `error`)

### Authentication Flow

1. Call `AuthPartnerLogin()` to get partner credentials
2. Call `AuthUserLogin(username, password)` to get user credentials
3. Subsequent calls use `c.userAuthToken` and `c.userID` automatically via `PandoraCall`

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

- **Prefer explicit structs**: Avoid `interface{}` unless type assertions are intentional
- **Constants**: Define constants for protocol-level values (reject magic numbers)
- **No GOPATH**: Keep GOPATH usage minimal; use modules

### Context

- **Long-running functions**: Accept `context.Context` as first parameter
- **Respect deadlines**: Check context cancellation and deadlines
- **Current limitation**: `PandoraCall` uses `context.Background()` - consider making it configurable

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

## Common Operations

### Adding a New Station Method

1. Check Pandora API docs for method signature
2. Add request struct to `requests/requests.go`
3. Add response struct to `responses/responses.go`
4. Add method to `station.go` following existing patterns
5. Use `BlowfishCall` for encrypted requests, `PandoraCall` for plain
6. Include `SyncTime: c.GetSyncTime()` for authenticated calls
7. Add doc comment describing what the method does

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

## Build and Test Commands

```bash
# Run all tests
go test ./...

# Run specific test
go test -run TestName ./path

# Lint and format check
golangci-lint run ./...

# Build
go build ./...
```
