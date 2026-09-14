# Policy Sources and Versioning

Policy packs are maintained as readable JSON under [`policies/`](../policies/). Each pack contains an identifier, interpretation version, “as of” date, official source URL, disclaimer, applicability rules, and evidence requirements. The complete file is SHA-256 fingerprinted in every assessment.

## EU AI Act

- Official source: [European Commission regulatory framework for AI](https://digital-strategy.ec.europa.eu/en/policies/regulatory-framework-ai)
- Pack: [`eu-ai-act-2026.json`](../policies/eu-ai-act-2026.json)
- Scope: selected public concepts relevant to prohibited-practice escalation, high-risk system controls, and transparency.
- Boundary: this is an engineering triage aid, not a legal determination of role, classification, obligation, conformity, or compliance.

## NIST AI Risk Management Framework

- Official source: [NIST AI Risk Management Framework](https://www.nist.gov/itl/ai-risk-management-framework)
- Pack: [`nist-ai-rmf-1.0.json`](../policies/nist-ai-rmf-1.0.json)
- Scope: original evidence checks mapped to the GOVERN, MAP, MEASURE, and MANAGE functions, plus generative-AI provenance and adversarial evaluation.
- Boundary: the pack is not produced, assessed, or endorsed by NIST. NIST has announced revision work, so the pinned version and date matter.

## ISO/IEC 42001

- Official source: [ISO/IEC 42001 overview](https://www.iso.org/standard/42001)
- Pack: [`iso-42001-alignment.json`](../policies/iso-42001-alignment.json)
- Scope: high-level, original alignment prompts for policy, accountability, risk, lifecycle, suppliers, monitoring, and improvement.
- Boundary: the repository does not reproduce proprietary ISO requirements, replace licensed access to the standard, perform a conformity assessment, or confer certification.

## Change Procedure

1. Record the official source and date reviewed.
2. Change the pack version whenever meaning or evidence requirements change.
3. Add tests and a case proving the new applicability path.
4. Rebuild committed outputs and inspect drift.
5. Request legal, risk, and domain review before organisational use.

Historical policy versions should remain addressable for assessment replay. A production implementation should add signed releases and an approved policy change log.
