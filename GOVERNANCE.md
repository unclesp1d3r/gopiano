# Project Governance

This document outlines how the gopiano project is managed, how decisions are made, and how the community can participate in the project's development.

## Project Maintainers

### Current Maintainer

- **UncleSp1d3r** - Current project maintainer

### Original Author

- **Joshua Gardner** ([@cellofellow](https://github.com/cellofellow)) - Original creator and author of gopiano (2014)

### Maintainer Responsibilities

Maintainers are responsible for:

- **Code Review**: Reviewing and merging pull requests
- **Release Management**: Managing releases, versioning, and changelogs
- **Issue Triage**: Triaging issues, labeling, and prioritizing
- **Community Management**: Enforcing the Code of Conduct, responding to questions
- **Project Direction**: Making decisions about project direction and major changes
- **Documentation**: Ensuring documentation is accurate and up-to-date

## Contribution Process

### Workflow

1. **Fork** the repository
2. **Create a branch** with a descriptive name (e.g., `feature/add-new-method`, `fix/encryption-bug`)
3. **Make changes** following project conventions and standards
4. **Test** your changes (run `go test ./...` and `golangci-lint run ./...`)
5. **Commit** with clear, descriptive commit messages
6. **Push** to your fork
7. **Create a Pull Request** with a clear description of changes
8. **Review** - Maintainers will review the PR
9. **Merge** - Once approved, the PR will be merged

For detailed contribution guidelines, see [CONTRIBUTING.md](CONTRIBUTING.md).

### Review Requirements

- At least one maintainer approval is required for merging pull requests
- All CI checks must pass (tests, linting, etc.)
- Code must follow project conventions and standards
- Documentation should be updated if adding new features

## Decision Making

### Small Changes

For small changes (bug fixes, documentation improvements, minor features):

- Maintainers have discretion to make decisions
- Pull requests can be merged after review and approval
- No community discussion required

### Major Changes

For major changes (API changes, architectural changes, breaking changes):

- Open an issue for discussion before implementing
- Gather community feedback
- Maintainers make final decisions based on:
  - Technical merit
  - Alignment with project goals
  - Community feedback
  - Maintenance burden

### Disagreements

If there's disagreement about a decision:

1. Discuss in the relevant issue or pull request
2. Maintainers will consider all viewpoints
3. Maintainers make the final decision
4. Decisions will be documented and explained

## Release Management

### Versioning

gopiano follows [Semantic Versioning](https://semver.org/) (SemVer):

- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions
- **PATCH** version for backwards-compatible bug fixes

### Release Process

1. **Testing**: Ensure all tests pass and code is linted
2. **Changelog**: Update changelog/release notes
3. **Version Tag**: Create a git tag with the version number (e.g., `v1.0.0`)
4. **Release Notes**: Create a GitHub release with release notes
5. **Announcement**: Announce the release (if applicable)

### Current Status

The project is currently in **alpha quality** (as noted in the README). This means:

- The API may change
- There may be bugs
- Proper tests and error handling are still being improved
- Use at your own risk

As the project matures, we aim to move toward beta and eventually stable releases.

## Adding Maintainers

As the project grows, additional maintainers may be added. Criteria for adding maintainers:

- **Consistent Contributions**: Demonstrated commitment through regular, quality contributions
- **Code Quality**: High-quality code that follows project standards
- **Community Engagement**: Active participation in issues, discussions, and reviews
- **Alignment**: Alignment with project goals and values
- **Consensus**: Agreement from existing maintainers

The process for adding maintainers:

1. Existing maintainers identify potential candidates
2. Discuss with the candidate
3. Invite the candidate to become a maintainer
4. Update this document with the new maintainer information

## Conflict Resolution

### Reporting Issues

If you have concerns about maintainer behavior or project decisions:

1. First, try to resolve through direct communication
2. If that's not possible, contact the maintainer privately
3. For Code of Conduct violations, see [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)

### Resolution Process

1. **Acknowledge**: The issue will be acknowledged promptly
2. **Investigate**: Gather information from all parties
3. **Discuss**: Discuss potential resolutions
4. **Decide**: Make a decision and communicate it clearly
5. **Document**: Document the resolution (if appropriate)

## Project Goals

The primary goals of gopiano are:

- **Thin Wrapper**: Provide a minimal, thin wrapper around Pandora's JSON API
- **API Fidelity**: Mirror the Pandora API exactly (not an abstraction)
- **Go Idioms**: Follow Go best practices and conventions
- **Reliability**: Improve test coverage and error handling
- **Documentation**: Maintain clear, accurate documentation

## Questions

If you have questions about governance, decision-making, or how to get involved:

- Open an issue for discussion
- Contact the maintainer directly
- See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines

## Changes to Governance

This governance document may evolve as the project grows. Changes will be:

- Discussed in issues or pull requests
- Approved by maintainers
- Documented in this file
