# Aegis Continuous-Assurance Report

- **System:** Talent Match Assistant (`talent-match-assistant`, version `2.0`)
- **Assessment:** `asg-d72abd990536049a`
- **Generated:** 2026-09-13T12:00:00Z
- **Decision gate:** **HOLD**

> **Important:** Aegis reports evidence readiness and traceability. It is not a legal opinion, certification, or determination of regulatory compliance.

## Executive Summary

Aegis routed this system to the **high** governance tier. 14 of 17 applicable controls are fully supported by verified and approved evidence (**82.4% evidence coverage**). The evidence review found 1 partial, 2 missing, and 0 expired control outcomes.

- Domain "employment" is routed to the high-impact control set.

## Decision Gate

| Applicable | Supported | Partial | Missing | Expired | Evidence Coverage |
| ---: | ---: | ---: | ---: | ---: | ---: |
| 17 | 14 | 1 | 2 | 0 | 82.4% |

## System Context

| Attribute | Value |
| --- | --- |
| Owner | People Analytics |
| Role | provider |
| Domain | employment |
| Lifecycle stage | production |
| Intended purpose | Prioritise applications for recruiter review and draft evidence-grounded recommendation explanations. |
| Decision impact | Influences access to employment interviews. |
| Geographies | EU, UK |
| Capabilities | ranking, automated_recommendation, content_generation |
| Affected groups | job applicants, recruiters |

## Framework Coverage

