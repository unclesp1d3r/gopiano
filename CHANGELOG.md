# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- **BREAKING**: `UserCreateUser` now performs client-side validation and returns a standard Go `error` (via `fmt.Errorf`) when `AuthPartnerLogin()` has not been called, instead of propagating a `responses.PandoraError` from the server. This provides clearer, more actionable error messages for misuse scenarios. Previously, calling `UserCreateUser` without first calling `AuthPartnerLogin()` would result in an opaque server-side error (typically "Pandora Error: 0 INTERNAL"). The new behavior fails fast with a descriptive client-side error message: "partner authentication token missing: must call AuthPartnerLogin() first to establish a partner session before creating a user".

  **Migration**: If your code relies on catching `responses.PandoraError` for this specific misuse scenario, update your error handling to check for standard Go errors instead. The error message clearly indicates the issue, so type assertions to `PandoraError` are no longer necessary for this case.

### Added

- Client-side validation in `UserCreateUser` to check for required `partnerAuthToken` before making API calls
- Enhanced error messages with troubleshooting guidance for error code 0 (INTERNAL)
- Comprehensive documentation in README.md and TROUBLESHOOTING.md
- Example code demonstrating proper authentication flow and user creation
- `GetErrorGuidance()` helper function in `responses` package for contextual error diagnosis
