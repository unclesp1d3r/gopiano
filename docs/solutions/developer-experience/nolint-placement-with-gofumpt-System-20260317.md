---
module: System
date: 2026-03-17
problem_type: developer_experience
component: tooling
symptoms:
  - golangci-lint pre-commit hook fails with gosec G117 despite //nolint:gosec comment
  - Pre-commit hook reports 'files were modified by this hook' on every commit attempt
  - 3 consecutive failed commits before identifying the root cause
root_cause: config_error
resolution_type: workflow_improvement
severity: medium
tags: [golangci-lint, gofumpt, nolint, gosec, pre-commit, go]
---

# Troubleshooting: nolint Directive Ignored After gofumpt Reformats Multi-Line Call

## Problem

When adding `//nolint:gosec` to a `json.Marshal()` call, the golangci-lint pre-commit hook repeatedly fails because gofumpt reformats the call onto multiple lines, moving the nolint comment to a line different from where gosec reports the issue.

## Environment

- Module: System (gopiano tooling)
- Go Version: 1.24+
- Affected Component: golangci-lint + gofumpt formatter interaction
- Date: 2026-03-17

## Symptoms

- `golangci-lint-mod` pre-commit hook fails with `G117: Marshaled struct field "Password" matches secret pattern (gosec)`
- Adding `//nolint:gosec` on the same line as `json.Marshal(requestData)` does not suppress the warning
- Hook reports "files were modified by this hook" because gofumpt reformats the code
- Each commit attempt: gofumpt splits the line, nolint moves to wrong line, gosec fires, commit fails

## What Didn't Work

**Attempted Solution 1:** Place nolint on the single-line `json.Marshal(requestData)` call

```go
requestDataEncoded, err := json.Marshal(requestData) //nolint:gosec // password is encrypted
```

- **Why it failed:** gofumpt reformats this to multi-line, moving the nolint to the closing `)` line. gosec reports on the `json.Marshal(` line (line with the function call), not the `)` line.

**Attempted Solution 2:** Place nolint on the closing `)` line after gofumpt reformats

```go
requestDataEncoded, err := json.Marshal(
    requestData,
) //nolint:gosec // password is encrypted
```

- **Why it failed:** gosec reports the issue on the `json.Marshal(` line, not the `)` line. The nolint directive must be on the same line as the diagnostic.

## Solution

Place the `//nolint:gosec` comment on the `json.Marshal(` line itself, which is where gosec reports the issue, and where gofumpt will leave it:

```go
// Before (nolint on wrong line after gofumpt):
requestDataEncoded, err := json.Marshal(
    requestData,
) //nolint:gosec // G117: password is encrypted via BlowfishCall + HTTPS

// After (nolint on the line gosec actually reports):
requestDataEncoded, err := json.Marshal( //nolint:gosec // G117: password is encrypted via BlowfishCall + HTTPS
    requestData,
)
```

## Why This Works

1. **gosec reports on the function call line** (`json.Marshal(`), not the closing paren
2. **gofumpt preserves trailing comments** on the function call line when it reformats
3. **nolint directives are line-scoped** in golangci-lint -- they only suppress diagnostics reported on that exact line
4. When gofumpt splits a call across lines, the nolint must be on the line where the linter reports the issue, which is always the line containing the function name

## Prevention

- When adding `//nolint` to a function call that gofumpt might reformat, always place the directive on the line with the function name (`json.Marshal(`)
- Run `golangci-lint run --fix ./...` first to let gofumpt reformat, then add the nolint on the correct line
- Test with `git add -A && git commit` to verify nolint survives the pre-commit hook cycle

## Related Issues

No related issues documented yet.
