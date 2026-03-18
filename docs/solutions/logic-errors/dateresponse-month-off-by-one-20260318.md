---
module: System
date: 2026-03-18
problem_type: logic_error
component: tooling
symptoms:
  - All dates from Pandora API are one month early
  - January responses parse as invalid month 0
  - DateResponse.GetDate returns incorrect time.Time
root_cause: logic_error
resolution_type: code_fix
severity: high
tags: [date-parsing, java-convention, off-by-one, pandora-api, go]
---

# Troubleshooting: DateResponse.GetDate Returns Month Off-By-One

## Problem

`DateResponse.GetDate()` returns dates with the month shifted by -1 (e.g., February data returns January, January returns invalid month 0). All dates from the Pandora API are silently wrong.

## Environment

- Module: responses package (DateResponse)
- Go Version: 1.24+
- Affected Component: `responses/common.go` - `DateResponse.GetDate()`
- Date: 2026-03-18

## Symptoms

- Dates from Pandora API appear one month early
- January (month 0 from Java) creates an invalid Go `time.Time` (month 0 is not valid)
- No error is returned -- the bug is silent

## What Didn't Work

**Direct solution:** The bug was identified by CodeRabbit during PR review. It was present since the original code and never caught because there were no unit tests for `GetDate()`.

## Solution

The Pandora API returns dates using Java's deprecated `Date.getMonth()` convention where months are **0-indexed** (0=January, 11=December). Go's `time.Month` is **1-indexed** (1=January, 12=December). Add +1 to convert:

```go
// Before (broken - month off by one):
return time.Date(1900+d.Year, time.Month(d.Month), ...)

// After (fixed):
return time.Date(1900+d.Year, time.Month(d.Month+1), ...)
```

## Why This Works

Java's `Date.getMonth()` returns 0 for January, 1 for February, etc. Go's `time.Month` type uses 1 for January, 2 for February, etc. The Pandora API serializes the Java convention into JSON. Without the +1 adjustment, every month is mapped to the previous month in Go.

## Prevention

- When wrapping APIs that originate from Java, check for 0-indexed months and 1900-offset years
- Add unit tests for date conversion functions -- the `GetDate()` method had 0% test coverage
- Test with known date values from API documentation, not just structural tests

## Related Issues

- See also: [nolint-placement-with-gofumpt-System-20260317.md](../developer-experience/nolint-placement-with-gofumpt-System-20260317.md) - another non-obvious bug from the same session
