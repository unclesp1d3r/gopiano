# Justfile for gopiano

# Use PowerShell on Windows, bash elsewhere
set windows-shell := ["pwsh.exe", "-NoLogo", "-Command"]
set shell := ["bash", "-lc"]

# Build the project
build:
    go build ./...

# Run all tests
test:
    go test ./...

# Run all tests with race detector and verbose output
test-race:
    go test -race -v ./...

# Run specific test by name pattern
test-run TEST:
    go test -run '{{TEST}}' ./...

# Run all integration tests (requires Pandora credentials)
test-integration:
    go test -tags=integration ./...

# Run specific integration test
test-integration-run TEST:
    go test -tags=integration -run '{{TEST}}' ./...

# Run tests with coverage reporting
test-coverage:
    go test -race -coverprofile=coverage.out -covermode=atomic ./...

# Generate and open HTML coverage report (macOS/Linux)
test-coverage-html: test-coverage
    #!/usr/bin/env sh
    go tool cover -html=coverage.out -o coverage.html
    case $(uname -s) in
        Darwin)
            open coverage.html
            ;;
        Linux)
            xdg-open coverage.html
            ;;
        *)
            echo "Coverage report generated at: coverage.html"
            ;;
    esac

# Generate and open HTML coverage report (Windows PowerShell)
test-coverage-html-win: test-coverage
    go tool cover -html=coverage.out -o coverage.html
    Start-Process coverage.html

# Display function-level coverage in terminal
test-coverage-func: test-coverage
    go tool cover -func=coverage.out

# Display package-level coverage summary
test-coverage-package: test-coverage
    go tool cover -func=coverage.out | grep total

# Run lint and code checks (golangci-lint with config in .golangci.yml)
lint:
    golangci-lint run ./...

# Run lint with autofix (formatting/imports and simple fixes)
lint-fix:
    golangci-lint run --fix ./...

# Run all checks and tests (CI)
ci-check: lint test

# Update dependencies
update-deps:
    go get -u ./...
    go mod tidy
    go mod verify

# Tidy module dependencies
tidy:
    go mod tidy
