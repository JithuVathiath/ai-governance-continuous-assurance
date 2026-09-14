# Security Policy

## Supported Version

Security fixes are applied to the latest `0.1.x` release and the `main` branch.

## Reporting a Vulnerability

Please use GitHub's private vulnerability reporting feature for this repository. Do not open a public issue containing exploit details, confidential evidence, credentials, or personal data.

Include the affected version, impact, reproduction steps, and any suggested mitigation. Expect acknowledgement within seven days. Coordinated disclosure is appreciated.

## Scope Note

The demonstration ledger detects modification but is not immutable against an attacker who can rewrite the entire database and all external references. Production deployments should externally anchor signed ledger heads, enforce identity and access control, encrypt sensitive material, and isolate evidence stores.
