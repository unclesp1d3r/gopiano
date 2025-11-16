# GitHub Copilot Instructions for gopiano

This document provides context-specific instructions for GitHub Copilot when working on the gopiano project.

## Project Overview

gopiano is a thin wrapper library around Pandora.com's reverse-engineered JSON API. The library provides a Go client for interacting with Pandora's API endpoints.

## Go Best Practices

When writing Go code for gopiano:

- **Idiomatic Go**: Follow Go idioms and conventions. Use `gofmt`/`gofumpt` formatting.
- **Error Handling**: Always check and return errors. Never ignore errors. Use `fmt.Errorf("context: %w", err)` for error wrapping.
- **Context Usage**: For HTTP requests, use `context.Context` as the first parameter. Currently, `PandoraCall` uses `context.Background()` - consider making it configurable in the future.
- **Documentation**: Add godoc comments for all exported types, functions, and methods.
- **Naming**: Use PascalCase for exported symbols, camelCase for unexported. Keep names concise and descriptive.

## Project-Specific Patterns

### API Method Implementation

All API methods should follow this pattern:

1. **Create request struct** from `requests` package
2. **Marshal to JSON** using `json.Marshal`
3. **Call `PandoraCall()` or `BlowfishCall()`** depending on encryption requirements
4. **Unmarshal response** into struct from `responses` package

Example pattern:

```go
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
    err = c.BlowfishCall("http://", "api.methodName", requestDataReader, &resp)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
```

### Encryption Requirements

- **Use `BlowfishCall()`** for methods requiring encryption (most user/station operations)
- **Use `PandoraCall()`** for unencrypted requests (typically authentication methods)
- Always include `SyncTime: c.GetSyncTime()` in authenticated requests
- Include appropriate auth tokens (`UserAuthToken` or `PartnerAuthToken`) in requests

### Struct Conventions

- **Request structs**: Add to `requests/requests.go`, named after API methods (e.g., `AuthPartnerLogin`, `UserGetStationList`)
- **Response structs**: Add to `responses/responses.go`, named after API methods
- **JSON tags**: Must match Pandora API exactly. Preserve typos if they exist in the API (use `//nolint:tagliatelle` comment)
- **Field types**: Match Pandora API types exactly (string, int, bool, etc.)

### Error Handling

- **Pandora API errors**: Returned as `responses.PandoraError` types (implements `error` interface)
- **Error codes**: Check against `responses.ErrorCodeMap` for human-readable error messages
- **Error wrapping**: Use `fmt.Errorf("context: %w", err)` to provide context
- **Immediate handling**: Check errors immediately, don't defer error handling

### Testing Patterns

- **Integration tests**: Use `//go:build integration` build tag (see `gopiano_test.go`)
- **Unit tests**: Test encryption/decryption logic, error handling, and utility functions
- **Test structure**: Use table-driven tests when appropriate
- **Test naming**: Follow Go conventions (`Test_FunctionName_Scenario`)

### Encryption Implementation

When working with Blowfish encryption:

- **ECB mode**: Maintain ECB (Electronic Codebook) mode as implemented in `encrypt()` and `decrypt()` methods
- **Padding**: Use proper padding as implemented (8-byte blocks)
- **Keys**: Use keys from `ClientDescription` (`EncryptKey` and `DecryptKey`)
- **Hex encoding**: Encrypted data is hex-encoded for transmission

### HTTP Client Usage

- **User-Agent**: Always set to "gopiano" (see `PandoraCall` implementation)
- **Content-Type**: Always set to "text/plain" for API requests
- **HTTP Client**: Use the existing `http.Client` in the `Client` struct
- **URL construction**: Build URLs with appropriate query parameters (method, partner_id, user_id, auth_token)

### Linting and Code Quality

- **Configuration**: Respect `.golangci.yml` configuration
- **Nolint directives**: Use sparingly with clear justification comments
- **Formatting**: Code must pass `golangci-lint run ./...`
- **Common nolints**:
  - `//nolint:staticcheck` for required but deprecated packages (e.g., `golang.org/x/crypto/blowfish`)
  - `//nolint:tagliatelle` for JSON tag names that match API typos
  - `//nolint:gochecknoglobals` for intentionally exported global variables

### Documentation Standards

- **Package documentation**: Add package-level doc comments
- **Exported symbols**: Document all exported types, functions, and methods
- **Examples**: Include usage examples in doc comments when helpful
- **API references**: Reference Pandora API documentation when relevant

## Common Operations

### Adding a New API Method

1. Check Pandora API documentation for method signature
2. Add request struct to `requests/requests.go`
3. Add response struct to `responses/responses.go`
4. Add method to appropriate file (`auth.go`, `station.go`, `user.go`, or `misc.go`)
5. Use `BlowfishCall` for encrypted requests, `PandoraCall` for plain
6. Include `SyncTime: c.GetSyncTime()` for authenticated calls
7. Add doc comment describing what the method does

### Modifying Existing Methods

- **Preserve API compatibility**: Don't change request/response structs unless matching API changes
- **Update doc comments**: Keep comments accurate and up-to-date
- **Maintain consistency**: Follow patterns in existing methods

## Things to Avoid

- **Don't change JSON tags**: They must match Pandora API exactly
- **Don't add dependencies**: Keep the library thin; only `golang.org/x/crypto` is allowed
- **Don't break API compatibility**: This is a wrapper, not an abstraction
- **Don't skip error handling**: Always check and return errors
- **Don't use magic numbers**: Define constants for protocol values
- **Don't ignore linter**: Fix all `golangci-lint` warnings before committing

## Architecture Notes

- **Client struct**: Main client type with HTTP client, encryption ciphers, and auth tokens
- **ClientDescription**: Describes device emulation (device model, credentials, encryption keys)
- **AndroidClient**: Global variable with Android client configuration
- **Three-layer structure**: Main package (Client methods), requests package (request structs), responses package (response structs)

## Dependencies

- **golang.org/x/crypto/blowfish**: Required for Blowfish encryption (marked as deprecated but required by Pandora API)
- **Standard library**: Use standard library packages for HTTP, JSON, encoding, etc.

## Additional Resources

- See `AGENTS.md` for comprehensive AI agent instructions
- See `CONTRIBUTING.md` for contribution guidelines
- See `.golangci.yml` for linting configuration
- Pandora API documentation: <https://6xq.net/pandora-apidoc/json/> and <https://6xq.net/pandora-apidoc/rest/>
