# Threat Model

## Assets

- System declarations and evidence metadata
- Evidence artefacts and their provenance
- Policy-pack integrity and interpretation history
- Assessment results, findings, and decision gates
- Audit-ledger continuity

## Trust Boundaries

System owners supply declarations and evidence status. Repository maintainers supply policy mappings. Aegis verifies file integrity and deterministic evaluation, but it cannot prove that an artefact is truthful, sufficient, independently produced, or correctly approved.

## Key Threats and Controls

| Threat | Current Control | Residual Risk |
| --- | --- | --- |
| Evidence file replaced after registration | SHA-256 recomputed at assessment time; mismatched declared hashes fail closed | A malicious owner can submit a different artefact and label it `auto` |
| Path or symbolic-link traversal from an evidence bundle | Canonical resolved artefact paths must remain inside the canonical bundle directory | Filesystem replacement races require a hardened evidence store in production |
| Past assessment modified | Canonical payload hash plus chained event hashes; full-chain verifier | An attacker with database write access can rewrite the entire chain unless the head is externally anchored |
| Policy silently changed | Policy-pack hash and version stored with each assessment | Governance of who approves pack changes is organisational |
| False “compliance score” interpretation | Language and outputs say evidence coverage and independent review | Users may still over-trust automation |
| Deliberately understated system declaration | Deterministic mismatch detection for declared tiers; traceable inputs | Undeclared capabilities require discovery and stakeholder review |
| Spreadsheet/checklist gaming | Minimum evidence status, integrity, validity, and lineage checks | Effective control operation requires sampling and assurance work |
| Sensitive evidence leakage | Demonstrations are synthetic; artefacts remain local | Production deployments need encryption, access control, retention, redaction, and secrets management |

## Production Hardening Priorities

Externally anchor ledger heads, sign policy packs and evidence manifests, resolve symbolic links before path validation, use strong identity and role-based approval, encrypt data at rest, isolate tenants, verify signatures, record reviewer identity, and integrate evidence from authoritative systems rather than self-attestation alone.

Report security issues using [`SECURITY.md`](../SECURITY.md).
