# Aegis Continuous-Assurance Report

- **System:** Talent Match Assistant (`talent-match-assistant`, version `1.0`)
- **Assessment:** `asg-a48809c62b6df35e`
- **Generated:** 2026-09-13T12:00:00Z
- **Decision gate:** **HOLD**

> **Important:** Aegis reports evidence readiness and traceability. It is not a legal opinion, certification, or determination of regulatory compliance.

## Executive Summary

Aegis routed this system to the **high** governance tier. 8 of 15 applicable controls are fully supported by verified and approved evidence (**53.3% evidence coverage**). The evidence review found 5 partial, 1 missing, and 1 expired control outcomes.

- Domain "employment" is routed to the high-impact control set.

## Decision Gate

| Applicable | Supported | Partial | Missing | Expired | Evidence Coverage |
| ---: | ---: | ---: | ---: | ---: | ---: |
| 15 | 8 | 5 | 1 | 1 | 53.3% |

## System Context

| Attribute | Value |
| --- | --- |
| Owner | People Analytics |
| Role | provider |
| Domain | employment |
| Lifecycle stage | pilot |
| Intended purpose | Prioritise job applications for trained recruiter review; never issue an autonomous employment decision. |
| Decision impact | Influences access to employment interviews. |
| Geographies | EU, UK |
| Capabilities | ranking, automated_recommendation |
| Affected groups | job applicants, recruiters |

## Framework Coverage

| Framework | Version | As Of | Applicable | Supported | Gaps | Pack SHA-256 |
| --- | --- | --- | ---: | ---: | ---: | --- |
| [EU AI Act](https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai) | 2026.09 | 2026-09-13 | 7 | 2 | 5 | `f710a9889c9c…` |
| [ISO/IEC 42001 Alignment](https://www.iso.org/standard/42001) | 2023-high-level | 2026-09-13 | 4 | 3 | 1 | `8ff9afd356a3…` |
| [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) | 1.0 | 2026-09-13 | 4 | 3 | 1 | `994a503d302f…` |

## Priority Findings

| Severity | Framework | Control | Status | Finding |
| --- | --- | --- | --- | --- |
| HIGH | EU AI Act | AIA-09 | partial | At least one requirement has draft or insufficient approved evidence. |
| HIGH | EU AI Act | AIA-10 | partial | At least one requirement has draft or insufficient approved evidence. |
| HIGH | EU AI Act | AIA-12 | partial | At least one requirement has draft or insufficient approved evidence. |
| HIGH | EU AI Act | AIA-13 | missing | At least one requirement has no matching evidence. |
| HIGH | EU AI Act | AIA-15 | expired | At least one requirement relies only on expired evidence. |
| HIGH | NIST AI RMF | MANAGE-1 | partial | At least one requirement has draft or insufficient approved evidence. |
| MEDIUM | ISO/IEC 42001 Alignment | AIMS-IMPROVE | partial | At least one requirement has draft or insufficient approved evidence. |

## Detailed Control Matrix

| Framework | Control | Source | Severity | Status | Evidence |
| --- | --- | --- | --- | --- | --- |
| EU AI Act | AIA-09 — Lifecycle Risk Management | Article 9 | high | **partial** | EV-003, EV-004 |
| EU AI Act | AIA-10 — Data Governance | Article 10 | high | **partial** | EV-005, EV-006 |
| EU AI Act | AIA-11 — Technical Documentation | Article 11 | high | **supported** | EV-008 |
| EU AI Act | AIA-12 — Event Logging | Article 12 | high | **partial** | EV-009 |
| EU AI Act | AIA-13 — Instructions and Transparency | Article 13 | high | **missing** | — |
| EU AI Act | AIA-14 — Human Oversight | Article 14 | high | **supported** | EV-010 |
| EU AI Act | AIA-15 — Accuracy, Robustness, and Cybersecurity | Article 15 | high | **expired** | EV-007, EV-011 |
| ISO/IEC 42001 Alignment | AIMS-LEAD — AI Policy and Accountability | High-level leadership alignment | high | **supported** | EV-001 |
| ISO/IEC 42001 Alignment | AIMS-RISK — Impact and Risk Process | High-level planning alignment | high | **supported** | EV-002, EV-003, EV-004 |
| ISO/IEC 42001 Alignment | AIMS-LIFE — Lifecycle and Supplier Governance | High-level operational alignment | medium | **supported** | EV-013 |
| ISO/IEC 42001 Alignment | AIMS-IMPROVE — Monitoring and Continual Improvement | High-level performance and improvement alignment | medium | **partial** | EV-012 |
| NIST AI RMF | GOVERN-1 — Accountability and Governance | GOVERN function | high | **supported** | EV-001 |
| NIST AI RMF | MAP-1 — Context and Impact Mapping | MAP function | high | **supported** | EV-002 |
| NIST AI RMF | MEASURE-1 — Independent Measurement | MEASURE function | high | **supported** | EV-006, EV-007 |
| NIST AI RMF | MANAGE-1 — Risk Treatment and Monitoring | MANAGE function | high | **partial** | EV-004, EV-012 |

## Evidence Lineage

| Evidence | Type | Status | Owner | Integrity | SHA-256 |
| --- | --- | --- | --- | --- | --- |
| EV-001 — Accountability Matrix | accountability_matrix | approved | AI Governance Council | verified | `4ce71f923fe4…` |
| EV-002 — Applicant Impact Assessment | impact_assessment | approved | Responsible AI Lead | verified | `222dadfecb46…` |
| EV-003 — Initial Risk Assessment | risk_assessment | approved | Enterprise Risk | verified | `569ad65a77b9…` |
| EV-004 — Risk Treatment Plan | risk_treatment | draft | Product Owner | verified | `a031c4fdfd6b…` |
| EV-005 — Training Data Card | data_card | approved | Data Steward | verified | `769842f7df72…` |
| EV-006 — Subgroup Evaluation | bias_test | in-review | Model Validation | verified | `ca697f145833…` |
| EV-007 — Performance Evaluation | performance_test | approved | Model Validation | verified | `ca697f145833…` |
| EV-008 — Model Card | model_card | approved | ML Engineering | verified | `3169da19db3a…` |
| EV-009 — Decision Log Specification | logging_specification | draft | Platform Engineering | verified | `424c277b5ec9…` |
| EV-010 — Recruiter Oversight Procedure | human_oversight_plan | approved | Recruiting Operations | verified | `cf5afa8c30b0…` |
| EV-011 — Security Assessment | security_assessment | approved | Information Security | verified | `2ebcaaa53ea8…` |
| EV-012 — Production Monitoring Plan | monitoring_plan | draft | ML Operations | verified | `424c277b5ec9…` |
| EV-013 — Release Change Record | change_record | approved | Change Advisory Board | verified | `09b72fb5c150…` |

## Evidence Graph and Audit Receipt

The evidence graph contains **39 nodes** and **46 typed edges**, linking the system to framework versions, controls, evidence artefacts, and findings.

- Ledger sequence: `1`
- Event hash: `59c75c08ba2f25cba9963ad09943dc761818638c48ddacceaef9d7dcc09f9b8f`
- Previous hash: `GENESIS`
- Payload hash: `ca6279982a755f53cf854b8746a49adc96bda81e0aff916fe2ac95230e980310`

## Interpretation and Limitations

Evidence coverage is the percentage of applicable controls supported by at least the configured minimum of verified, approved artefacts. It is not a compliance score. Applicability rules are a documented triage aid and must be reviewed by qualified legal, risk, security, and domain specialists. Framework mappings are versioned because regulatory and standards guidance changes over time.
