package report

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

// Write creates an executive and technical assurance report from an assessment.
func Write(path string, assessment domain.Assessment) error {
	var builder strings.Builder
	builder.WriteString("# Aegis Continuous-Assurance Report\n\n")
	fmt.Fprintf(&builder, "**System:** %s (`%s`, version `%s`)  \n", assessment.System.Name, assessment.System.ID, assessment.System.Version)
	fmt.Fprintf(&builder, "**Assessment:** `%s`  \n", assessment.AssessmentID)
	fmt.Fprintf(&builder, "**Generated:** %s  \n", assessment.GeneratedAt)
	fmt.Fprintf(&builder, "**Decision gate:** **%s**\n\n", strings.ToUpper(assessment.Summary.Decision))

	builder.WriteString("> **Important:** " + assessment.Disclaimer + "\n\n")
	builder.WriteString("## Executive Summary\n\n")
	fmt.Fprintf(&builder, "Aegis routed this system to the **%s** governance tier. ", assessment.Classification.Tier)
	fmt.Fprintf(&builder, "%d of %d applicable controls are fully supported by verified and approved evidence (**%.1f%% evidence coverage**). ",
		assessment.Summary.Supported, assessment.Summary.Applicable, assessment.Summary.EvidenceCoverage)
	fmt.Fprintf(&builder, "The evidence review found %d partial, %d missing, and %d expired control outcomes.\n\n",
		assessment.Summary.Partial, assessment.Summary.Missing, assessment.Summary.Expired)
	for _, reason := range assessment.Classification.Reasons {
		fmt.Fprintf(&builder, "- %s\n", reason)
	}
	if assessment.Classification.DeclarationMismatch {
		fmt.Fprintf(&builder, "- **Escalation:** The declared risk tier `%s` differs from engine routing `%s`.\n",
			assessment.Classification.DeclaredTier, assessment.Classification.Tier)
	}

	builder.WriteString("\n## Decision Gate\n\n")
	builder.WriteString("| Applicable | Supported | Partial | Missing | Expired | Evidence Coverage |\n")
	builder.WriteString("| ---: | ---: | ---: | ---: | ---: | ---: |\n")
	fmt.Fprintf(&builder, "| %d | %d | %d | %d | %d | %.1f%% |\n\n",
		assessment.Summary.Applicable, assessment.Summary.Supported, assessment.Summary.Partial,
		assessment.Summary.Missing, assessment.Summary.Expired, assessment.Summary.EvidenceCoverage)

	builder.WriteString("## System Context\n\n")
	builder.WriteString("| Attribute | Value |\n| --- | --- |\n")
	fmt.Fprintf(&builder, "| Owner | %s |\n", assessment.System.Owner)
	fmt.Fprintf(&builder, "| Role | %s |\n", assessment.System.Role)
	fmt.Fprintf(&builder, "| Domain | %s |\n", assessment.System.Domain)
	fmt.Fprintf(&builder, "| Lifecycle stage | %s |\n", assessment.System.LifecycleStage)
	fmt.Fprintf(&builder, "| Intended purpose | %s |\n", assessment.System.IntendedPurpose)
	fmt.Fprintf(&builder, "| Decision impact | %s |\n", assessment.System.DecisionImpact)
	fmt.Fprintf(&builder, "| Geographies | %s |\n", strings.Join(assessment.System.Geographies, ", "))
	fmt.Fprintf(&builder, "| Capabilities | %s |\n", strings.Join(assessment.System.Capabilities, ", "))
	fmt.Fprintf(&builder, "| Affected groups | %s |\n\n", strings.Join(assessment.System.AffectedGroups, ", "))

	builder.WriteString("## Framework Coverage\n\n")
	builder.WriteString("| Framework | Version | As Of | Applicable | Supported | Gaps | Pack SHA-256 |\n")
	builder.WriteString("| --- | --- | --- | ---: | ---: | ---: | --- |\n")
	for _, version := range assessment.PolicyVersions {
		applicable, supported := frameworkCounts(assessment.Controls, version.Framework)
		fmt.Fprintf(&builder, "| [%s](%s) | %s | %s | %d | %d | %d | `%s…` |\n",
			version.Framework, version.SourceURL, version.Version, version.AsOf,
			applicable, supported, applicable-supported, shortDigest(version.PackSHA256))
	}

	builder.WriteString("\n## Priority Findings\n\n")
	if len(assessment.Findings) == 0 {
		builder.WriteString("No evidence gaps were detected. Independent review is still required.\n")
	} else {
		builder.WriteString("| Severity | Framework | Control | Status | Finding |\n")
		builder.WriteString("| --- | --- | --- | --- | --- |\n")
		findings := append([]domain.Finding(nil), assessment.Findings...)
		sort.SliceStable(findings, func(i, j int) bool { return severityRank(findings[i].Severity) < severityRank(findings[j].Severity) })
		for _, finding := range findings {
			fmt.Fprintf(&builder, "| %s | %s | %s | %s | %s |\n",
				strings.ToUpper(finding.Severity), finding.Framework, finding.ControlID, finding.Status, finding.Message)
		}
	}

	builder.WriteString("\n## Detailed Control Matrix\n\n")
	builder.WriteString("| Framework | Control | Source | Severity | Status | Evidence |\n")
	builder.WriteString("| --- | --- | --- | --- | --- | --- |\n")
	for _, control := range assessment.Controls {
		if !control.Applicable {
			continue
		}
		evidenceIDs := controlEvidence(control)
		fmt.Fprintf(&builder, "| %s | %s — %s | %s | %s | **%s** | %s |\n",
			control.Framework, control.ID, control.Title, control.SourceRef, control.Severity,
			control.Status, strings.Join(evidenceIDs, ", "))
	}

	builder.WriteString("\n## Evidence Lineage\n\n")
	builder.WriteString("| Evidence | Type | Status | Owner | Integrity | SHA-256 |\n")
	builder.WriteString("| --- | --- | --- | --- | --- | --- |\n")
	for _, evidence := range assessment.Evidence {
		fmt.Fprintf(&builder, "| %s — %s | %s | %s | %s | %s | `%s…` |\n",
			evidence.ID, evidence.Title, evidence.Type, evidence.Status, evidence.Owner,
			evidence.Integrity, shortDigest(evidence.SHA256))
	}

	builder.WriteString("\n## Evidence Graph and Audit Receipt\n\n")
	fmt.Fprintf(&builder, "The evidence graph contains **%d nodes** and **%d typed edges**, linking the system to framework versions, controls, evidence artefacts, and findings.\n\n",
		len(assessment.Graph.Nodes), len(assessment.Graph.Edges))
	if assessment.Ledger != nil {
		fmt.Fprintf(&builder, "- Ledger sequence: `%d`\n", assessment.Ledger.Sequence)
		fmt.Fprintf(&builder, "- Event hash: `%s`\n", assessment.Ledger.EventHash)
		fmt.Fprintf(&builder, "- Previous hash: `%s`\n", assessment.Ledger.PreviousHash)
		fmt.Fprintf(&builder, "- Payload hash: `%s`\n", assessment.Ledger.PayloadHash)
	}

	builder.WriteString("\n## Interpretation and Limitations\n\n")
	builder.WriteString("Evidence coverage is the percentage of applicable controls supported by at least the configured minimum of verified, approved artefacts. It is not a compliance score. Applicability rules are a documented triage aid and must be reviewed by qualified legal, risk, security, and domain specialists. Framework mappings are versioned because regulatory and standards guidance changes over time.\n")

	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("create report directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(builder.String()), 0o600); err != nil {
		return fmt.Errorf("write assessment report: %w", err)
	}
	return nil
}

func shortDigest(value string) string {
	if len(value) <= 12 {
		return value
	}
	return value[:12]
}

func frameworkCounts(controls []domain.ControlResult, framework string) (int, int) {
	applicable, supported := 0, 0
	for _, control := range controls {
		if control.Framework != framework || !control.Applicable {
			continue
		}
		applicable++
		if control.Status == "supported" {
			supported++
		}
	}
	return applicable, supported
}

func controlEvidence(control domain.ControlResult) []string {
	result := []string{}
	for _, requirement := range control.Requirements {
		result = append(result, requirement.EvidenceIDs...)
	}
	sort.Strings(result)
	if len(result) == 0 {
		return []string{"—"}
	}
	return result
}

func severityRank(severity string) int {
	switch severity {
	case "critical":
		return 0
	case "high":
		return 1
	case "medium":
		return 2
	default:
		return 3
	}
}
