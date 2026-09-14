# ADR 0001: Report Evidence Readiness, Not Compliance

- Status: Accepted
- Date: 2026-09-13

## Context

Governance tools often compress legal interpretation, control design, evidence quality, and operating effectiveness into a single score. That presentation is precise-looking but unsafe: the dimensions require different expertise and cannot be established by file presence alone.

## Decision

Aegis reports the proportion of applicable controls supported by verified, approved evidence. It calls this **evidence coverage**. High-severity gaps trigger a hold, but complete coverage only means “ready for independent review.”

## Consequences

The output is less marketable than a compliance badge but more defensible. Organisations must define actual legal applicability, validate evidence quality, test control effectiveness, and make accountable release decisions outside the engine.
