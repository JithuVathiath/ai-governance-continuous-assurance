# Contributing

Contributions are welcome when they keep the project evidence-led, reviewable, and honest about its limits.

## Development

1. Use Go 1.27 or later.
2. Create a focused branch and keep generated evidence free of personal, confidential, or proprietary material.
3. Run `make verify` before opening a pull request.
4. Explain policy interpretation changes, cite the official public source, increment the pack version, and add a case or test covering the change.
5. Commit rebuilt files under `artifacts/public/` and `reports/` when assessment semantics change.

## Pull Request Checklist

- Tests cover success and failure paths.
- Public outputs are reproducible.
- New evidence paths cannot escape their bundle directory.
- Documentation distinguishes evidence readiness from legal compliance.
- No copyrighted standards text, secrets, personal data, or client material is included.
- Breaking changes and policy-pack changes are recorded in `CHANGELOG.md`.

By participating, contributors agree to respectful, evidence-based collaboration.
