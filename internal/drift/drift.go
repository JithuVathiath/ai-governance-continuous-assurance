package drift

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

// Compare identifies changes in control scope and evidence status.
func Compare(before, after domain.Assessment) (domain.DriftReport, error) {
	if before.System.ID != after.System.ID {
		return domain.DriftReport{}, fmt.Errorf("cannot compare different systems: %s and %s", before.System.ID, after.System.ID)
	}
	report := domain.DriftReport{
		SchemaVersion: "1.0.0", SystemID: before.System.ID,
		BeforeAssessment: before.AssessmentID, AfterAssessment: after.AssessmentID,
		BeforeTier: before.Classification.Tier, AfterTier: after.Classification.Tier,
	}
	beforeMap := controlMap(before.Controls)
	afterMap := controlMap(after.Controls)
	keys := make([]string, 0, len(beforeMap)+len(afterMap))
	seen := map[string]bool{}
	for key := range beforeMap {
		keys = append(keys, key)
		seen[key] = true
	}
	for key := range afterMap {
		if !seen[key] {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	for _, key := range keys {
		beforeControl, hadBefore := beforeMap[key]
		afterControl, hasAfter := afterMap[key]
		beforeStatus := "out-of-scope"
		afterStatus := "out-of-scope"
		control := afterControl
		if hadBefore {
			beforeStatus = statusFor(beforeControl)
			if !hasAfter {
				control = beforeControl
			}
		}
		if hasAfter {
			afterStatus = statusFor(afterControl)
		}
		if beforeStatus == afterStatus {
			continue
		}

		changeType := "changed"
		switch {
		case beforeStatus == "out-of-scope" && afterStatus != "out-of-scope":
			changeType = "newly-applicable"
			report.NewControls++
		case beforeStatus != "supported" && afterStatus == "supported":
			changeType = "resolved"
			report.ResolvedControls++
		case beforeStatus == "supported" && afterStatus != "supported":
			changeType = "regressed"
			report.RegressedControls++
		default:
			report.ChangedControls++
		}
		report.Items = append(report.Items, domain.DriftItem{
			Key: key, Framework: control.Framework, ControlID: control.ID, Title: control.Title,
			Before: beforeStatus, After: afterStatus, ChangeType: changeType, Severity: control.Severity,
		})
	}
	return report, nil
}

// WriteMarkdown renders a recruiter-readable change-impact report.
func WriteMarkdown(path string, report domain.DriftReport) error {
	var builder strings.Builder
	fmt.Fprintf(&builder, "# Aegis Control-Drift Report\n\n")
	fmt.Fprintf(&builder, "**System:** `%s`  \n", report.SystemID)
	fmt.Fprintf(&builder, "**Assessment transition:** `%s` → `%s`  \n", report.BeforeAssessment, report.AfterAssessment)
	fmt.Fprintf(&builder, "**Risk routing:** `%s` → `%s`\n\n", report.BeforeTier, report.AfterTier)
	fmt.Fprintf(&builder, "## Change Summary\n\n")
	fmt.Fprintf(&builder, "| Newly Applicable | Resolved | Regressed | Other Changes |\n")
	fmt.Fprintf(&builder, "| ---: | ---: | ---: | ---: |\n")
	fmt.Fprintf(&builder, "| %d | %d | %d | %d |\n\n", report.NewControls, report.ResolvedControls, report.RegressedControls, report.ChangedControls)
	fmt.Fprintf(&builder, "## Control Changes\n\n")
	fmt.Fprintf(&builder, "| Change | Framework | Control | Before | After | Severity |\n")
	fmt.Fprintf(&builder, "| --- | --- | --- | --- | --- | --- |\n")
	for _, item := range report.Items {
		fmt.Fprintf(&builder, "| %s | %s | %s — %s | %s | %s | %s |\n",
			item.ChangeType, item.Framework, item.ControlID, item.Title, item.Before, item.After, item.Severity)
	}
	builder.WriteString("\n## Interpretation\n\n")
	builder.WriteString("This report shows how policy scope and evidence readiness changed between two immutable assessments. It does not determine legal compliance. Newly applicable or regressed controls require independent review before release.\n")

	if err := os.MkdirAll(filepath.Dir(path), 0o750); err != nil {
		return fmt.Errorf("create drift report directory: %w", err)
	}
	if err := os.WriteFile(path, []byte(builder.String()), 0o600); err != nil {
		return fmt.Errorf("write drift report: %w", err)
	}
	return nil
}

func controlMap(controls []domain.ControlResult) map[string]domain.ControlResult {
	result := make(map[string]domain.ControlResult, len(controls))
	for _, control := range controls {
		result[control.PackID+":"+control.ID] = control
	}
	return result
}

func statusFor(control domain.ControlResult) string {
	if !control.Applicable {
		return "out-of-scope"
	}
	return control.Status
}
