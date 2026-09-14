# Aegis Continuous-Assurance Report

- **System:** Consumer Credit Access Score (`credit-access-score`, version `3.2`)
- **Assessment:** `asg-2e27729303f3427e`
- **Generated:** 2026-09-13T12:00:00Z
- **Decision gate:** **HOLD**

> **Important:** Aegis reports evidence readiness and traceability. It is not a legal opinion, certification, or determination of regulatory compliance.

## Executive Summary

Aegis routed this system to the **high** governance tier. 14 of 15 applicable controls are fully supported by verified and approved evidence (**93.3% evidence coverage**). The evidence review found 1 partial, 0 missing, and 0 expired control outcomes.

- Domain "creditworthiness" is routed to the high-impact control set.

## Decision Gate

| Applicable | Supported | Partial | Missing | Expired | Evidence Coverage |
| ---: | ---: | ---: | ---: | ---: | ---: |
| 15 | 14 | 1 | 0 | 0 | 93.3% |

## System Context

| Attribute | Value |
| --- | --- |
| Owner | Retail Credit Risk |
| Role | provider |
| Domain | creditworthiness |
| Lifecycle stage | production |
| Intended purpose | Recommend credit application risk bands for a trained underwriter who makes the final decision. |
| Decision impact | Influences access to consumer credit and pricing review. |
| Geographies | EU |
| Capabilities | classification, automated_recommendation |
| Affected groups | credit applicants, underwriters |

## Framework Coverage

| Framework | Version | As Of | Applicable | Supported | Gaps | Pack SHA-256 |
| --- | --- | --- | ---: | ---: | ---: | --- |
| [EU AI Act](https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai) | 2026.09 | 2026-09-13 | 7 | 6 | 1 | `f710a9889c9c…` |
| [ISO/IEC 42001 Alignment](https://www.iso.org/standard/42001) | 2023-high-level | 2026-09-13 | 4 | 4 | 0 | `8ff9afd356a3…` |
| [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) | 1.0 | 2026-09-13 | 4 | 4 | 0 | `994a503d302f…` |

## Priority Findings

| Severity | Framework | Control | Status | Finding |
| --- | --- | --- | --- | --- |
| HIGH | EU AI Act | AIA-10 | partial | At least one requirement has draft or insufficient approved evidence. |

## Detailed Control Matrix

| Framework | Control | Source | Severity | Status | Evidence |
| --- | --- | --- | --- | --- | --- |
| EU AI Act | AIA-09 — Lifecycle Risk Management | Article 9 | high | **supported** | CR-003, CR-004 |
| EU AI Act | AIA-10 — Data Governance | Article 10 | high | **partial** | CR-005, CR-006 |
| EU AI Act | AIA-11 — Technical Documentation | Article 11 | high | **supported** | CR-008 |
| EU AI Act | AIA-12 — Event Logging | Article 12 | high | **supported** | CR-009 |
| EU AI Act | AIA-13 — Instructions and Transparency | Article 13 | high | **supported** | CR-010 |
| EU AI Act | AIA-14 — Human Oversight | Article 14 | high | **supported** | CR-011 |
| EU AI Act | AIA-15 — Accuracy, Robustness, and Cybersecurity | Article 15 | high | **supported** | CR-007, CR-012 |
| ISO/IEC 42001 Alignment | AIMS-LEAD — AI Policy and Accountability | High-level leadership alignment | high | **supported** | CR-001 |
| ISO/IEC 42001 Alignment | AIMS-RISK — Impact and Risk Process | High-level planning alignment | high | **supported** | CR-002, CR-003, CR-004 |
| ISO/IEC 42001 Alignment | AIMS-LIFE — Lifecycle and Supplier Governance | High-level operational alignment | medium | **supported** | CR-008 |
| ISO/IEC 42001 Alignment | AIMS-IMPROVE — Monitoring and Continual Improvement | High-level performance and improvement alignment | medium | **supported** | CR-013, CR-014 |
| NIST AI RMF | GOVERN-1 — Accountability and Governance | GOVERN function | high | **supported** | CR-001 |
| NIST AI RMF | MAP-1 — Context and Impact Mapping | MAP function | high | **supported** | CR-002 |
| NIST AI RMF | MEASURE-1 — Independent Measurement | MEASURE function | high | **supported** | CR-006, CR-007 |
| NIST AI RMF | MANAGE-1 — Risk Treatment and Monitoring | MANAGE function | high | **supported** | CR-004, CR-013, CR-014 |

## Evidence Lineage

| Evidence | Type | Status | Owner | Integrity | SHA-256 |
| --- | --- | --- | --- | --- | --- |
| CR-001 — Credit AI Governance Policy | ai_policy | approved | Model Risk Committee | verified | `e14595dce401…` |
| CR-002 — Borrower Impact Assessment | impact_assessment | approved | Consumer Duty Lead | verified | `7c94dd4ee0f5…` |
| CR-003 — Model Risk Assessment | risk_assessment | approved | Independent Model Risk | verified | `3b309c01ca1e…` |
| CR-004 — Risk Treatment Record | risk_treatment | approved | Retail Credit Risk | verified | `3b309c01ca1e…` |
| CR-005 — Credit Data Card | data_card | approved | Credit Data Office | verified | `c7ee41c04f50…` |
| CR-006 — Fair-Lending Validation | bias_test | in-review | Independent Model Risk | verified | `7b321463c350…` |
| CR-007 — Performance and Calibration Test | performance_test | approved | Independent Model Risk | verified | `7b321463c350…` |
| CR-008 — Technical Model File | technical_documentation | approved | Model Development | verified | `9235e3d20f15…` |
| CR-009 — Decision Audit Sample | audit_log_sample | approved | Credit Operations | verified | `adcb1f66e6c1…` |
| CR-010 — Underwriter Operating Guide | instructions_for_use | approved | Credit Operations | verified | `ba4824760968…` |
| CR-011 — Manual Review and Appeal Process | appeals_process | approved | Customer Outcomes | verified | `ba4824760968…` |
| CR-012 — Security and Resilience Review | security_assessment | approved | Information Security | verified | `8b982b2c6f2d…` |
| CR-013 — Outcome Monitoring Plan | monitoring_plan | approved | Model Monitoring | verified | `adcb1f66e6c1…` |
| CR-014 — Model Incident Playbook | incident_response | approved | Operational Risk | verified | `adcb1f66e6c1…` |

## Evidence Graph and Audit Receipt

The evidence graph contains **34 nodes** and **43 typed edges**, linking the system to framework versions, controls, evidence artefacts, and findings.

- Ledger sequence: `3`
- Event hash: `733693ea344db8ada741ed7e420683a29963538f1e8319f164080a2dcd68f1f6`
- Previous hash: `c40b244dcf2a4216a9a0c29982d8043326e69772a65a8fa9299c6b8f19505a6c`
- Payload hash: `27dddeae794ce0d890d5b3e60dfd95e968c250f7f773671d66197e5013fc1684`

## Interpretation and Limitations

Evidence coverage is the percentage of applicable controls supported by at least the configured minimum of verified, approved artefacts. It is not a compliance score. Applicability rules are a documented triage aid and must be reviewed by qualified legal, risk, security, and domain specialists. Framework mappings are versioned because regulatory and standards guidance changes over time.
