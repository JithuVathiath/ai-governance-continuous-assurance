# Aegis Continuous-Assurance Report

- **System:** Enterprise Foundation Model Gateway (`enterprise-foundation-model-gateway`, version `0.8`)
- **Assessment:** `asg-00718728fb5eafbe`
- **Generated:** 2026-09-13T12:00:00Z
- **Decision gate:** **CONDITIONAL-REVIEW**

> **Important:** Aegis reports evidence readiness and traceability. It is not a legal opinion, certification, or determination of regulatory compliance.

## Executive Summary

Aegis routed this system to the **minimal** governance tier. 6 of 7 applicable controls are fully supported by verified and approved evidence (**85.7% evidence coverage**). The evidence review found 1 partial, 0 missing, and 0 expired control outcomes.

- No configured high-risk or transparency trigger matched.

## Decision Gate

| Applicable | Supported | Partial | Missing | Expired | Evidence Coverage |
| ---: | ---: | ---: | ---: | ---: | ---: |
| 7 | 6 | 1 | 0 | 0 | 85.7% |

## System Context

| Attribute | Value |
| --- | --- |
| Owner | Technology Strategy |
| Role | deployer |
| Domain | technology_procurement |
| Lifecycle stage | design |
| Intended purpose | Evaluate and govern access to third-party foundation models before downstream use cases are approved. |
| Decision impact | Controls technology selection; it does not make decisions about individuals. |
| Geographies | EU, UK, US |
| Capabilities | foundation_model_dependency, model_routing |
| Affected groups | employees, application owners |

## Framework Coverage

| Framework | Version | As Of | Applicable | Supported | Gaps | Pack SHA-256 |
| --- | --- | --- | ---: | ---: | ---: | --- |
| [EU AI Act](https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai) | 2026.09 | 2026-09-13 | 0 | 0 | 0 | `f710a9889c9c…` |
| [ISO/IEC 42001 Alignment](https://www.iso.org/standard/42001) | 2023-high-level | 2026-09-13 | 3 | 2 | 1 | `8ff9afd356a3…` |
| [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) | 1.0 | 2026-09-13 | 4 | 4 | 0 | `994a503d302f…` |

## Priority Findings

| Severity | Framework | Control | Status | Finding |
| --- | --- | --- | --- | --- |
| MEDIUM | ISO/IEC 42001 Alignment | AIMS-LIFE | partial | At least one requirement has draft or insufficient approved evidence. |

## Detailed Control Matrix

| Framework | Control | Source | Severity | Status | Evidence |
| --- | --- | --- | --- | --- | --- |
| ISO/IEC 42001 Alignment | AIMS-LEAD — AI Policy and Accountability | High-level leadership alignment | high | **supported** | FM-001 |
| ISO/IEC 42001 Alignment | AIMS-RISK — Impact and Risk Process | High-level planning alignment | high | **supported** | FM-002, FM-003, FM-004 |
| ISO/IEC 42001 Alignment | AIMS-LIFE — Lifecycle and Supplier Governance | High-level operational alignment | medium | **partial** | FM-008 |
| NIST AI RMF | GOVERN-1 — Accountability and Governance | GOVERN function | high | **supported** | FM-001 |
| NIST AI RMF | MAP-1 — Context and Impact Mapping | MAP function | high | **supported** | FM-002 |
| NIST AI RMF | MEASURE-1 — Independent Measurement | MEASURE function | high | **supported** | FM-005 |
| NIST AI RMF | MANAGE-1 — Risk Treatment and Monitoring | MANAGE function | high | **supported** | FM-004, FM-006, FM-007 |

## Evidence Lineage

| Evidence | Type | Status | Owner | Integrity | SHA-256 |
| --- | --- | --- | --- | --- | --- |
| FM-001 — Enterprise AI Use Policy | ai_policy | approved | Chief Technology Office | verified | `5ace935026fe…` |
| FM-002 — Gateway Impact Assessment | impact_assessment | approved | Responsible AI | verified | `c2ed7a8eddce…` |
| FM-003 — Third-Party Model Risk Assessment | risk_assessment | approved | Technology Risk | verified | `c2ed7a8eddce…` |
| FM-004 — Gateway Risk Treatment Plan | risk_treatment | approved | Technology Strategy | verified | `c2ed7a8eddce…` |
| FM-005 — Comparative Safety Evaluation | red_team_report | approved | AI Security | verified | `ad1f00bcc586…` |
| FM-006 — Supplier and Model Change Monitoring | monitoring_plan | approved | AI Platform | verified | `1011d831530c…` |
| FM-007 — Gateway Incident Playbook | incident_response | approved | AI Platform | verified | `1011d831530c…` |
| FM-008 — Candidate Supplier Due Diligence | supplier_assessment | in-review | Third-Party Risk | verified | `2b1dd6bce1ab…` |

## Evidence Graph and Audit Receipt

The evidence graph contains **20 nodes** and **22 typed edges**, linking the system to framework versions, controls, evidence artefacts, and findings.

- Ledger sequence: `5`
- Event hash: `957903985af2cf0db265e79174270c7266efd48595067021140120eeff018262`
- Previous hash: `8b69c3d65c5a155aaa3dd011e805557133213828038f8f08335d7c54d9aa3d74`
- Payload hash: `08e4623fe77ce4f9e74d432762737131ee96e9bcca8cea3107277e99553eda60`

## Interpretation and Limitations

Evidence coverage is the percentage of applicable controls supported by at least the configured minimum of verified, approved artefacts. It is not a compliance score. Applicability rules are a documented triage aid and must be reviewed by qualified legal, risk, security, and domain specialists. Framework mappings are versioned because regulatory and standards guidance changes over time.
