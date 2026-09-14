# Methodology

## Purpose

Aegis answers a narrow, useful question: **is the documented evidence for an AI system sufficiently complete, current, approved, and traceable for the next independent review?** It does not answer whether a system is lawful, safe in every context, certified, or compliant.

## Assessment Pipeline

1. A system owner declares purpose, role, domain, lifecycle stage, geography, capabilities, risk factors, data sensitivity, and affected groups.
2. The routing engine assigns `prohibited-review`, `high`, `transparency`, or `minimal` using explicit rules and records its reasons.
3. Aegis loads every policy pack, validates its structure, and records its version and SHA-256 fingerprint.
4. Evidence paths are constrained to the bundle directory. Every artefact is hashed and any supplied digest must match.
5. Applicable controls are evaluated against evidence type, status, integrity, minimum count, and validity date.
6. A decision gate is calculated, a typed evidence graph is assembled, and findings are exported to JSON, Markdown, and SARIF.
7. The canonical assessment payload is appended to a SQLite hash chain.

## Evidence Semantics

Only an item with `status: approved` and `integrity: verified` can satisfy a requirement. Draft or in-review material produces `partial`; a matching artefact outside its declared validity window produces `expired`; absence produces `missing`.

An evidence artefact may support more than one requirement when that relationship is defensible. Aegis preserves every link in the evidence graph rather than hiding reuse.

## Coverage and Decision Gate

Evidence coverage is:

```text
supported applicable controls / all applicable controls × 100
```

It is intentionally not weighted, not a maturity score, and not a compliance percentage. Any unsupported high- or critical-severity control produces `hold`. Lower-severity gaps produce `conditional-review`. Complete evidence produces `ready-for-independent-review`, never “compliant.”

## Determinism

Given the same system, evidence artefacts, policy packs, and timestamp, Aegis produces the same assessment identifier and substantive result. The demonstration fixes the timestamp at `2026-09-13T12:00:00Z` so repository outputs can be rebuilt and reviewed.

## Human Review Boundary

Policy interpretation, legal applicability, evidence quality, control effectiveness, stakeholder participation, proportionality, residual risk, and release approval remain human decisions. Aegis makes gaps visible; it does not make those decisions legitimate by itself.
