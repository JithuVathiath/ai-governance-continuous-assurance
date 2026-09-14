package drift

import (
	"path/filepath"
	"testing"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

func TestCompareClassifiesControlChanges(t *testing.T) {
	before := domain.Assessment{AssessmentID: "before", System: domain.System{ID: "system"}, Classification: domain.Classification{Tier: "high"}, Controls: []domain.ControlResult{
		{PackID: "p", ID: "resolved", Framework: "F", Status: "missing", Applicable: true},
		{PackID: "p", ID: "regressed", Framework: "F", Status: "supported", Applicable: true},
	}}
	after := domain.Assessment{AssessmentID: "after", System: domain.System{ID: "system"}, Classification: domain.Classification{Tier: "high"}, Controls: []domain.ControlResult{
		{PackID: "p", ID: "resolved", Framework: "F", Status: "supported", Applicable: true},
		{PackID: "p", ID: "regressed", Framework: "F", Status: "partial", Applicable: true},
		{PackID: "p", ID: "new", Framework: "F", Status: "missing", Applicable: true},
	}}
	report, err := Compare(before, after)
	if err != nil {
		t.Fatal(err)
	}
	if report.ResolvedControls != 1 || report.RegressedControls != 1 || report.NewControls != 1 {
		t.Fatalf("unexpected drift: %#v", report)
	}
	if err := WriteMarkdown(filepath.Join(t.TempDir(), "drift.md"), report); err != nil {
		t.Fatal(err)
	}
}

func TestCompareRejectsDifferentSystems(t *testing.T) {
	_, err := Compare(domain.Assessment{System: domain.System{ID: "one"}}, domain.Assessment{System: domain.System{ID: "two"}})
	if err == nil {
		t.Fatal("expected identity mismatch")
	}
}
