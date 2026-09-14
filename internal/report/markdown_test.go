package report

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/JithuVathiath/ai-governance-continuous-assurance/internal/domain"
)

func TestWriteHandlesShortDigests(t *testing.T) {
	path := filepath.Join(t.TempDir(), "report.md")
	assessment := domain.Assessment{
		AssessmentID: "asg", GeneratedAt: "2026-09-13T12:00:00Z", Disclaimer: "Not legal advice.",
		System:         domain.System{Name: "Demo", ID: "demo", Version: "1", Owner: "Owner"},
		Classification: domain.Classification{Tier: "minimal", Reasons: []string{"demo"}},
		PolicyVersions: []domain.PolicyVersion{{Framework: "F", Version: "1", SourceURL: "https://example.test", PackSHA256: "short"}},
		Evidence:       []domain.Evidence{{ID: "E", Title: "Evidence", SHA256: "tiny"}},
		Summary:        domain.Summary{Decision: "conditional-review"},
	}
	if err := Write(path, assessment); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "Aegis Continuous-Assurance Report") || !strings.Contains(string(data), "CONDITIONAL-REVIEW") {
		t.Fatalf("unexpected report: %s", data)
	}
}

func TestReportHelpers(t *testing.T) {
	controls := []domain.ControlResult{
		{Framework: "F", Applicable: true, Status: "supported", Requirements: []domain.RequirementResult{{EvidenceIDs: []string{"B", "A"}}}},
		{Framework: "F", Applicable: true, Status: "missing"},
		{Framework: "Other", Applicable: true, Status: "supported"},
	}
	applicable, supported := frameworkCounts(controls, "F")
	if applicable != 2 || supported != 1 {
		t.Fatalf("unexpected counts: %d %d", applicable, supported)
	}
	ids := controlEvidence(controls[0])
	if strings.Join(ids, ",") != "A,B" || controlEvidence(controls[1])[0] != "—" {
		t.Fatalf("unexpected evidence IDs: %#v", ids)
	}
	if severityRank("critical") >= severityRank("high") || severityRank("high") >= severityRank("medium") || severityRank("medium") >= severityRank("low") {
		t.Fatal("severity order is incorrect")
	}
	if shortDigest("123456789012345") != "123456789012" {
		t.Fatal("digest was not shortened")
	}
}
