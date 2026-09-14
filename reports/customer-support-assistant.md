# Aegis Continuous-Assurance Report

- **System:** Customer Support Answer Assistant (`helpdesk-answer-assistant`, version `1.4`)
- **Assessment:** `asg-c95f1a09e734a74e`
- **Generated:** 2026-09-13T12:00:00Z
- **Decision gate:** **READY-FOR-INDEPENDENT-REVIEW**

> **Important:** Aegis reports evidence readiness and traceability. It is not a legal opinion, certification, or determination of regulatory compliance.

## Executive Summary

Aegis routed this system to the **transparency** governance tier. 11 of 11 applicable controls are fully supported by verified and approved evidence (**100.0% evidence coverage**). The evidence review found 0 partial, 0 missing, and 0 expired control outcomes.

- Capability "conversational_agent" triggers transparency controls.

## Decision Gate

| Applicable | Supported | Partial | Missing | Expired | Evidence Coverage |
| ---: | ---: | ---: | ---: | ---: | ---: |
| 11 | 11 | 0 | 0 | 0 | 100.0% |

## System Context

| Attribute | Value |
| --- | --- |
| Owner | Customer Experience |
| Role | deployer |
| Domain | customer_support |
| Lifecycle stage | production |
| Intended purpose | Draft grounded customer-support answers from an approved knowledge base with agent review for sensitive topics. |
| Decision impact | Provides service information but cannot approve claims, credit, or account restrictions. |
| Geographies | EU, UK |
| Capabilities | conversational_agent, content_generation, retrieval_augmented_generation |
| Affected groups | customers, support agents |

## Framework Coverage

| Framework | Version | As Of | Applicable | Supported | Gaps | Pack SHA-256 |
| --- | --- | --- | ---: | ---: | ---: | --- |
| [EU AI Act](https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai) | 2026.09 | 2026-09-13 | 2 | 2 | 0 | `f710a9889c9c…` |
| [ISO/IEC 42001 Alignment](https://www.iso.org/standard/42001) | 2023-high-level | 2026-09-13 | 4 | 4 | 0 | `8ff9afd356a3…` |
| [NIST AI RMF](https://www.nist.gov/itl/ai-risk-management-framework) | 1.0 | 2026-09-13 | 5 | 5 | 0 | `994a503d302f…` |

## Priority Findings

No evidence gaps were detected. Independent review is still required.

## Detailed Control Matrix

| Framework | Control | Source | Severity | Status | Evidence |
| --- | --- | --- | --- | --- | --- |
| EU AI Act | AIA-50-I — AI Interaction Transparency | Article 50 | medium | **supported** | CS-008 |
| EU AI Act | AIA-50-C — AI-Generated Content Provenance | Article 50 | medium | **supported** | CS-007, CS-008 |
| ISO/IEC 42001 Alignment | AIMS-LEAD — AI Policy and Accountability | High-level leadership alignment | high | **supported** | CS-001 |
| ISO/IEC 42001 Alignment | AIMS-RISK — Impact and Risk Process | High-level planning alignment | high | **supported** | CS-002, CS-003, CS-004 |
| ISO/IEC 42001 Alignment | AIMS-LIFE — Lifecycle and Supplier Governance | High-level operational alignment | medium | **supported** | CS-011 |
| ISO/IEC 42001 Alignment | AIMS-IMPROVE — Monitoring and Continual Improvement | High-level performance and improvement alignment | medium | **supported** | CS-009, CS-010 |
| NIST AI RMF | GOVERN-1 — Accountability and Governance | GOVERN function | high | **supported** | CS-001 |
| NIST AI RMF | MAP-1 — Context and Impact Mapping | MAP function | high | **supported** | CS-002 |
| NIST AI RMF | MEASURE-1 — Independent Measurement | MEASURE function | high | **supported** | CS-005, CS-006 |
| NIST AI RMF | MANAGE-1 — Risk Treatment and Monitoring | MANAGE function | high | **supported** | CS-004, CS-009, CS-010 |
| NIST AI RMF | GAI-1 — Generative-AI Content Provenance | Generative AI Profile | high | **supported** | CS-006, CS-007 |

## Evidence Lineage

| Evidence | Type | Status | Owner | Integrity | SHA-256 |
| --- | --- | --- | --- | --- | --- |
| CS-001 — Assistant Accountability Matrix | accountability_matrix | approved | AI Product Council | verified | `73ad9ec00e1c…` |
| CS-002 — Customer Impact Assessment | impact_assessment | approved | Customer Outcomes | verified | `5aecb4695411…` |
| CS-003 — Generative AI Risk Assessment | risk_assessment | approved | Responsible AI | verified | `5aecb4695411…` |
| CS-004 — Generative AI Treatment Plan | risk_treatment | approved | AI Product Council | verified | `5aecb4695411…` |
| CS-005 — Groundedness Evaluation | performance_test | approved | AI Quality | verified | `91cca1147a6f…` |
| CS-006 — Safety Red-Team Report | red_team_report | approved | AI Security | verified | `91cca1147a6f…` |
| CS-007 — Answer Provenance Design | content_provenance | approved | AI Engineering | verified | `e8ec36ebf4f1…` |
| CS-008 — Customer AI Disclosure | transparency_notice | approved | Customer Experience | verified | `e8ec36ebf4f1…` |
| CS-009 — Quality and Safety Monitoring | monitoring_plan | approved | AI Operations | verified | `8196993dd970…` |
| CS-010 — AI Incident Playbook | incident_response | approved | AI Operations | verified | `8196993dd970…` |
| CS-011 — Model Service Supplier Review | supplier_assessment | approved | Third-Party Risk | verified | `aea53cba29e1…` |

## Evidence Graph and Audit Receipt

The evidence graph contains **26 nodes** and **33 typed edges**, linking the system to framework versions, controls, evidence artefacts, and findings.

- Ledger sequence: `4`
- Event hash: `8b69c3d65c5a155aaa3dd011e805557133213828038f8f08335d7c54d9aa3d74`
- Previous hash: `733693ea344db8ada741ed7e420683a29963538f1e8319f164080a2dcd68f1f6`
- Payload hash: `e974673b82e0571bd00dbefbfd73cc8e009023e86cbac8785d6f249e7c5ef734`

## Interpretation and Limitations

Evidence coverage is the percentage of applicable controls supported by at least the configured minimum of verified, approved artefacts. It is not a compliance score. Applicability rules are a documented triage aid and must be reviewed by qualified legal, risk, security, and domain specialists. Framework mappings are versioned because regulatory and standards guidance changes over time.
