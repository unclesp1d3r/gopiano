# Security Policy

## Supported Versions

The following versions of gopiano currently receive security updates:

| Version | Supported          |
| ------- | ------------------ |
| Latest release | :white_check_mark: |
| main branch | :white_check_mark: |
| Older releases | :x:                |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security vulnerability in gopiano, please report it responsibly.

### How to Report

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, please report them via one of the following methods:

1. **GitHub Security Advisories** (Preferred): Use the [Security tab](https://github.com/unclesp1d3r/gopiano/security/advisories) in this repository to create a private security advisory.

2. **Email**: Contact the maintainer directly (see repository maintainer information).

### What to Include

When reporting a vulnerability, please include:

- A description of the vulnerability
- Steps to reproduce the issue
- Potential impact
- Suggested fix (if any)

### Response Timeline

Please note that gopiano is maintained as a part-time hobby project by a single maintainer. While we take security seriously and will address all reported vulnerabilities, response times may be longer than those of commercial projects with dedicated security teams. We appreciate your patience and understanding.

- **Acknowledgment**: We will acknowledge receipt of your report within 1 week
- **Initial Assessment**: We will provide an initial assessment within 2-3 weeks
- **Resolution**: Critical vulnerabilities will be prioritized and addressed as quickly as possible, typically within 60-90 days depending on severity and complexity
- **Updates**: We will provide periodic updates on the status of the vulnerability, with more frequent communication for critical issues

## Security Update Process

1. **Assessment**: The maintainer will assess the reported vulnerability
2. **Development**: A security patch will be developed in a private branch
3. **Testing**: The patch will be thoroughly tested
4. **Release**: A security release will be published with appropriate versioning
5. **Disclosure**: After the fix is released, a security advisory will be published detailing the vulnerability (with appropriate delay for users to update)

## Scope

### In Scope

Security vulnerabilities in gopiano include, but are not limited to:

- **Credential Exposure**: Issues that could lead to exposure of user credentials or authentication tokens
- **Encryption Flaws**: Vulnerabilities in the Blowfish encryption/decryption implementation
- **Insecure HTTP Communication**: Issues related to insecure API communication
- **Input Validation**: Vulnerabilities that could lead to injection attacks or data corruption
- **Memory Safety**: Issues that could lead to buffer overflows or memory corruption
- **Dependency Vulnerabilities**: Known vulnerabilities in project dependencies

### Out of Scope

The following are considered out of scope for gopiano security reporting:

- **Pandora API Issues**: Vulnerabilities or issues with Pandora.com's API itself should be reported directly to Pandora
- **Feature Requests**: General feature requests or enhancements
- **Non-Security Bugs**: Regular bugs that don't have security implications
- **Denial of Service**: DoS issues that require excessive resources or are inherent to the API design

## Security Best Practices

When using gopiano:

- Always use the latest version
- Never commit credentials or authentication tokens to version control
- Use secure storage for user credentials
- Regularly update dependencies
- Review the code before using in production (especially given the alpha status)

## Acknowledgments

We appreciate the security research community's efforts to help keep gopiano secure. Security researchers who responsibly disclose vulnerabilities will be acknowledged (with their permission) in security advisories.
