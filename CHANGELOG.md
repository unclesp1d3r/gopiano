# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- **BREAKING**: `UserCreateUser` and other authentication-requiring methods now perform client-side validation and return standard Go `error` values instead of `responses.PandoraError` when required authentication tokens are missing. This provides clearer error messages and fails fast instead of making unsuccessful API calls.

  **Migration**: Update error handling code that specifically checks for `responses.PandoraError` in missing authentication scenarios to handle standard Go errors instead.

### Added

- Client-side validation in all authentication-requiring methods (`UserCreateUser`, `AuthUserLogin`, `UserEmailPassword`, and all user/station/misc methods) to check for required authentication tokens before making API calls
- Enhanced error messages with troubleshooting guidance for error code 0 (INTERNAL)
- Comprehensive documentation in README.md and TROUBLESHOOTING.md
- Example code demonstrating proper authentication flow and user creation
- `GetErrorGuidance()` helper function in `responses` package for contextual error diagnosis