| Framework | Version | As Of | Applicable | Supported | Gaps | Pack SHA-256 |
| --- | --- | --- | ---: | ---: | ---: | --- |
| [EU AI Act](https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai) | 2026.09 | 2026-09-13 | 8 | 7 | 1 | `f710a9889c9c…` |
| [ISO/IEC 42001 Alignment](https://www.iso.org/standard/42001) | 2023-high-level | 2026-09-13 | 4 | 4 | 0 | `8ff9afd356a3…` |
| [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) | 1.0 | 2026-09-13 | 5 | 3 | 2 | `994a503d302f…` |

## Priority Findings

| Severity | Framework | Control | Status | Finding |
| --- | --- | --- | --- | --- |
| HIGH | NIST AI RMF | MAP-1 | partial | At least one requirement has draft or insufficient approved evidence. |
| HIGH | NIST AI RMF | GAI-1 | missing | At least one requirement has no matching evidence. |
| MEDIUM | EU AI Act | AIA-50-C | missing | At least one requirement has no matching evidence. |

## Detailed Control Matrix

| Framework | Control | Source | Severity | Status | Evidence |
| --- | --- | --- | --- | --- | --- |
| EU AI Act | AIA-09 — Lifecycle Risk Management | Article 9 | high | **supported** | EV2-003, EV2-004 |
| EU AI Act | AIA-10 — Data Governance | Article 10 | high | **supported** | EV2-005, EV2-006 |
| EU AI Act | AIA-11 — Technical Documentation | Article 11 | high | **supported** | EV2-008 |
| EU AI Act | AIA-12 — Event Logging | Article 12 | high | **supported** | EV2-009 |
| EU AI Act | AIA-13 — Instructions and Transparency | Article 13 | high | **supported** | EV2-010 |
| EU AI Act | AIA-14 — Human Oversight | Article 14 | high | **supported** | EV2-011 |
| EU AI Act | AIA-15 — Accuracy, Robustness, and Cybersecurity | Article 15 | high | **supported** | EV2-007, EV2-012 |
| EU AI Act | AIA-50-C — AI-Generated Content Provenance | Article 50 | medium | **missing** | — |
| ISO/IEC 42001 Alignment | AIMS-LEAD — AI Policy and Accountability | High-level leadership alignment | high | **supported** | EV2-001 |
| ISO/IEC 42001 Alignment | AIMS-RISK — Impact and Risk Process | High-level planning alignment | high | **supported** | EV2-002, EV2-003, EV2-004 |
| ISO/IEC 42001 Alignment | AIMS-LIFE — Lifecycle and Supplier Governance | High-level operational alignment | medium | **supported** | EV2-008, EV2-015 |
| ISO/IEC 42001 Alignment | AIMS-IMPROVE — Monitoring and Continual Improvement | High-level performance and improvement alignment | medium | **supported** | EV2-013 |
| NIST AI RMF | GOVERN-1 — Accountability and Governance | GOVERN function | high | **supported** | EV2-001 |
| NIST AI RMF | MAP-1 — Context and Impact Mapping | MAP function | high | **partial** | EV2-002 |
| NIST AI RMF | MEASURE-1 — Independent Measurement | MEASURE function | high | **supported** | EV2-006, EV2-007, EV2-014 |
| NIST AI RMF | MANAGE-1 — Risk Treatment and Monitoring | MANAGE function | high | **supported** | EV2-004, EV2-013 |
| NIST AI RMF | GAI-1 — Generative-AI Content Provenance | Generative AI Profile | high | **missing** | EV2-014 |

## Evidence Lineage

| Evidence | Type | Status | Owner | Integrity | SHA-256 |
| --- | --- | --- | --- | --- | --- |
| EV2-001 — Updated Accountability Matrix | accountability_matrix | approved | AI Governance Council | verified | `dc63e99fe1b5…` |
| EV2-002 — Generative Feature Impact Addendum | impact_assessment | in-review | Responsible AI Lead | verified | `29634d6a3cbe…` |
| EV2-003 — Version 2 Risk Assessment | risk_assessment | approved | Enterprise Risk | verified | `26d5f36c1f9c…` |
| EV2-004 — Approved Risk Treatment Plan | risk_treatment | approved | AI Governance Council | verified | `26d5f36c1f9c…` |
| EV2-005 — Versioned Data Card | data_card | approved | Data Steward | verified | `a182619916b3…` |
| EV2-006 — Approved Subgroup Evaluation | bias_test | approved | Model Validation | verified | `a182619916b3…` |
| EV2-007 — Version 2 Performance Evaluation | performance_test | approved | Model Validation | verified | `a182619916b3…` |
| EV2-008 — Version 2 Technical File | technical_documentation | approved | ML Engineering | verified | `59f194a18daa…` |
| EV2-009 — Decision and Prompt Log Specification | logging_specification | approved | Platform Engineering | verified | `f51e8da1505a…` |
| EV2-010 — Recruiter Instructions | instructions_for_use | approved | Recruiting Operations | verified | `c7d4ce45777f…` |
| EV2-011 — Updated Human Oversight Procedure | human_oversight_plan | approved | Recruiting Operations | verified | `c7d4ce45777f…` |
| EV2-012 — Current Security Assessment | security_assessment | approved | Information Security | verified | `c490a5fc0c09…` |
| EV2-013 — Production Monitoring and Response | monitoring_plan | approved | ML Operations | verified | `f51e8da1505a…` |
| EV2-014 — Explanation Red-Team Report | red_team_report | approved | AI Safety | verified | `b1199518f6f3…` |
| EV2-015 — Foundation Model Supplier Review | supplier_assessment | approved | Third-Party Risk | verified | `583464f89edc…` |

## Evidence Graph and Audit Receipt

The evidence graph contains **39 nodes** and **48 typed edges**, linking the system to framework versions, controls, evidence artefacts, and findings.

- Ledger sequence: `2`
- Event hash: `c40b244dcf2a4216a9a0c29982d8043326e69772a65a8fa9299c6b8f19505a6c`
- Previous hash: `59c75c08ba2f25cba9963ad09943dc761818638c48ddacceaef9d7dcc09f9b8f`
- Payload hash: `c9aff2876ec6f17ac5597bf0b49e3c4d1cbcb9d7491e2f06c7bdc144701981dd`

## Interpretation and Limitations

Evidence coverage is the percentage of applicable controls supported by at least the configured minimum of verified, approved artefacts. It is not a compliance score. Applicability rules are a documented triage aid and must be reviewed by qualified legal, risk, security, and domain specialists. Framework mappings are versioned because regulatory and standards guidance changes over time.
