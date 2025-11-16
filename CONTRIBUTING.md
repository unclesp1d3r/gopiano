# Contributing to gopiano

Thank you for your interest in contributing to gopiano! This document provides guidelines and instructions for contributing to the project.

## Welcome

We welcome contributions of all kinds:

- Bug reports
- Feature requests
- Documentation improvements
- Code contributions
- Testing and test improvements
- Code reviews

Every contribution helps make gopiano better for everyone.

## Ways to Contribute

### Bug Reports

Found a bug? Please open an issue with:

- A clear, descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Environment details (Go version, OS, etc.)
- Code examples or error messages (if applicable)

### Feature Requests

Have an idea for a new feature? Open an issue with:

- A clear description of the feature
- Use cases and examples
- Potential implementation approach (if you have ideas)
- Any related issues or discussions

### Documentation

Documentation improvements are always welcome:

- Fix typos or clarify explanations
- Add examples or usage patterns
- Improve code comments
- Update README or other documentation files

### Code Contributions

See the [Development Workflow](#development-workflow) section below for details on contributing code.

### Testing

Help improve test coverage:

- Write tests for existing functionality
- Add integration tests (see testing guidelines below)
- Report test failures or flaky tests

## Getting Started

### Prerequisites

- Go 1.24 or later
- Git
- Basic familiarity with Go and the project structure
- [just](https://github.com/casey/just) (optional, but recommended for running common tasks)

### Setup

1. **Fork the repository** on GitHub

2. **Clone your fork**:

   ```sh
   git clone https://github.com/YOUR_USERNAME/gopiano.git
   cd gopiano
   ```

3. **Add upstream remote**:

   ```sh
   git remote add upstream https://github.com/unclesp1d3r/gopiano.git
   ```

   **Note**: The `gopiano` project originated from `https://github.com/cellofellow/gopiano` but is now maintained under `https://github.com/unclesp1d3r/gopiano`. Please use the latter as the canonical upstream repository for contributions.

4. **Install dependencies** (if any):

   ```sh
   go mod download
   ```

5. **Verify setup**:

   ```sh
   # Using justfile (recommended)
   just test
   just lint

   # Or using Go commands directly
   go test ./...
   golangci-lint run ./...
   ```

### Project Structure

- **Main package**: `gopiano.go`, `auth.go`, `user.go`, `station.go`, `misc.go` - Client methods
- **Requests package**: `requests/requests.go` - Request structs for API calls
- **Responses package**: `responses/responses.go` - Response structs from API calls
- **Tests**: `*_test.go` files alongside source files

For more architectural details, see [AGENTS.md](AGENTS.md).

### Using the Justfile

This project includes a `justfile` with convenient recipes for common development tasks. If you have [just](https://github.com/casey/just) installed, you can use these commands:

**Build and Test:**

- `just build` - Build the project
- `just test` - Run all tests
- `just test-race` - Run tests with race detector
- `just test-run TEST` - Run specific test by name pattern
- `just test-integration` - Run integration tests
- `just test-coverage` - Run tests with coverage
- `just test-coverage-html` - Generate and view HTML coverage report

**Linting:**

- `just lint` - Run golangci-lint
- `just lint-fix` - Run linter with autofix

**CI and Dependencies:**

- `just ci-check` - Run all checks and tests (for CI)
- `just update-deps` - Update dependencies
- `just tidy` - Tidy module dependencies

You can also run `just` without arguments to see all available recipes. All justfile recipes are optional - you can use the underlying Go commands directly if you prefer.

## Development Workflow

### 1. Create a Branch

Create a feature branch from `main`:

```sh
git checkout -b feature/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

Use descriptive branch names:

- `feature/add-new-method` - For new features
- `fix/encryption-bug` - For bug fixes
- `docs/update-readme` - For documentation

### 2. Make Changes

- Follow Go best practices and idioms
- Follow project patterns (see [AGENTS.md](AGENTS.md) for details)
- Add godoc comments for exported symbols
- Maintain consistency with existing code

### 3. Test Your Changes

Run tests:

```sh
# Using justfile (recommended)
just test                    # Run all tests
just test-run TestName       # Run specific test by name pattern
just test-integration        # Run integration tests (requires Pandora credentials)
just test-race               # Run tests with race detector
just test-coverage           # Run tests with coverage
just test-coverage-html      # Generate and view HTML coverage report

# Or using Go commands directly
go test ./...
go test -run TestName ./path
go test -tags=integration ./...
go test -race -v ./...
```

Run the linter:

```sh
# Using justfile (recommended)
just lint        # Run linter
just lint-fix    # Run linter with autofix

# Or using golangci-lint directly
golangci-lint run ./...
golangci-lint run --fix ./...
```

Verify compilation:

```sh
# Using justfile (recommended)
just build

# Or using Go commands directly
go build ./...
```

### 4. Commit Your Changes

Write clear, descriptive commit messages:

```text
Short summary (50 chars or less)

More detailed explanation if needed. Wrap at 72 characters.
Explain what and why, not how.

- Bullet points are okay too
- Use present tense ("Add feature" not "Added feature")
```

### 5. Push and Create Pull Request

```sh
git push origin feature/your-feature-name
```

Then create a pull request on GitHub with:

- **Clear title**: Describe what the PR does
- **Description**: Explain the changes and why
- **Related issues**: Link to any related issues
- **Testing**: Describe how you tested the changes

## Code Standards

### Go Best Practices

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt`/`gofumpt` for formatting (enforced by linter)
- Use tabs for indentation (not spaces)
- Keep lines under ~120 characters
- Sort imports: standard lib, third-party, local

### Project-Specific Patterns

When adding new API methods:

1. Add request struct to `requests/requests.go`
2. Add response struct to `responses/responses.go`
3. Add method to appropriate file (`auth.go`, `station.go`, `user.go`, or `misc.go`)
4. Use `BlowfishCall` for encrypted requests, `PandoraCall` for plain
5. Include `SyncTime: c.GetSyncTime()` for authenticated calls
6. Add doc comment describing what the method does

See [AGENTS.md](AGENTS.md) for detailed patterns and examples.

### Documentation

- Add godoc comments for all exported types, functions, and methods
- Keep comments accurate and up-to-date
- Include usage examples when helpful

### Error Handling

- Always check and return errors
- Use `fmt.Errorf("context: %w", err)` for error wrapping
- Don't ignore errors
- Return `PandoraError` types for API errors

### Linting

- Code must pass `golangci-lint run ./...`
- Respect `.golangci.yml` configuration
- Use nolint directives sparingly with clear justification

## Testing Guidelines

### Unit Tests

- Test individual functions and methods
- Use table-driven tests when appropriate
- Test error cases and edge cases
- Keep tests deterministic

### Integration Tests

Integration tests require Pandora credentials and use the `//go:build integration` build tag:

```go
//go:build integration

package gopiano

import "testing"

func Test_IntegrationExample(t *testing.T) {
    // Integration test code
}
```

Run with:

```sh
# Using justfile (recommended)
just test-integration

# Or using Go commands directly
go test -tags=integration ./...
```

**Note**: Integration tests may require valid Pandora credentials and make actual API calls.

### Test Coverage

While the project currently has minimal test coverage, we aim to improve it. When adding new features:

- Add tests for new functionality
- Aim for good coverage of error paths
- Test edge cases and boundary conditions

## Pull Request Process

### Before Submitting

- [ ] Code follows project conventions
- [ ] Tests pass (`just test` or `go test ./...`)
- [ ] Linter passes (`just lint` or `golangci-lint run ./...`)
- [ ] Code compiles (`just build` or `go build ./...`)
- [ ] Documentation is updated (if needed)
- [ ] Commit messages are clear

You can run `just ci-check` to run both linting and tests in one command.

### Review Process

1. **Automated Checks**: CI will run tests and linting
2. **Maintainer Review**: A maintainer will review your PR
3. **Feedback**: You may receive feedback or requested changes
4. **Approval**: Once approved, the PR will be merged

### What to Include in PR Description

- Summary of changes
- Motivation and context
- How to test the changes
- Screenshots or examples (if applicable)
- Related issues

## Issue Reporting

### Bug Reports

When reporting a bug, include:

- **Title**: Clear, descriptive title
- **Description**: What happened vs. what you expected
- **Steps to Reproduce**: Detailed steps to reproduce
- **Environment**: Go version, OS, gopiano version
- **Code Example**: Minimal code example (if applicable)
- **Error Messages**: Full error messages or logs

### Feature Requests

When requesting a feature:

- **Title**: Clear feature description
- **Description**: What you want and why
- **Use Cases**: Examples of how it would be used
- **Alternatives**: Other solutions you've considered

## Questions

If you have questions:

- Check existing issues and discussions
- Open an issue for discussion
- See [AGENTS.md](AGENTS.md) for project architecture details
- See [GOVERNANCE.md](GOVERNANCE.md) for project governance

## Code of Conduct

Please note that this project follows a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Recognition

Contributors will be recognized in [CONTRIBUTORS.md](CONTRIBUTORS.md). Thank you for contributing to gopiano!
